/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015-2016 Intel Corporation

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

package main

import (
	"os"
	"testing"

	"github.com/intelsdi-x/snap-plugin-collector-swap/swap"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	ioNewMockFile  = "/tmp/vmstat_test"
	ioOldMockFile  = "/tmp/stat_test"
	perDevMockFile = "/tmp/swaps_test"
	compMockFile   = "/tmp/meminfo_test"
)

func TestMain(t *testing.T) {
	createMockFiles()
	Convey("ensure plugin loads and responds if data sources available", t, func() {
		swap.SourceIOnew = "/tmp/nonexistingfile"
		swap.SourceIOold = "/tmp/nonexistinigfile"
		swap.SourcePerDev = "/tmp/nonexistinigfile"
		swap.SourceCombined = "/tmp/nonexistinigfile"
		os.Args = []string{"", "{\"NoDaemon\": true}"}
		So(func() { main() }, ShouldPanic)
	})
	Convey("ensure plugin fails to load if data sources unavailable", t, func() {
		swap.SourceIOnew = ioNewMockFile
		swap.SourceIOold = ioOldMockFile
		swap.SourcePerDev = perDevMockFile
		swap.SourceCombined = compMockFile
		os.Args = []string{"", "{\"NoDaemon\": true}"}
		So(func() { main() }, ShouldNotPanic)
	})
	deleteMockFiles()
}

func createMockFiles() {
	deleteMockFiles()
	ioNewMockFileCont := []byte("This is new IO mock\n")
	ioOldMockFileCont := []byte("This is old IO mock\n")
	perDevMockFileCont := []byte("This is dev mock\n")
	compMockFileCont := []byte("This is comp mock\n")

	f, _ := os.Create(ioNewMockFile)
	f.Write(ioNewMockFileCont)
	f, _ = os.Create(ioOldMockFile)
	f.Write(ioOldMockFileCont)
	f, _ = os.Create(perDevMockFile)
	f.Write(perDevMockFileCont)
	f, _ = os.Create(compMockFile)
	f.Write(compMockFileCont)
}

func deleteMockFiles() {
	os.Remove(ioNewMockFile)
	os.Remove(ioOldMockFile)
	os.Remove(perDevMockFile)
	os.Remove(compMockFile)
}
