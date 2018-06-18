package main

import (
	"context"
	"runtime"

	"github.com/pkg/errors"
	"gopkg.in/segmentio/analytics-go.v3"
	"gopkg.in/urfave/cli.v1"

	"github.com/Southclaws/sampctl/print"
	"github.com/Southclaws/sampctl/rook"
	"github.com/Southclaws/sampctl/util"
)

var packageReleaseFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "dir",
		Value: ".",
		Usage: "working directory for the project - by default, uses the current directory",
	},
}

func packageRelease(c *cli.Context) error {
	if c.Bool("verbose") {
		print.SetVerbose()
	}

	if config.Metrics {
		segment.Enqueue(analytics.Track{
			Event:  "package release",
			UserId: config.UserID,
		})
	}

	dir := util.FullPath(c.String("dir"))

	pkg, err := rook.PackageFromDir(true, dir, runtime.GOOS, "")
	if err != nil {
		return errors.Wrap(err, "failed to interpret directory as Pawn package")
	}

	err = rook.Release(context.Background(), gh, gitAuth, pkg)
	if err != nil {
		return errors.Wrap(err, "failed to release")
	}

	return nil
}
