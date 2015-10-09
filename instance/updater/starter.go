package updater

import (
	bosherr "github.com/cloudfoundry/bosh-agent/errors"
	boshlog "github.com/cloudfoundry/bosh-agent/logger"

	bpagclient "github.com/sclevine/bosh-provisioner/agent/client"
)

const starterLogTag = "Starter"

type Starter struct {
	agentClient bpagclient.Client
	logger      boshlog.Logger
}

func NewStarter(
	agentClient bpagclient.Client,
	logger boshlog.Logger,
) Starter {
	return Starter{
		agentClient: agentClient,
		logger:      logger,
	}
}

func (s Starter) Start() error {
	s.logger.Debug(starterLogTag, "Starting instance")

	_, err := s.agentClient.Start()
	if err != nil {
		return bosherr.WrapError(err, "Starting")
	}

	return nil
}
