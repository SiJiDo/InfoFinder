package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

func OpenWordlist(wordlist string) []string {
	var results []string
	fi, err := os.Open(wordlist)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	tmp := strings.Replace(string(fd), "\r", "", -1)
	results = strings.Split(tmp, "\n")

	return results
}

func InitTarget(target string, wordlist string) []string {
	var targets []string
	if target[len(target)-1:] != "/" {
		target = target + "/"
	}
	paths := OpenWordlist(wordlist)
	for i := 0; i < len(paths); i++ {
		if paths[i][0] == '/' {
			paths[i] = paths[i][1:len(paths[i])]
		}
		targets = append(targets, target+paths[i])
	}

	return targets
}
