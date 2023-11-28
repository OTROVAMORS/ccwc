// cli.go
package cli

import (
	"cc/bison/ccwc/lib"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)


func Execute() {

	countBytes := flag.Bool("c", false, "Count bytes")
	countLines := flag.Bool("l", false, "Count lines")
	countWords := flag.Bool("w", false, "Count words")
	countCharacters := flag.Bool("m", false, "Count characters")
	flag.Parse()

