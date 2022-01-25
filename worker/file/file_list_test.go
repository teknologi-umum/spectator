package file_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/google/uuid"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func TestListFiles(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		t.Errorf("failed to generate uuid: %v", err)
	}

	writeSessionAPI := deps.DB.WriteAPIBlocking(deps.DBOrganization, deps.BucketSessionEvents)

	for i := 0; i < 50; i++ {
		for _, x := range []string{"keystroke", "mouse_click", "mouse_move", "personal_info", "sam_test"} {
			e := influxdb2.NewPointWithMeasurement("test_result")
			studentNumber := "a"
			e.AddTag("session_id", id.String())
			e.AddTag("student_number", studentNumber)
			e.AddField("file_csv_url", "/public/"+studentNumber+"_"+x+".csv")
			e.AddField("file_json_url", "/public/"+studentNumber+"_"+x+".json")
			e.SetTime(time.Now())

			err = writeSessionAPI.WritePoint(ctx, e)
			if err != nil {
				t.Fatalf("failed to write %s test result: %v", x, err)
				return
			}

			f, err := os.Create("./" + studentNumber + "_" + x + ".csv")
			if err != nil {
				t.Errorf("creating a file: %v", err)
				return
			}
			defer f.Close()

			_, err = f.Write([]byte(x))
			if err != nil {
				t.Errorf("writing to a file: %v", err)
				return
			}

			err = f.Sync()
			if err != nil {
				t.Errorf("syncing a file: %v", err)
				return
			}

			_, err = f.Stat()
			if err != nil {
				t.Errorf("getting file stat: %v", err)
				return
			}

			f, err = os.Open("./" + studentNumber + "_" + x + ".json")
			if err != nil {
				t.Errorf("opening a file: %v", err)
				return
			}
			defer f.Close()
		}
	}

	result, err := deps.ListFiles(ctx, id)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	pathJSON, err := filepath.Glob("./*_*.json")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	pathCSV, err := filepath.Glob("./*_*.csv")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(result) != 50 {
		t.Errorf("Expected 50 file, got %d", len(result))
	}

	for _, i := range append(pathJSON, pathCSV...) {
		err = os.Remove(i)
		if err != nil {
			t.Errorf("removing a file: %v", err)
			return
		}
	}
}
