package util

import (
	"encoding/json"
	"errors"
	"os"

	"gosync/internals/data"
	"gosync/internals/errors/jsonerrors"
)

func ReadJSONManifest(path string) (*data.Manifest, error) {
	if path == "" {
		return nil, &jsonerrors.JSONReadError{Err: errors.New("JSON Manifest path is not specified.")}
	}

	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, &jsonerrors.JSONReadError{Err: err}
	}
	var manifestData data.Manifest
	json.Unmarshal(jsonBytes, &manifestData)

	return &manifestData, nil
}
