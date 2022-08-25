package excav

import (
	"fmt"
	"testing"
)

func TestContextSaveLoad(t *testing.T) {
	tmpDir := TempDirectory()

	ctx := PatchContext{
		RepoName: "hello/world",
	}

	saveContext(tmpDir, &ctx)
	fmt.Println(tmpDir)
}
