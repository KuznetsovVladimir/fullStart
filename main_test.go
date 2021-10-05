package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOK = `1
2
3
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOK))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("test for Ok failed - error")
	}
	if out.String() != testOkResult {
		t.Errorf("test fo Ok failed - results not match")
	}
}

var testFail = `1
2
1`

func TestForError(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)
	if err == nil {
		t.Errorf("test for Ok failed - error")
	}
}
