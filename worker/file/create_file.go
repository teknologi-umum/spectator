package file

import (
	"context"
	"log"
	"time"

	"worker/common"
	loggerpb "worker/logger_proto"

	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

// CreateFile creates a file and concurently uploads it to the MinIO bucket.
// This function should be called as a goroutine.
//
// It will not panic, instead the panic will be caught
func (d *Dependency) CreateFile(requestID string, sessionID uuid.UUID) {
	// Defer a func that will recover from panic.
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		log.Println(r.(error))
		d.Logger.Log(
			r.(error).Error(),
			loggerpb.Level_ERROR.Enum(),
			requestID,
			map[string]string{
				"session_id": sessionID.String(),
				"function":   "CreateFile",
				"info":       "recovering from panic",
			},
		)
	}()

	log.Printf("[%s] Got request to create file for session %s", requestID, sessionID.String())

	// Let's create a new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// Now we fetch all the data and put it on 4 files.
	// We will not use a goroutine to do this, because we want to allow
	// the core backend to continue inserting other data, so we won't cause
	// a bottleneck to the InfluxDB system.
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// cfDeps contains the struct that implements a log function
	// to make the code more sane.
	cfDeps := &createFile{deps: d}

	// Keystroke events queries
	outputKeystroke, err := d.QueryKeystrokes(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query keystrokes", requestID, sessionID)
		return
	}

	// Mouse events queries
	outputMouseClick, err := d.QueryMouseClick(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query mouse click", requestID, sessionID)
		return
	}

	outputMouseMove, err := d.QueryMouseMove(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query mouse move", requestID, sessionID)
		return
	}

	outputMouseScrolled, err := d.QueryMouseScrolled(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query mouse scrolled", requestID, sessionID)
		return
	}

	outputMouseDistanceTraveled, err := d.QueryMouseDistanceTraveled(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query mouse distance traveled", requestID, sessionID)
		return
	}

	// Session events queries (regarding the user information and so on forth)
	outputPersonalInfo, err := d.QueryPersonalInfo(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query personal info", requestID, sessionID)
		return
	}

	outputSamBeforeTest, err := d.QueryBeforeExamSam(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query before exam sam", requestID, sessionID)
		return
	}

	outputSamAfterTest, err := d.QueryAfterExamSam(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query after exam sam", requestID, sessionID)
		return
	}

	outputExamStarted, err := d.QueryExamStarted(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query exam started", requestID, sessionID)
		return
	}

	outputExamEnded, err := d.QueryExamEnded(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query exam ended", requestID, sessionID)
		return
	}

	outputExamForfeited, err := d.QueryExamForfeited(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query exam forfeited", requestID, sessionID)
		return
	}

	outputExamIDEReloaded, err := d.QueryExamIDEReloaded(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query exam idereloaded", requestID, sessionID)
		return
	}

	outputDeadlinePassed, err := d.QueryDeadlinePassed(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query deadline passed", requestID, sessionID)
		return
	}

	outputFunfact, err := d.QueryFunfact(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query funfact", requestID, sessionID)
		return
	}

	// Solution events queries
	outputSolutionRejected, err := d.QuerySolutionRejected(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query solution rejected", requestID, sessionID)
		return
	}

	outputSolutionAccepted, err := d.QuerySolutionAccepted(ctx, queryAPI, sessionID)
	if err != nil {
		cfDeps.sendErrorLog(err, "failed to query solution accepted", requestID, sessionID)
		return
	}

	// Then, we'll write to 2 different files with 2 different formats.
	// Do this repeatedly for each event.
	//
	// So in the end, we have multiple files,
	// one is about the keystroke & mouse events
	// one is about coding test results
	// one is all about the user (personal info, sam test)
	//
	// After that, store the file into MinIO
	// then, put the MinIO link on the influxdb database
	// in a different bucket. You might want to check and do a
	// create if not exists on the bucket.
	// So you'd make sure you're not inserting data into a
	// nil bucket.

	// Join every events into their own event types.
	userEvents := &UserEvents{
		SelfAssessmentManekinBeforeTest: outputSamBeforeTest,
		SelfAssessmentManekinAfterTest:  outputSamAfterTest,
		PersonalInfo:                    outputPersonalInfo,
		ExamStarted:                     outputExamStarted,
		ExamEnded:                       outputExamEnded,
		ExamForfeited:                   outputExamForfeited,
		ExamIDEReloaded:                 outputExamIDEReloaded,
		DeadlinePassed:                  outputDeadlinePassed,
		Funfact:                         outputFunfact,
	}

	mouseEvents := &MouseEvents{
		MouseClick:            outputMouseClick,
		MouseMoved:            outputMouseMove,
		MouseScrolled:         outputMouseScrolled,
		MouseDistanceTraveled: outputMouseDistanceTraveled,
	}

	keystrokeEvents := &KeystrokeEvents{
		Keystroke: outputKeystroke,
	}

	solutionEvents := &SolutionEvents{
		SolutionAccepted: outputSolutionAccepted,
		SolutionRejected: outputSolutionRejected,
	}

	writeAPI := d.DB.WriteAPIBlocking(d.DBOrganization, common.BucketFileEvents)

	studentNumber := outputPersonalInfo.StudentNumber

	g, gctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, userEvents, "user_events", studentNumber, requestID, sessionID)
	})
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, mouseEvents, "mouse_events", studentNumber, requestID, sessionID)
	})
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, keystrokeEvents, "keystroke_events", studentNumber, requestID, sessionID)
	})
	g.Go(func() error {
		return d.convertAndUpload(gctx, writeAPI, solutionEvents, "solution_events", studentNumber, requestID, sessionID)
	})

	if err := g.Wait(); err != nil {
		cfDeps.sendErrorLog(err, "failed to convert and upload", requestID, sessionID)
		return
	}

	log.Printf("[%s] Successfully converted and uploaded all events for session: %s", requestID, sessionID.String())
}

// createFile is a struct that implements sendErrorLog method.
// This struct must be used strictly within the CreateFile method.
type createFile struct {
	deps *Dependency
}

func (d *createFile) sendErrorLog(err error, additionalInfo string, requestID string, sessionID uuid.UUID) {
	d.deps.Logger.Log(
		err.Error(),
		loggerpb.Level_ERROR.Enum(),
		requestID,
		map[string]string{
			"session_id": sessionID.String(),
			"function":   "CreateFile",
			"info":       additionalInfo,
		},
	)
}
