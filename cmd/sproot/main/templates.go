package main

import "fmt"

// Nil / skip
var Skip Project = nil

// Go-Chai
var GoChi Project = &project{
	name: "Go-Chi",
	slug: "go-chi",
	setup: func(args ...string) error {
		fmt.Println("Setting up Go-Chi project...")
		// Add setup logic here
		return nil
	},
}

// SvelteKit
var SvelteKit Project = &project{
	name: "Svelte-Kit",
	slug: "sveltekit",
	setup: func(args ...string) error {
		fmt.Println("Setting up Svelte-Kit project...")
		// Add setup logic here
		return nil
	},
}
