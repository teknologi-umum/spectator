package file

import (
	"context"

	"github.com/google/uuid"
)

type File struct {
	CSVFile  string
	JSONFile string
	// TODO: add more here
}

func (d *Dependency) ListFiles(ctx context.Context, sessionID uuid.UUID) ([]File, error) {
	// TODO: add more here
	return []File{}, nil
}
