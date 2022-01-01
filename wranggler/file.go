package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"

	pb "worker/proto"
)

// GenerateFile is the handler for generating file into CSV and JSON based on
// the input data (which only contains the Session ID).
func (d *Dependency) GenerateFile(ctx context.Context, in *pb.Member) (*pb.EmptyResponse, error) {
	sessionID, err := uuid.Parse(in.GetSessionId())
	if err != nil {
		return &pb.EmptyResponse{}, fmt.Errorf("parsing uuid: %v", err)
	}

	go d.CreateFile(sessionID)

	return &pb.EmptyResponse{}, nil
}

func (d *Dependency) CreateFile(sessionID uuid.UUID) {
	// Defer a func that will recover from panic.
	// TODO: Send this data into the Logging service.
	defer func(){
		r := recover()
		if r != nil {
			log.Println(r.(error))
		}
	}()

	// Let's create a new context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	// Now we fetch all the data with the _actor being sessionID.String()
	queryAPI := d.DB.QueryAPI(d.DBOrganization)

	// keystroke and mouse
	_, err := queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketInputEvents+`")
		|> filter(fn : (r) => r["session_id"] == "`+sessionID.String()+` and (
			(r["_measurement"] == "coding_event_mouseclick") or
			(r["_measurement"] == "coding_event_movemove") or
			(r["_measurement"] == "coding_event_mouseclick"))
		`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	// coding test result
	_, err = queryAPI.Query(
		ctx,
		`from(bucket: "`+BucketSessionEvents+`")
		|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
		|> fliter(fn: (r) => r["_measurement"] == "code_submission")`,
	)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
		return
	}

	// user
	_, err = queryAPI.Query(ctx, `
	from(bucket: "`+BucketSessionEvents+`")
	|> filter(fn: (r) => r["session_id"] == "`+sessionID.String()+`")
	|> filter(fn: (r) => (r["event"] == "sam_test_before") or
		(r["event"] == "personal_info"))
	`)
	if err != nil {
		// we send a http request to the logger service
		// for now, we'll just do this:
		log.Println(err)
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
}

// TODO: delete this one
type SampleInput struct {
	Time  time.Time `json:"timestamp" csv:"timestamp"`
	Actor string    `json:"actor" csv:"actor"`
	X     int       `json:"x" csv:"x"`
	Y     int       `json:"y" csv:"y"`
}

// TODO: change the SampleInput type with an actual working type
// that resembles the influxdb schema
func ConvertDataToJSON(input []SampleInput) ([]byte, error) {
	data, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return []byte{}, err
	}

	return data, err
}

// TODO: change the SampleInput type with an actual working type
// that resembles the influxdb schema
func ConvertDataToCSV(input []SampleInput) ([]byte, error) {
	w := &bytes.Buffer{}
	writer := csv.NewWriter(w)
	// Because csv package does not have something like
	// json.Marshal, we'll gonna do what Thanos did.
	//
	// "Fine. I'll do it myself."

	// Create the CSV headers first
	structType := reflect.TypeOf(input[0])
	headers := make([]string, structType.NumField())
	for i := 0; i < structType.NumField(); i++ {
		headers = append(headers, structType.Field(i).Tag.Get("csv"))
	}

	err := writer.Write(headers)
	if err != nil {
		return []byte{}, err
	}

	for _, inputItem := range input {
		// Struct are always in-order, so it's easy to
		// put it into the temporary
		structValue := reflect.ValueOf(inputItem)
		data := make([]string, structValue.NumField())

		for k := 0; k < structValue.NumField(); k++ {
			currentValue := structValue.Field(k)

			switch currentValue.Interface().(type) {
			case bool:
				data = append(data, strconv.FormatBool(currentValue.Bool()))
				continue
			case string:
				data = append(data, currentValue.String())
				continue
			case uint:
				data = append(data, strconv.FormatUint(currentValue.Uint(), 10))
			case int64:
				data = append(data, strconv.FormatInt(currentValue.Int(), 10))
				continue
			case int:
				data = append(data, strconv.FormatInt(currentValue.Int(), 10))
				continue
			case time.Time:
				t, ok := currentValue.Interface().(time.Time)
				if !ok {
					return []byte{}, fmt.Errorf("struct name of %s has a type of time.Time yet cannot be parsed", currentValue.Type().Name())
				}
				data = append(data, t.Format(time.RFC3339Nano))
				continue
			default:
				return []byte{}, fmt.Errorf("struct name of %s has a weird and unsupported type", currentValue.Type().Name())
			}
		}

		err := writer.Write(data)
		if err != nil {
			return []byte{}, err
		}
	}

	writer.Flush()
	if writer.Error() != nil {
		return []byte{}, fmt.Errorf("last csv write error: %v", err)
	}

	return w.Bytes(), nil
}
