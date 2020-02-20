# CLI-fops
Implement a command-line application fops written in Golang.

## Quick Start
![](https://raw.githubusercontent.com/amosricky/CLI-fops/master/src/demo_start.gif)
```
$ go run main.go 
or
$ ./CLI-fops 
```
## Test subcommand
![](https://raw.githubusercontent.com/amosricky/CLI-fops/master/src/demo_test.png)
```
$ go test
```
## Features
1. Use spf13/cobra as command framework for this project.
2. Use go-ini/ini to read configuration more structured.
3. Implement "fops" command that could count line or generate checksum with the specific file.

## How it works
```
.
├── cli               // Use Cobra to define the CLI command.
│   └── cli.go
├── CLI-fops          // An executive file.
├── conf              // System configuration
│   └── app.ini
├── main.go           // Create an loop to read stdin stream.
├── myfile.txt        // Test file.
├── setting           // Initialize the configuration
│   └── setting.go
...

```

## CLI Document
```
$ fops help
File Ops

Usage:
   fops [command]

Available Commands:
  checksum    Get checksum
  linecount   Count line for file.
  version     Get system version.

Flags:
  -h, --help   help for fops
```
```
$ fops linecount --help
Count line for file.

Usage:
   fops linecount [flags]

Flags:
  -f, --file string   File path
  -h, --help          help for linecount

```
```
$ fops checksum --help
Get checksum

Usage:
   fops checksum [flags]

Flags:
  -f, --file string   File path
  -h, --help          help for checksum
      --md5           Get checksum in hash function-sha256
      --sha1          Get checksum in hash function-sha1
      --sha256        Get checksum in hash function-sha256
```
```
$ fops version --help
Get system version.

Usage:
   fops version [flags]

Flags:
  -h, --help   help for version
```

## Demo
Command [lincount]
* Count line for myfile.txt.
* Count line for a not exist file.
* Count line for a binary file.
* Count line for a directory.

![](https://raw.githubusercontent.com/amosricky/CLI-fops/master/src/demo_linecount.gif)

Command [checksum]
* Generate checksum with hash function - MD5.
* Generate checksum with hash function - SHA1.
* Generate checksum with hash function - SHA256.

![](https://raw.githubusercontent.com/amosricky/CLI-fops/master/src/demo_checksum.gif)

Command [version]
* Get version.

![](https://raw.githubusercontent.com/amosricky/CLI-fops/master/src/demo_version.gif)