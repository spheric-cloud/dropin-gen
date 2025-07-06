// SPDX-FileCopyrightText: 2025 Axel Christ and Spheric contributors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bytes"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func readFile(t *testing.T, filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func assertFile(t *testing.T, filename string, expectedData []byte) {
	t.Helper()

	actualData, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Error reading file %s: %v", filename, err)
		return
	}

	actual := bytes.TrimSpace(actualData)
	expectedData = bytes.TrimSpace(expectedData)
	if !bytes.Equal(actual, expectedData) {
		t.Errorf("File %s content does not match expected data.\nExpected:\n%s\n\nActual:\n%s", filename, expectedData, actualData)
	}
}

func Test_generate(t *testing.T) {
	type args struct {
		nonTempDir     bool
		srcDir         string
		importPath     string
		outputPackage  string
		outputFilename string
		goHeaderFile   string
		requireFiles   map[string][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Generate simple dropin",
			args: args{
				srcDir:         filepath.Join("testdata", "orig"),
				importPath:     ".",
				outputPackage:  "dropin",
				outputFilename: "zz_generated.dropin.go",
				goHeaderFile:   filepath.Join("hack", "boilerplate.go.txt"),
				requireFiles: map[string][]byte{
					"zz_generated.dropin.go": readFile(t, filepath.Join("testdata", "expected", "simple-dropin.go")),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.args.nonTempDir {
				tt.args.outputPackage = filepath.Join(t.TempDir(), tt.args.outputPackage)
			}

			if err := generate(tt.args.srcDir, tt.args.importPath, tt.args.outputPackage, tt.args.outputFilename, tt.args.goHeaderFile); (err != nil) != tt.wantErr {
				t.Errorf("generate() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				requireFilenames := slices.Sorted(maps.Keys(tt.args.requireFiles))
				for _, requireFilename := range requireFilenames {
					expectedData := tt.args.requireFiles[requireFilename]
					if !tt.args.nonTempDir {
						requireFilename = filepath.Join(tt.args.outputPackage, requireFilename)
					}
					assertFile(t, requireFilename, expectedData)
				}
			}
		})
	}
}
