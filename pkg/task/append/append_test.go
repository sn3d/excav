package append

import (
	"fmt"
	"testing"
)

func Test_AppendBegin(t *testing.T) {

	// when we insert to beginning some text
	data := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	appended := appendAtBegin("START. ", []byte(data))

	// then the new text must start with content
	if string(appended) != "START. Lorem ipsum dolor sit amet, consectetur adipiscing elit." {
		t.FailNow()
	}
}

func Test_AppendEnd(t *testing.T) {

	// when we insert to end some text
	data := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	appended := appendAtEnd(" END.", []byte(data))

	// then the new text must end with content
	if string(appended) != "Lorem ipsum dolor sit amet, consectetur adipiscing elit. END." {
		t.FailNow()
	}
}

func Test_AppendBefore(t *testing.T) {
	text := `
		func() {
			// +excav: Hello World
		}`

	appended := appendBefore("\\/\\/ \\+excav:(.*)", "// line1 \n//line2", []byte(text))

	fmt.Print(string(appended))
	fmt.Print("\n")
}

func Test_AppendAfter(t *testing.T) {
	text := `
		func() {
			// +excav: Hello World
		}`

	appended := appendAfter("\\/\\/ \\+excav:(.*)", "// line1\n// line2", []byte(text))

	fmt.Print(string(appended))
	fmt.Print("\n")
}
