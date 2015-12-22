# snap-plugin-collector-swap

snap plugin for collecting swap metrics from /proc linux filesystem

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
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
#### Download df plugin binary:
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

## Documentation

### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | 
----------|-----------|
/intel/linux/swap/io/in_bytes_per_sec | float64 | 
/intel/linux/swap/io/in_pages_per_sec | float64 | 
/intel/linux/swap/io/out_bytes_per_sec | float64 | 
/intel/linux/swap/io/out_pages_per_sec | float64 | 
/intel/linux/swap/device/{device}/used_bytes | float64 | 
/intel/linux/swap/device/{device}/used_percent | float64 | 
/intel/linux/swap/device/{device}/free_bytes | float64 | 
/intel/linux/swap/device/{device}/free_percent | float64 | 
/intel/linux/swap/all/used_bytes | float64 | 
/intel/linux/swap/all/used_percent | float64 | 
/intel/linux/swap/all/free_bytes | float64 | 
/intel/linux/swap/all/free_percent | float64 | 
/intel/linux/swap/all/cached_bytes | float64 | 
/intel/linux/swap/all/cached_percent | float64 | 

### Examples
Example task manifest to use swap plugin:
```
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "5s"
    },
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/linux/swap/device/dev_sda5/used_bytes": {},
                "/intel/linux/swap/device/dev_sda5/used_percent": {},
                "/intel/linux/swap/device/dev_sda5/free_bytes": {},
                "/intel/linux/swap/device/dev_sda5/free_percent": {},
                "/intel/linux/swap/all/cached_bytes": {},
                "/intel/linux/swap/all/cached_percent":{},
                "/intel/linux/swap/all/free_bytes":{},
                "/intel/linux/swap/all/free_percent":{},
                "/intel/linux/swap/all/used_bytes":{},
                "/intel/linux/swap/all/used_percent":{},
                "/intel/linux/swap/io/in_bytes_per_sec":{},
                "/intel/linux/swap/io/in_pages_per_sec":{},
                "/intel/linux/swap/io/out_bytes_per_sec":{},
                "/intel/linux/swap/io/out_pages_per_sec":{}
            },
            "config": {
            },
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


### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [Andrzej Kuriata](https://github.com/andrzej-k)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
