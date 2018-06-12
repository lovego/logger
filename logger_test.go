package logger

import (
	"bytes"
	"testing"
)

func TestInfo(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	log := New(buf)
	log.Info("the", `message`)
	t.Log(buf.String())
}
