package termui

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

// CliError just wrap the error and add
// CLI context. We need this context for printing
// appropriate help
type cliError struct {
	Err error
	Ctx *cli.Context
}

func (e cliError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func (e cliError) Unwrap() error {
	return e.Err
}

// Wrap any error and add CLI context into error
func CliError(ctx *cli.Context, err error) error {
	return cliError{
		Err: err,
		Ctx: ctx,
	}
}

// ShowError is responsible for printing all errors. Errors shouldn't
// be printed directly but returned to main function
func ShowError(err error) {
	var cerr cliError
	if errors.As(err, &cerr) {
		fmt.Printf("%s: %v\n\n",Red("ERROR"), cerr)
		cli.ShowSubcommandHelp(cerr.Ctx)
	} else {
		fmt.Printf("%s: %v\n",Red("UNKNOWN ERROR"), cerr)
	}
}
