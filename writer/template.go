package writer

import (
	"io"
	"text/template"

	"github.com/haya14busa/errorformat"
)

type Template struct {
	tmpl *template.Template
	w    *trackingWriter
}

func NewTemplate(tmpl *template.Template, w io.Writer) *Template {
	return &Template{tmpl: tmpl, w: newTrackingWriter(w)}
}

func (t *Template) Write(e *errorformat.Entry) error {
	if err := t.tmpl.Execute(t.w, e); err != nil {
		return err
	}
	if t.w.NeedNL() {
		if _, err := t.w.WriteNL(); err != nil {
			return err
		}
	}
	return nil
}
