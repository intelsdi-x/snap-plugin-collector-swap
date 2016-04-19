# snap plugin collector - swap

## Collected Metrics

This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description
----------|-----------|-----------------------
/intel/procfs/swap/io/in_bytes_per_sec | float64 | amount of the memory the system paged in (B per sec)
/intel/procfs/swap/io/in_pages_per_sec | float64 | number of pages the system paged in (pages per sec)
/intel/procfs/swap/io/out_bytes_per_sec | float64 | amount of the memory the system paged out (B per sec)
/intel/procfs/swap/io/out_pages_per_sec | float64 | number of pages the system paged out (pages per sec)
/intel/procfs/swap/device/{device}/used_bytes | float64 | used swap space (MB)
/intel/procfs/swap/device/{device}/used_percent | float64 | used swap space (percentage)
/intel/procfs/swap/device/{device}/free_bytes | float64 | free swap space (MB)
/intel/procfs/swap/device/{device}/free_percent | float64 | free swap space (percentage)
/intel/procfs/swap/all/used_bytes | float64 | total amount of swap space available (MB)
/intel/procfs/swap/all/used_percent | float64 | total amount of swap space available (percentage)
/intel/procfs/swap/all/free_bytes | float64 | amount of swap space that is currently unused (MB)
/intel/procfs/swap/all/free_percent | float64 |  amount of swap space that is currently unused (percentage)
/intel/procfs/swap/all/cached_bytes | float64 | amount of memory that once was swapped out, is swapped back in but still also is in the swap file (MB)
/intel/procfs/swap/all/cached_percent | float64 | amount of memory that once was swapped out, is swapped back in but still also is in the swap file (percentage)