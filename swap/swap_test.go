// +build linux

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

package swap

import (
	"fmt"
	"os"
	"testing"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/core"
	"github.com/intelsdi-x/snap/core/cdata"
	"github.com/intelsdi-x/snap/core/ctypes"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	mockMts = []plugin.MetricType{
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "io", "in_bytes_per_sec"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "io", "in_pages_per_sec"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "io", "out_bytes_per_sec"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "io", "out_pages_per_sec"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "device", "*", "used_bytes"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "device", "*", "used_percent"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "device", "*", "free_bytes"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "device", "*", "free_percent"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "all", "used_bytes"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "all", "used_percent"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "all", "free_bytes"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "all", "free_percent"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "all", "cached_bytes"),
		},
		plugin.MetricType{
			Namespace_: core.NewNamespace("intel", "procfs", "swap", "all", "cached_percent"),
		},
	}
	ioNewMockFile  = "/tmp/vmstat"
	ioOldMockFile  = "/tmp/stat"
	perDevMockFile = "/tmp/swaps"
	compMockFile   = "/tmp/meminfo"
)

func TestGetConfigPolicy(t *testing.T) {
	SourceIOnew = ioNewMockFile
	SourceIOold = ioOldMockFile
	SourcePerDev = perDevMockFile
	SourceCombined = compMockFile
	createMockFiles()
	swap := NewSwapCollector()
	Convey("normal case", t, func() {
		So(func() { swap.GetConfigPolicy() }, ShouldNotPanic)
		_, err := swap.GetConfigPolicy()
		So(err, ShouldBeNil)
	})
	deleteMockFiles()
}

func TestGetMetricTypes(t *testing.T) {
	SourceIOnew = ioNewMockFile
	SourceIOold = ioOldMockFile
	SourcePerDev = perDevMockFile
	SourceCombined = compMockFile
	createMockFiles()
	swap := NewSwapCollector()
	cfg := plugin.NewPluginConfigType()
	cfg.AddItem(ProcPathCfg, ctypes.ConfigValueStr{Value: "/dummy"})
	Convey("proc_path does not exist", t, func() {
		m, err := swap.GetMetricTypes(cfg)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "no such file or directory")
		So(m, ShouldBeNil)
	})
	cfg = plugin.NewPluginConfigType()
	cfg.AddItem(ProcPathCfg, ctypes.ConfigValueStr{Value: "/tmp"})
	Convey("source files available", t, func() {
		swap.initialized = false
		m, err := swap.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		// 4 - IO metrics, 4 - dev metrics, 6 - combined metrics
		So(m, ShouldNotBeNil)
		So(len(m), ShouldEqual, 14)
	})
	Convey("Dummy IO new source file, should switch to old mode", t, func() {
		SourceIOnew = ioNewMockFile + "-dummy"
		m, err := swap.GetMetricTypes(cfg)
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)
		So(len(m), ShouldEqual, 14)
	})
	Convey("Dummy IO new+old source files, should switch to old mode and fail", t, func() {
		SourceIOnew = ioNewMockFile + "-dummy"
		SourceIOold = ioOldMockFile + "-dummy"
		m, err := swap.GetMetricTypes(cfg)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "Data source")
		So(err.Error(), ShouldContainSubstring, "not accessible")
		So(m, ShouldBeNil)
	})
	Convey("dev source file not available", t, func() {
		SourceIOnew = ioNewMockFile + "-dummy"
		SourceIOold = ioOldMockFile
		os.Remove(perDevMockFile)
		m, err := swap.GetMetricTypes(cfg)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "Data source")
		So(err.Error(), ShouldContainSubstring, "not accessible")
		So(m, ShouldBeNil)
	})
	deleteMockFiles()
}

func TestCollectMetrics(t *testing.T) {
	SourceIOnew = ioNewMockFile
	SourceIOold = ioOldMockFile
	SourcePerDev = perDevMockFile
	SourceCombined = compMockFile
	createMockFiles()
	swap := NewSwapCollector()
	Convey("proc_path does not exist", t, func() {
		node := cdata.NewNode()
		node.AddItem(ProcPathCfg, ctypes.ConfigValueStr{Value: "/dummy"})
		mts := []plugin.MetricType{
			plugin.MetricType{
				Namespace_: core.NewNamespace("intel", "procfs", "swap", "io", "in_bytes_per_sec"),
				Config_:    node,
			},
		}
		m, err := swap.CollectMetrics(mts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
	})
	swap = NewSwapCollector()
	Convey("source files available", t, func() {
		m, err := swap.CollectMetrics(mockMts)
		So(err, ShouldBeNil)
		// 4 - IO metrics, 8 - dev metrics (2 devices), 6 - combined metrics
		So(len(m), ShouldEqual, 18)
	})
	Convey("source files available old IO mode", t, func() {
		swap.newIOfile = false
		m, err := swap.CollectMetrics(mockMts)
		So(err, ShouldBeNil)
		// 4 - IO metrics, 8 - dev metrics (2 devices), 6 - combined metrics
		So(len(m), ShouldEqual, 18)
	})
	Convey("metric exists", t, func() {
		mts := []plugin.MetricType{
			plugin.MetricType{
				Namespace_: core.NewNamespace("intel", "procfs", "swap", "device", "dev_sda5", "free_bytes"),
			},
		}
		m, err := swap.CollectMetrics(mts)
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)
		So(len(m), ShouldEqual, 1)
	})
	Convey("metrics do not exist", t, func() {
		mts := []plugin.MetricType{
			plugin.MetricType{
				Namespace_: core.NewNamespace("intel", "procfs", "swap", "io", "dummy"),
			},
		}
		m, err := swap.CollectMetrics(mts)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "Requested IO swap stat")
		So(err.Error(), ShouldContainSubstring, "is not available!")
		So(m, ShouldNotBeNil)
		So(len(m), ShouldEqual, 0)
		mts = []plugin.MetricType{
			plugin.MetricType{
				Namespace_: core.NewNamespace("intel", "procfs", "swap", "device", "dummy", "free_bytes"),
			},
		}
		m, err = swap.CollectMetrics(mts)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "Requested per device swap stat")
		So(err.Error(), ShouldContainSubstring, "is not available!")
		So(m, ShouldNotBeNil)
		So(len(m), ShouldEqual, 0)
		mts = []plugin.MetricType{
			plugin.MetricType{
				Namespace_: core.NewNamespace("intel", "procfs", "swap", "all", "dummy"),
			},
		}
		m, err = swap.CollectMetrics(mts)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "Requested combined swap stat")
		So(err.Error(), ShouldContainSubstring, "is not available!")
		So(m, ShouldNotBeNil)
		So(len(m), ShouldEqual, 0)
	})
	Convey("dev source file not available", t, func() {
		os.Remove(perDevMockFile)
		m, err := swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
	})
	deleteMockFiles()
	SourceIOnew = ioNewMockFile
	SourceIOold = ioOldMockFile
	SourcePerDev = perDevMockFile
	SourceCombined = compMockFile
	Convey("source files available with errors or specific cases", t, func() {
		createMockFilesWithErrors(
			"not-an-int", "1010", "2020",
			"11111", "22222",
			"33333", "44444",
			"55555", "6666")
		swap = NewSwapCollector()
		m, err := swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "SwapTotal is not a number")

		createMockFilesWithErrors(
			"99999", "not-an-int", "2020",
			"11111", "22222",
			"33333", "44444",
			"55555", "6666")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "SwapFree is not a number")

		createMockFilesWithErrors(
			"99999", "1010", "not-an-int",
			"11111", "22222",
			"33333", "44444",
			"55555", "6666")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "SwapCached is not a number")

		createMockFilesWithErrors(
			"99999", "1010", "2020",
			"not-an-int", "22222",
			"33333", "44444",
			"55555", "6666")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "pswpin is not a number")

		createMockFilesWithErrors(
			"99999", "1010", "2020",
			"11111", "not-an-int",
			"33333", "44444",
			"55555", "6666")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "pswpout is not a number")

		createMockFilesWithErrors(
			"99999", "1010", "2020",
			"11111", "22222",
			"33333", "44444",
			"not-an-int", "6666")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Swap size for")
		So(err.Error(), ShouldContainSubstring, "is not a number")

		createMockFilesWithErrors(
			"99999", "1010", "2020",
			"11111", "22222",
			"33333", "44444",
			"55555", "not-an-int")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Used swap size for")
		So(err.Error(), ShouldContainSubstring, "is not a number")

		createMockFilesWithErrors(
			"0", "1010", "2020",
			"11111", "22222",
			"33333", "44444",
			"55555", "6666")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldBeNil)
		So(m, ShouldNotBeNil)
		// 4 - IO metrics, 8 - dev metrics (2 devices), 6 - combined metrics
		So(len(m), ShouldEqual, 18)

		swap.newIOfile = false
		createMockFilesWithErrors(
			"99999", "1010", "2020",
			"11111", "22222",
			"not-an-int", "44444",
			"55555", "6666")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Swap in metric is not a number")

		createMockFilesWithErrors(
			"99999", "1010", "2020",
			"11111", "22222",
			"33333", "not-an-int",
			"55555", "6666")
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Swap out metric is not a number")

		swap.newIOfile = true
		err = os.Chmod(SourceCombined, 0)
		So(err, ShouldBeNil)
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Failed to open following file for reading")

		err = os.Chmod(SourceIOnew, 0)
		So(err, ShouldBeNil)
		m, err = swap.CollectMetrics(mockMts)
		So(err, ShouldNotBeNil)
		So(m, ShouldBeNil)
		So(err.Error(), ShouldContainSubstring, "Failed to open following file for reading")
	})
	deleteMockFiles()
}

func TestHelperRoutines(t *testing.T) {
	SourceIOnew = ioNewMockFile
	SourceIOold = ioOldMockFile
	SourcePerDev = perDevMockFile
	SourceCombined = compMockFile
	createMockFiles()
	Convey("Helper Routines", t, func() {
		m := Meta()

		Convey("Then meta value should be reported as not nil", func() {
			So(m, ShouldNotBeNil)
		})

		Convey("Set config variables", func() {
			swap := NewSwapCollector()
			Convey("Swap collector should not be nil", func() {
				So(swap, ShouldNotBeNil)
				cfg := plugin.NewPluginConfigType()
				cfg.AddItem(ProcPathCfg, ctypes.ConfigValueStr{Value: "/dummy"})
				err := swap.setProcPath(cfg)
				Convey("Then error should be reported (no such file or directory)", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldContainSubstring, "no such file or directory")
				})
				cfg = plugin.NewPluginConfigType()
				cfg.AddItem(ProcPathCfg, ctypes.ConfigValueStr{Value: "/etc/hosts"})
				err = swap.setProcPath(cfg)
				Convey("Then error should be reported (not a directory)", func() {
					So(err, ShouldNotBeNil)
					So(err.Error(), ShouldContainSubstring, "is not a directory")
				})
				SourceIOnew = ioNewMockFile + "-dummy"
				SourceIOold = ioOldMockFile + "-dummy"
				swap := NewSwapCollector()
				Convey("Then error should be reported", func() {
					So(swap, ShouldBeNil)
				})
				SourceIOnew = ioNewMockFile
				SourceIOold = ioOldMockFile
			})
		})

		f := calcPercentage(1.0, 0)

		Convey("Then returned value should be 0", func() {
			So(f, ShouldEqual, 0)
		})
	})
	deleteMockFiles()
}

func createMockFiles() {
	// Proper fields
	createMockFilesWithErrors(
		"99999", "1010", "2020",
		"11111", "22222",
		"33333", "44444",
		"55555", "6666",
	)
}

func createMockFilesWithErrors(
	swapTotal string, swapFree string, swapCached string,
	swapIn string, swapOut string,
	pagesIn string, pagesOut string,
	swapSize string, usedSwapSize string,
) {
	deleteMockFiles()
	ioNewMockFileCont := []byte(
		fmt.Sprintf(
			"pswpin %s\npswpout %s\nbadentry\n",
			swapIn, swapOut))
	ioOldMockFileCont := []byte(
		fmt.Sprintf("page %s %s\nbadentry\n",
			pagesIn, pagesOut))
	perDevMockFileCont := []byte(
		fmt.Sprintf(
			"Filename Type Size Used Priority\n/dev/sda5 partition %s %s -1\n"+
				"/dev/sda6 partition  77777 8888   -1\nbadentry\n",
			swapSize, usedSwapSize))
	compMockFileCont := []byte(
		fmt.Sprintf(
			"SwapTotal: %s kB\nSwapFree: %s kB\nSwapCached: %s kB\nbad-entry\n",
			swapTotal, swapFree, swapCached))
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
