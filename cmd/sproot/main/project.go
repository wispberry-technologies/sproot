package main

import "github.com/charmbracelet/huh"

// ---------------
// Project
// ---------------

type project struct {
	name     string
	slug     string
	features map[string]Feature // slug -> Feature
	setup    func(...string) error
}

type Project interface {
	Name() *string
	Slug() *string
	Setup(...string) error
	Features() map[string]Feature // slug -> Feature
	AddFeature(f Feature) error
}

func NewProject() Project {
	return &project{}
}

func (p *project) Name() *string {
	return &p.name
}
func (p *project) Slug() *string {
	return &p.slug
}
func (p *project) Features() map[string]Feature {
	return p.features
}
func (p *project) AddFeature(f Feature) error {
	if p.features == nil {
		p.features = make(map[string]Feature)
	}
	p.features[*f.Name()] = f
	return nil
}
func (p *project) Setup(args ...string) error {
	return p.setup(args...)
}

// ---------------
// Project -> Features
// ---------------
type feature struct {
	name        string
	description string
	enabled     bool
	order       int
	setup       func() error
}

type Feature interface {
	Name() *string
	Description() *string
	Enabled() *bool
	Order() *int
	Setup() error
}

func (f *feature) Name() *string {
	return &f.name
}
func (f *feature) Description() *string {
	return &f.description
}
func (f *feature) Enabled() *bool {
	return &f.enabled
}
func (f *feature) Order() *int {
	return &f.order
}
func (f *feature) Setup() error {
	return f.setup()
}

// Project setup/options form
func ProjectSetupForm(project Project) *huh.Group {
	return huh.NewGroup(
		huh.NewInput().
		
	)
}
