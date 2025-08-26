package main

import "fmt"

// Nil / skip
var Skip Project = nil

// Go-Chai
var GoChai Project = &project{
	name: "Go-Chai",
	slug: "go-chai",
	setup: func(args ...string) error {
		fmt.Println("Setting up Go-Chai project...")
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
