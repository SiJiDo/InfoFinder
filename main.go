package main

import (
	"InfoFinder/cmd"
	"flag"
)

const TARGETURL int = 0
const TARGETFILE int = 0

func main() {
	var url string
	var filename string
	var wordlist string
	var threads int
	var timeout int
	var output string
	var jsurlflag bool
	var jsinfoflag bool
	var dirbflag bool

	flag.StringVar(&url, "u", "", "input target host(s) to scan")
	flag.StringVar(&filename, "l", "", "input file containing list of hosts to process")
	flag.StringVar(&wordlist, "w", "db/dic.txt", "dir brute wordlist(deaflut db/dic.txt)")
	flag.IntVar(&threads, "t", 50, "number of threads to use (default 50)")
	flag.IntVar(&timeout, "timeout", 5, "timeout in seconds (default 5)")
	flag.StringVar(&output, "o", "", "file to write output results(default target name)")
	flag.BoolVar(&jsurlflag, "jsurl", true, "find url in js(default true)")
	flag.BoolVar(&jsinfoflag, "jsinfo", true, "find info by config in js(deafult true)")
	flag.BoolVar(&dirbflag, "dirb", true, "scan by dir brute(deaflut true)")

	flag.Parse()

	if url != "" {
		cmd.Execute(TARGETURL, url, threads, timeout, output, jsurlflag, jsinfoflag, dirbflag, wordlist)
	} else if filename != "" {
		cmd.Execute(TARGETFILE, filename, threads, timeout, output, jsurlflag, jsinfoflag, dirbflag, wordlist)
	} else {
		flag.Usage()
	}
}
