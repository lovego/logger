package logger

import (
	"os"
	"testing"
)

func TestPrintf(t *testing.T) {
	log := New(`test: `, nil, nil)
	log.Printf("Print: %s", `message`)
	println()
}

func TestErrorf(t *testing.T) {
	log := New(`test: `, nil, nil)
	log.Errorf("Error: %s", `error`)
	println()
}

func TestFatalf(t *testing.T) {
	log := New(`test: `, nil, nil)
	log.Fatalf("Fatal: %s", `fatal error`)
	println()
}

func TestPanicf(t *testing.T) {
	log := New(`test: `, os.Stdout, nil)
	log.Panicf("Fatal: %s", `panic error`)
	println()
}
