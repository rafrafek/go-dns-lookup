# Bulk IP address DNS lookup

Lookup hostnames for IP addresses loaded from file.

Program uses goroutines to speed up the process.

Written in Go (Golang).

## Usage

Requires Go installed.

Build and run:

```
go build
./main
```
Example input `addresses.txt`:
```
172.217.20.206
216.58.215.110
127.58.215.110
99.83.207.202
212.77.98.9
```
Example output:
```
172.217.20.206
        waw02s08-in-f206.1e100.net
        waw02s08-in-f14.1e100.net

216.58.215.110
        waw02s17-in-f14.1e100.net

212.77.98.9
        www.wp.pl

99.83.207.202
        aafc88a28d9997374.awsglobalaccelerator.com

127.58.215.110
        err
```
Command line arguments `./main -help`:
```
Usage:
  -e string
        Hostnames value on error (default "err")
  -f string
        Format output (ip, hostnames) (default "%v\n\t%v\n")
  -i string
        Input file name (default "./addresses.txt")
  -s string
        Hostnames separator (default "\n\t")
  -t bool
        Trim hostnames "." suffix (default true)
```
Note that `-t` argument requires `-t=` notation e.g. `-t=false` as it is bool type.
