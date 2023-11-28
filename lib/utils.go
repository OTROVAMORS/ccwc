// lib.go
package lib

import (
	"fmt"
	"os"
	"unicode"
)

func CountAndPrint(filePath string, countBytes, countLines, countWords, countCharacters bool) {
    data, err := ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println("File Analysis:")
    ProcessData(data, countBytes, countLines, countWords, countCharacters)
}

