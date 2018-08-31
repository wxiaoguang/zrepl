package client

import (
	"github.com/pkg/errors"
	"github.com/zrepl/zrepl/config"
	"github.com/zrepl/zrepl/daemon"
)

func RunWakeup(config *config.Config, args []string) error {
	if len(args) != 1 {
		return errors.Errorf("Expected 1 argument: job")
	}

	httpc, err := controlHttpClient(config.Global.Control.SockPath)
	if err != nil {
		return err
	}

	err = jsonRequestResponse(httpc, daemon.ControlJobEndpointWakeup,
		struct {
			Name string
		}{
			Name: args[0],
		},
		struct{}{},
	)
	return err
}
