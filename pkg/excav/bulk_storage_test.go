package excav

import (
	"fmt"
	"testing"

	"github.com/sn3d/excav/pkg/dir"
)

func TestContextSaveLoad(t *testing.T) {
	tmpDir := dir.Temp()

	ctx := PatchContext{
		RepoName: "hello/world",
	}

	saveContext(tmpDir, &ctx)
	fmt.Println(tmpDir)
}
