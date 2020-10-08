# cidrgrep

[![Go Report Card](https://goreportcard.com/badge/github.com/tomdoherty/cidrgrep)](https://goreportcard.com/report/github.com/tomdoherty/cidrgrep)
[![Go Actions Status](https://github.com/tomdoherty/cidrgrep/workflows/Go/badge.svg)](https://github.com/tomdoherty/cidrgrep/actions)

```shell
cidrgrep "10.0.0.0/24" < file.log
cidrgrep "10.0.0.0/24" file1.log file2.log file3.log
```

Outputs lines which contain an IP within the CIDR block specified
