package cmd

import (
	"github.com/sn3d/excav/cmd/apply"
	"github.com/sn3d/excav/cmd/bulk"
	"github.com/sn3d/excav/cmd/diff"
	"github.com/sn3d/excav/cmd/discard"
	"github.com/sn3d/excav/cmd/initialize"
	"github.com/sn3d/excav/cmd/inventory"
	"github.com/sn3d/excav/cmd/patch"
	"github.com/sn3d/excav/cmd/push"
	"github.com/sn3d/excav/cmd/show"
	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	initialize.Cmd,
	apply.Cmd,
	bulk.Cmd,
	push.Cmd,
	diff.Cmd,
	show.Cmd,
	discard.Cmd,
	patch.Cmd,
	inventory.Cmd,
}
