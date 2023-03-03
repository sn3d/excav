package termui

import (
	"fmt"

	"github.com/sn3d/excav/api"
)

type Listener struct {
}

func (l *Listener) OnEvent(ev api.Event) {
	switch ev := ev.(type) {

	case api.TaskEnd:
		if ev.Error == nil {
			fmt.Printf("   [%s] %s%s\n", Green(CheckMark), Grey("task:"), ev.Task)
		} else {
			fmt.Printf("   [%s] %s%s\n", Red(XMark), ("task:"), ev.Task)
			fmt.Printf("        %s: %v\n", Red("error"), ev.Error)
		}

	case api.PatchingStarted:
		fmt.Printf("\n%s:\n", BrightWhite(ev.Repo))

	case api.Pushed:
		if ev.ErrorMsg != "" {
			fmt.Printf("%s: error: %s\n", ev.Repo, ev.ErrorMsg)
		} else {
			fmt.Printf("%s: pushed\n", ev.Repo)
		}

	case api.ReposSelected:
		fmt.Print("\n")
		for _, repoName := range ev.RepoNames {
			fmt.Printf("%s \n", repoName)
		}
		fmt.Print("\n")

	default:
		// do nothing
	}
}
