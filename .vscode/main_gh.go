// Copyright 2021 Rafal Glinski. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package implements functions for bulk IP address DNS lookup.
// It uses goroutines to speed up the process.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	errVal   = flag.String("e", "err", "Hostnames value on error")
	format   = flag.String("f", "%v\n\t%v\n", "Format output (ip, hostnames)")
	fileName = flag.String("i", "./addresses.txt", "Input file name")
	sep      = flag.String("s", "\n\t", "Hostnames separator")
	trimDot  = flag.Bool("t", true, "Trim hostnames \".\" suffix")
)

func loadLines(lines *[]string) {
	file, err := os.Open(*fileName)
	if err != nil {
		panic(fmt.Sprintf("Error opening file \"%v\"", *fileName))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
}

func lookup(ip string, c chan string) {
	names, err := net.LookupAddr(ip)
	if err != nil {
		c <- fmt.Sprintf(*format, ip, *errVal)
		return
	}
	if *trimDot {
		for i := range names {
			names[i] = strings.TrimSuffix(names[i], ".")
		}
	}
	c <- fmt.Sprintf(*format, ip, strings.Join(names, *sep))
}

func main() {
	flag.Parse()
	var ipAddrs []string
	loadLines(&ipAddrs)
	c := make(chan string, len(ipAddrs))
	for _, ip := range ipAddrs {
		go lookup(ip, c)
	}
	for range ipAddrs {
		fmt.Println(<-c)
	}
}
