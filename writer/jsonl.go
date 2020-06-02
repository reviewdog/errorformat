package writer

import (
	"encoding/json"
	"io"

	"github.com/reviewdog/errorformat"
)

// JSONL represents JSONL (http://jsonlines.org/) based writer.
type JSONL struct {
	encoder *json.Encoder
}

func NewJSONL(w io.Writer) *JSONL {
	return &JSONL{encoder: json.NewEncoder(w)}
}

func (j *JSONL) Write(e *errorformat.Entry) error {
	return j.encoder.Encode(e)
}
