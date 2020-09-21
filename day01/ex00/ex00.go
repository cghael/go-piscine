package main

import (
	"flag"
	"strings"
)

const (
	Usage       = "USAGE:\t-f databasename.xml\n\t\t-f databasename.json"
	ExitSuccess = 0
	ExitFailure = 1
)

func ftCheckFlags (str string) (err int, baseType string) {
	if str == "" {
		return ExitFailure, ""
	}
	baseName := strings.Split(str, ".")
	if len(baseName) != 2 || (baseName[1] != "xml" && baseName[1] != "json") {
		return ExitFailure, ""
	}
	return ExitSuccess, baseName[1]
}

func main() {
	var strFlags = flag.String("f", "", "-f databasename.type")
	flag.Parse()
	err, baseType := ftCheckFlags(*strFlags)
	if err == ExitFailure {
		println(Usage)
	} else {
		println(baseType)
	}
}
