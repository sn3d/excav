package termui

import "fmt"

// Info is used for regular messages they're print to terminal
func Info(text string, a ...interface{}) {
	fmt.Printf(text+"\n", a)
}

func PrintRepositories(repos []string) {
	for _, repo := range repos {
		fmt.Printf("%s\n", repo)
	}
}

func PrintRepository(repository string, tags []string, params map[string]string) {
	fmt.Printf("%s: %s\n", Magenta("Repository"), BrightWhite(repository))

	// print tags
	if len(tags) > 0 {
		fmt.Printf("   %s: %s\n", BrightWhite("tags"), tags)
	}

	// print params
	if len(params) > 0 {
		fmt.Printf("   %s:\n", BrightWhite("params"))
		for key, val := range params {
			fmt.Printf("     %s: '%s'\n", key, val)
		}
	}

	fmt.Printf("\n")
}
