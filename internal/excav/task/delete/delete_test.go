package delete

import (
	"testing"
)

// scenario:
//   when: we delete text between '>' and '<'
//   then: the text should be without part in a middle
func TestDeleteLines(t *testing.T) {
	text := "Hello > We want to cut < Johnny"
	out := del([]byte(text), ">", "<")

	if string(out) != "Hello  Johnny" {
		t.FailNow()
	}
}

// scenario:
//   when: we delete text where is no end present
//   then: the text should be without part in a middle
func TestDeleteLinesNoEnd(t *testing.T) {
	text := "Hello > We want to cut"
	out := del([]byte(text), ">", "<")

	if string(out) != "Hello " {
		t.FailNow()
	}
}

// scenario:
//   when: we delete text where is no begin anchor '>' present
//   then: the text should be without part in a middle
func TestDeleteLinesNoBegin(t *testing.T) {
	text := "Hello + We want to cut < Johhny"
	out := del([]byte(text), ">", "<")

	if string(out) != text {
		t.FailNow()
	}
}

// scenario:
//   given: text with '>' and '<' anchors
//   when: I call del for those anchors
//   then: text shouldn't contain any text between '>' and '<'
func TestDeleteAll(t *testing.T) {
	text := "Hello > We want to cut < Johnny > Yet another to cut < Bravo"
	out := del([]byte(text), ">", "<")

	if string(out) != "Hello  Johnny  Bravo" {
		t.FailNow()
	}
}
