package main

import (
	"github.com/aquasecurity/trivy/pkg/module/api"
	"github.com/aquasecurity/trivy/pkg/module/serialize"
	"github.com/aquasecurity/trivy/pkg/module/wasm"
)

const (
	moduleName    = "Custom Trivy Module"
	moduleVersion = 1
)

func main() {
	wasm.RegisterModule(TrivyModule{})
}

type TrivyModule struct{}

func (TrivyModule) Name() string {
	return moduleName
}
func (TrivyModule) Version() int {
	return moduleVersion
}

func (TrivyModule) Analyze(filePath string) (*serialize.AnalysisResult, error) {
	wasm.Warn("Hello from Analyze")
	return nil, nil
}
func (TrivyModule) RequiredFiles() []string {
	wasm.Warn("Hello from RequiredFiles")
	return nil
}

func (TrivyModule) PostScanSpec() serialize.PostScanSpec {
	wasm.Warn("Hello from PostScanSpec")
	return serialize.PostScanSpec{
		Action: api.ActionUpdate,
	}
}
func (TrivyModule) PostScan(result serialize.Results) (serialize.Results, error) {
	wasm.Warn("Hello from PostScan")
	return result, nil
}
