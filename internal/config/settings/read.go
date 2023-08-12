package settings

import (
	"io"
	"os"
	"path/filepath"

	envsubst "github.com/a8m/envsubst"
	"gopkg.in/yaml.v3"
)

const callerOffset = 1

func Read(env, path string) (*Settings, error) {
	projectDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Open file
	file, err := os.Open(filepath.Join(projectDir, path))
	if err != nil {
		return nil, err
	}

	// Read file
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse file
	confContent, err := envsubst.Bytes(bytes)
	if err != nil {
		return nil, err
	}

	settingsFile := make(map[string]*Settings)

	// Unmarshal file
	err = yaml.Unmarshal(confContent, settingsFile)
	if err != nil {
		return nil, err
	}
	return settingsFile[env], nil
}
