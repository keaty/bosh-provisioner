package provisioner

import (
	bosherr "github.com/cloudfoundry/bosh-agent/errors"
	boshlog "github.com/cloudfoundry/bosh-agent/logger"

	bpeventlog "github.com/sclevine/bosh-provisioner/eventlog"
	bpvm "github.com/sclevine/bosh-provisioner/vm"
)

// SingleNonConfiguredVMProvisioner configures 1 VM as a regular empty BOSH VM.
type SingleNonConfiguredVMProvisioner struct {
	vmProvisioner bpvm.Provisioner
	eventLog      bpeventlog.Log
	logger        boshlog.Logger
}

func NewSingleNonConfiguredVMProvisioner(
	vmProvisioner bpvm.Provisioner,
	eventLog bpeventlog.Log,
	logger boshlog.Logger,
) SingleNonConfiguredVMProvisioner {
	return SingleNonConfiguredVMProvisioner{
		vmProvisioner: vmProvisioner,
		eventLog:      eventLog,
		logger:        logger,
	}
}

func (p SingleNonConfiguredVMProvisioner) Provision() error {
	// todo VM was possibly provisioned last time
	_, err := p.vmProvisioner.ProvisionNonConfigured()
	if err != nil {
		return bosherr.WrapError(err, "Provisioning VM")
	}

	// Do not Deprovision() VM to keep instance running

	return nil
}
