package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// -----------------
// Repo + Tools
// -----------------

type repo struct {
	name     string
	location string
	tools    []Tool
	projects []Project
}

type Repo interface {
	Name() *string
	Location() *string
	Tools() *[]Tool
	Projects() *[]Project
}

func NewRepo() Repo {
	return &repo{
		name:     "",
		location: "./",
		tools:    []Tool{},
	}
}

func (r *repo) Name() *string {
	return &r.name
}

func (r *repo) Location() *string {
	return &r.location
}

func (r *repo) Tools() *[]Tool {
	return &r.tools
}

func (r *repo) Projects() *[]Project {
	return &r.projects
}

// -----------------
// Repo CLI Setup
// -----------------

func CreateProject() {

}

func SetupRepoCli() {
	currentRepo := NewRepo()

	initForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Repo Name").
				Placeholder("Enter the name of the repo").
				Value(currentRepo.Name()),
		),
		//
		huh.NewGroup(
			huh.NewInput().
				Title("Repo Location").
				Placeholder("Enter the location of the repo").
				Value(currentRepo.Location()),
		),
		//
		huh.NewGroup(
			huh.NewMultiSelect[Project]().
				Title("Let's scaffold some projects!").
				Options(
					huh.NewOption("Go Chi", GoChi),
					huh.NewOption("Svelte Kit", SvelteKit),
				).
				Value(currentRepo.Projects()),
		),
		//
	)

	if err := initForm.Run(); err != nil {
		fmt.Printf("Error running form: %v\n", err)
		return
	}

	projectFormGroups := []*huh.Group{}
	for _, project := range *currentRepo.Projects() {
		projectFormGroups = append(projectFormGroups, ProjectSetupForm(project))
	}

	projectsForm := huh.NewForm()

	if err := projectsForm.Run(); err != nil {
		fmt.Printf("Error running projects form: %v\n", err)
		return
	}

	fmt.Println("Repo name chosen:", *currentRepo.Name())
	fmt.Println("Repo location:", *currentRepo.Location())
	fmt.Println("Repo tools:", *currentRepo.Tools())
	fmt.Println("Repo projects:", *currentRepo.Projects())
}
