package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// saplingDir is where we keep our config
const saplingDir = ".sapling"
const configFile = "repo.toml"

// ConfigExists checks whether /.sapling/repo.toml exists
func ConfigExists() bool {
	path := filepath.Join(saplingDir, configFile)
	_, err := os.Stat(path)
	return err == nil
}

// LoadConfig reads the TOML file into a Repo
func LoadConfig() (Repo, error) {
	path := filepath.Join(saplingDir, configFile)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("config does not exist at %s", path)
	}

	r := &repo{}
	if _, err := toml.DecodeFile(path, r); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return r, nil
}

// CreateConfig writes a Repo into /.sapling/repo.toml
func CreateConfig(repo Repo) error {
	if err := os.MkdirAll(saplingDir, 0755); err != nil {
		return fmt.Errorf("failed to create sapling dir: %w", err)
	}

	path := filepath.Join(saplingDir, configFile)

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer f.Close()

	enc := toml.NewEncoder(f)
	if err := enc.Encode(repo); err != nil {
		return fmt.Errorf("failed to encode repo config: %w", err)
	}

	return nil
}
