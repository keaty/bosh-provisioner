package updater

import (
	"fmt"
	"time"

	boshlog "github.com/cloudfoundry/bosh-agent/logger"

	bpagclient "github.com/sclevine/bosh-provisioner/agent/client"
	bpdep "github.com/sclevine/bosh-provisioner/deployment"
	bpeventlog "github.com/sclevine/bosh-provisioner/eventlog"
	bptplcomp "github.com/sclevine/bosh-provisioner/instance/templatescompiler"
	bpapplier "github.com/sclevine/bosh-provisioner/instance/updater/applier"
	bppkgscomp "github.com/sclevine/bosh-provisioner/packagescompiler"
)

type Factory struct {
	templatesCompiler       bptplcomp.TemplatesCompiler
	packagesCompilerFactory bppkgscomp.ConcretePackagesCompilerFactory

	eventLog bpeventlog.Log
	logger   boshlog.Logger
}

func NewFactory(
	templatesCompiler bptplcomp.TemplatesCompiler,
	packagesCompilerFactory bppkgscomp.ConcretePackagesCompilerFactory,
	eventLog bpeventlog.Log,
	logger boshlog.Logger,
) Factory {
	return Factory{
		templatesCompiler:       templatesCompiler,
		packagesCompilerFactory: packagesCompilerFactory,

		eventLog: eventLog,
		logger:   logger,
	}
}

func (f Factory) NewUpdater(
	agentClient bpagclient.Client,
	depJob bpdep.Job,
	instance bpdep.Instance,
) Updater {
	drainer := NewDrainer(agentClient, f.logger)

	stopper := NewStopper(agentClient, f.logger)

	applier := bpapplier.NewApplier(
		depJob,
		instance,
		f.templatesCompiler,
		f.packagesCompilerFactory.NewCompiler(agentClient),
		agentClient,
		f.logger,
	)

	starter := NewStarter(agentClient, f.logger)

	waiter := NewWaiter(
		instance.WatchTime.Start(),
		instance.WatchTime.End(),
		time.Sleep,
		agentClient,
		f.logger,
	)

	updater := NewUpdater(
		fmt.Sprintf("%s/%d", instance.JobName, instance.Index),
		drainer,
		stopper,
		applier,
		starter,
		waiter,
		f.eventLog,
		f.logger,
	)

	return updater
}
