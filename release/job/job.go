package job

import (
	bpdep "github.com/sclevine/bosh-provisioner/deployment"
	bpreljobman "github.com/sclevine/bosh-provisioner/release/job/manifest"
)

type Job struct {
	Manifest bpreljobman.Manifest `json:"-"`

	Name        string
	Description string

	MonitTemplate Template

	Templates []Template

	DeploymentJobTemplates []bpdep.Template

	// Runtime package dependencies for this job
	Packages []Package

	Properties []Property
}

type Template struct {
	SrcPathEnd string
	DstPathEnd string // End of the path on the VM

	Path string
}

type Package struct {
	Name string
}

type Property struct {
	Name        string
	Description string

	Default interface{}
	Example interface{}
}

// populateFromManifest populates job information interpreted from job manifest.
func (j *Job) populateFromManifest(manifest bpreljobman.Manifest) {
	j.populateJob(manifest.Job)
	j.populateTemplates(manifest.Job.TemplateNames)
	j.populatePackages(manifest.Job.PackageNames)
	j.populateProperties(manifest.Job.PropertyMappings)
	j.Manifest = manifest
}

func (j *Job) populateJob(manJob bpreljobman.Job) {
	j.Name = manJob.Name
	j.Description = manJob.Description
}

func (j *Job) populateTemplates(manTemplateNames bpreljobman.TemplateNames) {
	j.MonitTemplate = Template{
		SrcPathEnd: "monit",
		DstPathEnd: "monit",
	}

	for srcPathEnd, dstPathEnd := range manTemplateNames {
		template := Template{
			SrcPathEnd: srcPathEnd,
			DstPathEnd: dstPathEnd,
		}

		j.Templates = append(j.Templates, template)
	}
}

func (j *Job) populatePackages(manPackageNames []string) {
	for _, name := range manPackageNames {
		j.Packages = append(j.Packages, Package{Name: name})
	}
}

func (j *Job) populateProperties(manPropMappings bpreljobman.PropertyMappings) {
	for propName, propDef := range manPropMappings {
		property := Property{
			Name:        propName,
			Description: propDef.Description,

			Default: propDef.Default,
			Example: propDef.Example,
		}

		j.Properties = append(j.Properties, property)
	}
}
