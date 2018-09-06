// Copyright (c) 2018, Hayden Eskriett <hayden@eskriett.com>
// See LICENSE for licensing details

// yuri is a blazing fast tool for extracting URIs from plaintext
//  NAME
//
//    yuri [OPTION]... [FILE]...
//
//  SYNOPSIS
//
//     yuri - Extract URIs from FILE(s) or standard input
//
//  DESCRIPTION
//
//     yuri tries to extracting URIs of numerous schemes from text as fast as
//     possible. Compared to most similar tools which use regular expressions,
//     yuri uses a DFA built using ragel for performance.
//
//  OPTIONS
//
//     -h, -help
//         displays the help text and exits
//
//  EXAMPLES
//
//     yuri f - g
//         Outputs f's URIs, then standard input's URIs, then g's URIs
//
//  AUTHOR
//     Written by Hayden Eskriett <hayden@eskriett.com>
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/eskriett/yuri"
)

func init() {
	flag.Usage = func() {
		p := func(str string) {
			fmt.Fprintln(os.Stderr, str)
		}
		p("Usage: yuri [OPTION]... [FILE]...\n")
		p("Extracts URIs from FILE(s) or standard input if no files provided")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		args = []string{"-"}
	}

	for _, path := range args {
		if err := yankURIs(path); err != nil {
			log.Fatal(err)
		}
	}
}

func yankURIs(path string) error {
	r := os.Stdin

	if path != "-" {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()
		r = f
	}

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		for _, uri := range yuri.YankURIs([]byte(scanner.Text())) {
			fmt.Println(uri)
		}
	}

	return scanner.Err()
}
