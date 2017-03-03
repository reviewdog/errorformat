package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		in       string
		efms     []string
		entryFmt string
		name     string
		want     string
	}{
		{
			in: `golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
`,
			efms:     []string{"%f:%l:%c: %m"},
			entryFmt: "{{.String}}",
			want: `golint.new.go|3 col 5| exported var V should have comment or be unexported
golint.new.go|5 col 5| exported var NewError1 should have comment or be unexported
golint.new.go|7 col 1| comment on exported function F should be of the form "F ..."
golint.new.go|11 col 1| comment on exported function F2 should be of the form "F2 ..."
`,
		},
		{
			in: `golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
`,
			efms:     []string{"%f:%l:%c: %m"},
			entryFmt: "{{.Filename}}",
			want: `golint.new.go
golint.new.go
golint.new.go
golint.new.go
`,
		},
		{
			in: `golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
`,
			entryFmt: "{{.Filename}}",
			name:     "golint",
			want: `golint.new.go
golint.new.go
golint.new.go
golint.new.go
`,
		},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		if err := run(strings.NewReader(tt.in), out, tt.efms, "", tt.entryFmt, tt.name, false); err != nil {
			t.Error(err)
		}
		if got := out.String(); got != tt.want {
			t.Errorf("got:\n%v\nwant:\n%v", got, tt.want)
		}
	}
}

func TestRun_checkstyle(t *testing.T) {
	tests := []struct {
		in   string
		efms []string
		want string
	}{
		{
			in: `golint.new.go:3:5: exported var V should have comment or be unexported
golint.new.go:5:5: exported var NewError1 should have comment or be unexported
golint.new.go:7:1: comment on exported function F should be of the form "F ..."
golint.new.go:11:1: comment on exported function F2 should be of the form "F2 ..."
`,
			efms: []string{"%f:%l:%c: %m"},
			want: `<?xml version="1.0" encoding="UTF-8"?>
<checkstyle version="1.0">
  <file name="golint.new.go">
    <error column="5" line="3" message="exported var V should have comment or be unexported"></error>
    <error column="5" line="5" message="exported var NewError1 should have comment or be unexported"></error>
    <error column="1" line="7" message="comment on exported function F should be of the form &#34;F ...&#34;"></error>
    <error column="1" line="11" message="comment on exported function F2 should be of the form &#34;F2 ...&#34;"></error>
  </file>
</checkstyle>
`,
		},
	}
	for _, tt := range tests {
		out := new(bytes.Buffer)
		if err := run(strings.NewReader(tt.in), out, tt.efms, "checkstyle", "", "", false); err != nil {
			t.Error(err)
		}
		if got := out.String(); got != tt.want {
			t.Errorf("got:\n%v\nwant:\n%v", got, tt.want)
		}
	}
}

func TestRun_unknown_writer(t *testing.T) {
	if err := run(nil, nil, nil, "unknown", "", "", false); err == nil {
		t.Error("error expected but got nil")
	}
}

func TestRun_list(t *testing.T) {
	out := new(bytes.Buffer)
	if err := run(nil, out, nil, "", "", "", true); err != nil {
		t.Error(err)
	}
	t.Log(out.String())
}
