# snap plugin collector - swap

snap plugin for collecting swap metrics from /proc linux filesystem

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

 Plugin collects specified metrics in-band on OS level

### System Requirements
 - Linux system

### Installation
#### Download the plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [Github Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-collector-swap

Clone repo into `$GOPATH/src/github/intelsdi-x/`:
```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-swap
```
Build the plugin by running make in repo:
```
$ make
```
This builds the plugin in `/build/rootfs`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started).
* Load the plugin and create a task, see example in [Examples](https://github.com/intelsdi-x/snap-plugin-collector-swap/blob/master/README.md#examples).

## Documentation

### Collected Metrics
The path to the procfs can be provided in configuration as `proc_path`. If configuration is not provided, the plugin will use the default of `/proc`.

It can be set in the snap global config that is loaded with snapd, e.g.:
```json
{
    "log_level": 1,
    "control": {
        "plugin_trust_level": 0,
        "plugins": {
            "collector": {
				"swap": {
					"all": {
						"proc_path": "/hostproc"
					}
				}
            }
        }
    }
}
```

The list of collected metrics is described in [METRICS.md](https://github.com/intelsdi-x/snap-plugin-collector-swap/blob/master/METRICS.md).

### Examples

Example running snap-plugin-collector-swap plugin and writing data to a file.

Make sure that your `$SNAP_PATH` is set, if not:
```
$ export SNAP_PATH=<snapDirectoryPath>/build
```

Other paths to files should be set according to your configuration, using a file you should indicate where it is located.

In one terminal window, open the snap daemon (in this case with logging set to 1,  trust disabled):
```
$ $SNAP_PATH/bin/snapd -l 1 -t 0
```
In another terminal window:

Load snap-plugin-collector-swap plugin:
```
$ $SNAP_PATH/bin/snapctl plugin load snap-plugin-collector-swap
```
Load file plugin for publishing:
```
$ $SNAP_PATH/bin/snapctl plugin load $SNAP_PATH/plugin/snap-publisher-file
```
See available metrics for your system:
```
$ $SNAP_PATH/bin/snapctl metric list
```

Create a task manifest file to use snap-plugin-collector-swap plugin (exemplary files in [examples/tasks/] (https://github.com/intelsdi-x/snap-plugin-collector-swap/blob/master/examples/tasks/)):
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "5s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/procfs/swap/device/dev_sda5/used_bytes": {},
                "/intel/procfs/swap/device/dev_sda5/used_percent": {},
                "/intel/procfs/swap/device/dev_sda5/free_bytes": {},
                "/intel/procfs/swap/device/dev_sda5/free_percent": {},
                "/intel/procfs/swap/all/cached_bytes": {},
                "/intel/procfs/swap/all/cached_percent":{},
                "/intel/procfs/swap/all/free_bytes":{},
                "/intel/procfs/swap/all/free_percent":{},
                "/intel/procfs/swap/all/used_bytes":{},
                "/intel/procfs/swap/all/used_percent":{},
                "/intel/procfs/swap/io/in_bytes_per_sec":{},
                "/intel/procfs/swap/io/in_pages_per_sec":{},
                "/intel/procfs/swap/io/out_bytes_per_sec":{},
                "/intel/procfs/swap/io/out_pages_per_sec":{}
            },
            "config": {},
            "process": null,
            "publish": [
                {
                    "plugin_name": "file",
                    "config": {
                        "file": "/tmp/published_swap"
                    }
                }
            ]
        }
    }
}
```
Create a task:
```
$ $SNAP_PATH/bin/snapctl task create -t swap-file.json
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-swap/issues) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-swap/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support) or visit [snap Gitter channel](https://gitter.im/intelsdi-x/snap).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

And **thank you!** Your contribution, through code and participation, is incredibly important to us.

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [Andrzej Kuriata](https://github.com/andrzej-k)