// Test tests for echo_threeways

package ch1

import (
	"bytes"
	"flag"
	"io"
	"os"
	"testing"
)

var origStdout *os.File

// Run tests through TestMain so that we can redirect stdout and also populate
// os.Args.
func TestMain(m *testing.M) {
	flag.Parse()
	arglen := 100
	os.Args = make([]string, arglen, arglen)
	for i, _ := range os.Args {
		os.Args[i] = "Yet another argument " + string(i)
	}
	os.Exit(m.Run())
}

func redirectStdout() {
	origStdout = os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		// Zero the buffer, don't worry abou the bytes copied.
		buf.Reset()
	}()
}

func restoreStdout() {
	os.Stdout.Close()
	os.Stdout = origStdout
}

func BenchmarkEchoStringCat(b *testing.B) {
	redirectStdout()
	for n := 0; n < b.N; n++ {
		EchoStringCat()
	}
	restoreStdout()
}

func BenchmarkEchoPrintParts(b *testing.B) {
	redirectStdout()
	for n := 0; n < b.N; n++ {
		EchoPrintParts()
	}
	restoreStdout()
}

func BenchmarkEchoStringJoin(b *testing.B) {
	redirectStdout()
	for n := 0; n < b.N; n++ {
		EchoStringJoin()
	}
	restoreStdout()
}
