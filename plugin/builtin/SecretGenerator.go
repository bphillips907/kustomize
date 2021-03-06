// Code generated by pluginator on SecretGenerator; DO NOT EDIT.
package builtin

import (
	"sigs.k8s.io/kustomize/pkg/ifc"
	"sigs.k8s.io/kustomize/pkg/resmap"
	"sigs.k8s.io/kustomize/pkg/types"
	"sigs.k8s.io/yaml"
)

type SecretGeneratorPlugin struct {
	ldr ifc.Loader
	rf  *resmap.Factory
	types.GeneratorOptions
	types.SecretArgs
}

func NewSecretGeneratorPlugin() *SecretGeneratorPlugin {
	return &SecretGeneratorPlugin{}
}

func (p *SecretGeneratorPlugin) Config(
	ldr ifc.Loader, rf *resmap.Factory, config []byte) (err error) {
	p.GeneratorOptions = types.GeneratorOptions{}
	p.SecretArgs = types.SecretArgs{}
	err = yaml.Unmarshal(config, p)
	p.ldr = ldr
	p.rf = rf
	return
}

func (p *SecretGeneratorPlugin) Generate() (resmap.ResMap, error) {
	argsList := make([]types.SecretArgs, 1)
	argsList[0] = p.SecretArgs
	return p.rf.NewResMapFromSecretArgs(
		p.ldr, &p.GeneratorOptions, argsList)
}
