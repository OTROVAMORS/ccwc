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


	if !(*countBytes || *countLines || *countWords || *countCharacters) {
		*countBytes = true
		*countLines = true
		*countWords = true
	}


	filePath := flag.Arg(0)
	if filePath == "" {

		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from standard input:", err)
			return
		}
		lib.ProcessData(data, *countBytes, *countLines, *countWords, *countCharacters)
		return
	}


	lib.CountAndPrint(filePath, *countBytes, *countLines, *countWords, *countCharacters)
}


