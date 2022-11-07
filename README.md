DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.

# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


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
You can get the pre-built binaries for your OS and architecture at plugin's [Github Releases](https://github.com/intelsdi-x/snap-plugin-collector-swap/releases) page.

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
This builds the plugin in `./build/`

### Configuration and Usage

The path to the procfs can be provided in configuration as `proc_path`. If configuration is not provided, the plugin will use the default of `/proc`.

It can be set in the Snap global config that is loaded with snapteld, e.g.:
```json
{
    "log_level": 1,
    "control": {
        "plugin_trust_level": 0,
        "plugins": {
            "collector": {
                "swap": {
                    "all": {
                        "proc_path": "/proc"
                    }
                }
            }
        }
    }
}
```
## Documentation

### Collected Metrics


The list of collected metrics is described in [METRICS.md](https://github.com/intelsdi-x/snap-plugin-collector-swap/blob/master/METRICS.md).

### Examples

Example of running Snap swap collector and writing data to file.

Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap):
* initd: `sudo service snap-telemetry start`
* systemd: `sudo systemctl start snap-telemetry`
* command line: `sudo snapteld -l 1 -t 0 &`

Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-swap/latest/linux/x86_64/snap-plugin-collector-swap
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ snaptel plugin load snap-plugin-collector-swap
$ snaptel plugin load snap-plugin-publisher-file
```

See available metrics for your system:
```
$ snaptel metric list
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
                "/intel/procfs/swap/device/*/used_bytes": {},
                "/intel/procfs/swap/device/*/used_percent": {},
                "/intel/procfs/swap/device/*/free_bytes": {},
                "/intel/procfs/swap/device/*/free_percent": {},
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
                        "file": "/tmp/published_swap.log"
                    }
                }
            ]
        }
    }
}
```
Create a task:
```
$ snaptel task create -t swap-file.json
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-swap/issues) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-swap/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

And **thank you!** Your contribution, through code and participation, is incredibly important to us.

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [Andrzej Kuriata](https://github.com/andrzej-k)
