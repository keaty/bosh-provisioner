package fakes

import (
	boshaction "github.com/cloudfoundry/bosh-agent/agent/action"
	boshas "github.com/cloudfoundry/bosh-agent/agent/applier/applyspec"
	boshcomp "github.com/cloudfoundry/bosh-agent/agent/compiler"
	bosherr "github.com/cloudfoundry/bosh-agent/errors"

	bpagclient "github.com/sclevine/bosh-provisioner/agent/client"
)

type FakeClient struct {
	GetStateState  boshaction.GetStateV1ApplySpec
	GetStateStates []boshaction.GetStateV1ApplySpec
	GetStateErr    error
}

func (c *FakeClient) Ping() (string, error) {
	return "", bosherr.New("fake-ping-err")
}

func (c *FakeClient) GetTask(string) (interface{}, error) {
	return "", bosherr.New("fake-get-task-err")
}

func (c *FakeClient) CancelTask(string) (string, error) {
	return "", bosherr.New("fake-cancel-task-err")
}

func (c *FakeClient) SSH(cmd string, params boshaction.SSHParams) (map[string]interface{}, error) {
	return nil, bosherr.New("fake-ssh-err")
}

func (c *FakeClient) FetchLogs(logType string, filters []string) (map[string]interface{}, error) {
	return nil, bosherr.New("fake-fetch-logs-err")
}

func (c *FakeClient) Prepare(boshas.V1ApplySpec) (string, error) {
	return "", bosherr.New("fake-prepare-err")
}

func (c *FakeClient) Apply(boshas.V1ApplySpec) (string, error) {
	return "", bosherr.New("fake-apply-err")
}

func (c *FakeClient) GetState(filters ...string) (boshaction.GetStateV1ApplySpec, error) {
	state := c.GetStateState

	if c.GetStateStates != nil {
		state = c.GetStateStates[0]
		c.GetStateStates = c.GetStateStates[1:]
	}

	return state, c.GetStateErr
}

func (c *FakeClient) Start() (string, error) {
	return "", bosherr.New("fake-start-err")
}

func (c *FakeClient) Stop() (string, error) {
	return "", bosherr.New("fake-stop-err")
}

func (c *FakeClient) Drain(boshaction.DrainType, ...boshas.V1ApplySpec) (int, error) {
	return 0, bosherr.New("fake-drain-err")
}

func (c *FakeClient) RunErrand() (boshaction.ErrandResult, error) {
	return boshaction.ErrandResult{}, bosherr.New("fake-run-errand-err")
}

func (c *FakeClient) CompilePackage(
	blobID string,
	sha1 string,
	name string,
	version string,
	deps boshcomp.Dependencies,
) (bpagclient.CompiledPackage, error) {
	return bpagclient.CompiledPackage{}, bosherr.New("fake-ping-err")
}
