package writer

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/haya14busa/go-sarif/sarif"
	"github.com/reviewdog/errorformat"
)

type Sarif struct {
	mu      sync.Mutex
	data    *sarif.Sarif
	run     *sarif.Run
	w       io.Writer
	srcRoot string
}

type SarifOption struct {
	ToolName string
}

func NewSarif(w io.Writer, opt SarifOption) (*Sarif, error) {
	if opt.ToolName == "" {
		return nil, errors.New("-sarif.tool-name is required.")
	}
	s := &Sarif{w: w, data: sarif.NewSarif(), run: &sarif.Run{
		Tool: sarif.Tool{
			Driver: sarif.ToolComponent{
				Name: opt.ToolName,
			},
		},
	}}
	s.srcRoot, _ = os.Getwd()
	return s, nil
}

func (s *Sarif) Write(e *errorformat.Entry) error {
	result := sarif.Result{}

	// Set Level
	switch e.Type {
	case 'e', 'E':
		result.Level = sarif.Error.Ptr()
	case 'w', 'W':
		result.Level = sarif.Warning.Ptr()
	case 'n', 'N', 'i', 'I': // Handle info as note.
		result.Level = sarif.Note.Ptr()
	}

	// Set Message
	result.Message.Text = sarif.String(e.Text)

	// Set Location
	loc := &sarif.PhysicalLocation{
		ArtifactLocation: &sarif.ArtifactLocation{
			URIBaseID: sarif.String("%SRCROOT%"),
		},
		Region: &sarif.Region{},
	}
	if e.Lnum != 0 {
		loc.Region.StartLine = sarif.Int64(int64(e.Lnum))
	}
	if e.EndLnum != 0 {
		loc.Region.EndLine = sarif.Int64(int64(e.EndLnum))
	}
	if e.Col != 0 {
		// Errorformat.Col is not usually unicodeCodePoints, but let's just keep it
		// as is...
		loc.Region.StartColumn = sarif.Int64(int64(e.Col))
	}
	if e.EndCol != 0 {
		loc.Region.EndColumn = sarif.Int64(int64(e.EndCol))
	}
	if e.Filename != "" {
		uri := e.Filename
		if filepath.IsAbs(e.Filename) && s.srcRoot != "" {
			uri, _ = filepath.Rel(s.srcRoot, e.Filename)
		}
		loc.ArtifactLocation.URI = sarif.String(uri)
	}
	result.Locations = append(result.Locations, sarif.Location{
		PhysicalLocation: loc,
	})

	s.mu.Lock()
	defer s.mu.Unlock()
	s.run.Results = append(s.run.Results, result)
	return nil
}

func (s *Sarif) Flush() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	e := json.NewEncoder(s.w)
	e.SetIndent("", "  ")
	s.data.Runs = append(s.data.Runs, *s.run)
	return e.Encode(s.data)
}
