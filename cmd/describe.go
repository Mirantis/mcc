package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/Mirantis/mcc/pkg/analytics"
	"github.com/Mirantis/mcc/pkg/config"
	"github.com/Mirantis/mcc/pkg/exec"
	"github.com/urfave/cli/v2"
	event "gopkg.in/segmentio/analytics-go.v3"

	log "github.com/sirupsen/logrus"
)

var reports = []string{"hosts", "mke", "msr", "config", "dummy"}

func reportIsKnown(n string) bool {
	for _, v := range reports {
		if v == n {
			return true
		}
	}
	return false
}

// NewDescribeCommand creates new describe command to be called from cli
func NewDescribeCommand() *cli.Command {
	return &cli.Command{
		Name:      "describe",
		Usage:     "Display cluster status",
		ArgsUsage: fmt.Sprintf("<%s>", strings.Join(reports, "|")),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "config",
				Usage:     "Path to cluster config yaml",
				Aliases:   []string{"c"},
				Value:     "launchpad.yaml",
				TakesFile: true,
			},
			&cli.BoolFlag{
				Name:  "confirm",
				Usage: "Ask confirmation for all commands",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "disable-redact",
				Usage: "Do not hide sensitive information in the output",
				Value: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			report := ctx.Args().First()
			if report == "" {
				return fmt.Errorf("missing report name argument")
			}
			if !reportIsKnown(report) {
				return fmt.Errorf("unknown report %s - must be one of %s", report, strings.Join(reports, ","))
			}

			if !(ctx.Bool("debug") || ctx.Bool("trace")) {
				log.SetLevel(log.FatalLevel)
			}

			start := time.Now()
			analytics.TrackEvent("Cluster Describe Started", nil)

			product, err := config.ProductFromFile(ctx.String("config"))
			if err != nil {
				return err
			}

			err = product.Describe(ctx.Args().First())

			if err != nil {
				analytics.TrackEvent("Cluster Describe Failed", nil)
			} else {
				duration := time.Since(start)
				props := event.Properties{
					"duration": duration.Seconds(),
				}
				analytics.TrackEvent("Cluster Describe Completed", props)
			}
			return err
		},
		Before: func(ctx *cli.Context) error {
			exec.Confirm = ctx.Bool("confirm")
			exec.DisableRedact = ctx.Bool("disable-redact")
			if !ctx.Bool("accept-license") {
				return analytics.RequireRegisteredUser()
			}
			return nil
		},
	}
}
