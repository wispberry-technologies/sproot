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

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Repo Name").
				Placeholder("Enter the name of the repo").
				Value(currentRepo.Name()),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("Repo Location").
				Placeholder("Enter the location of the repo").
				Value(currentRepo.Location()),
		),
		// huh.NewGroup(
		// 	huh.NewMultiSelect[Tool]().
		// 		Title("Need any cool features setup?").
		// 		Options(
		// 			huh.NewOption("Crush", Crush),
		// 		).
		// 		Value(currentRepo.Tools()),
		// ),

		huh.NewGroup(
			huh.NewMultiSelect[Project]().
				Title("Scaffold backend project?").
				Options(
					huh.NewOption("-- skip", Skip),
					huh.NewOption("Go Chi", GoChi),
				).
				Value(currentRepo.Projects()),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Printf("Error running form: %v\n", err)
		return
	}

	fmt.Println("Repo name chosen:", *currentRepo.Name())
	fmt.Println("Repo location:", *currentRepo.Location())
	fmt.Println("Repo tools:", *currentRepo.Tools())
}
