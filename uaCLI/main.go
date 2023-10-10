package main

import (
	"github.com/eyethereal/go-archercl"
	"github.com/op/go-logging"
	"github.com/tomseago/uArm/io"
)

var log = logging.MustGetLogger("uaCLI")

func main() {
	// Phase 0 of the boot sequence: Configuration

	opts := &archercl.Opts{
		Name: "uaCLI",
		DefaultText: `
		logging level: debug
/*
		logging backends {
			ui {
				type: delayed
				format: "%{time:15:04} %{level} %{message}"
			}

			syslog {
				type: syslog
				facility: user
				prefix: bchat
				level: debug
			}
		}
*/

		dumpConfig: true
		dumpColor: false
		`,

		// Set this so we can see early log messages, but not desirable
		// in most deployments
		AddColorConsoleLogging: false,
	}

	cfg, err := archercl.Load(opts)
	if err != nil {
		panic(err)
	}

	if cfg.ChildAsBool("configDebug") {
		log.Error("Stopping because configDebug was set")
		return
	}

	log.Debug("Starting up")

	sl := io.NewSlinger()

	log.Infof("Hello nextIndex=%v\n", sl.NextIndex())
}
