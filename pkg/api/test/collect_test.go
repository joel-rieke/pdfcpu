/*
Copyright 2020 The pdf Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"path/filepath"
	"testing"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func TestCollect(t *testing.T) {
	msg := "TestCollect"

	inFile := filepath.Join(inDir, "pike-stanford.pdf")
	outFile := filepath.Join(outDir, "myPageSequence.pdf")

	// Start with all odd pages but page 1, then append pages 8-11 and the last page.
	if err := api.CollectFile(inFile, outFile, []string{"odd", "!1", "8-11", "l"}, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}

	if err := api.ValidateFile(outFile, nil); err != nil {
		t.Fatalf("%s: %v\n", msg, err)
	}
}