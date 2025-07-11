package stdhelper

import (
	"bytes"
	"io"
	"os"
)

var OutputWriter io.Writer = os.Stdout

func CaptureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
