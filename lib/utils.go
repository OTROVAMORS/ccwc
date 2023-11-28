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

func ProcessData(data []byte, countBytes, countLines, countWords, countCharacters bool) {
    if countBytes {
        byteCount := len(data)
        fmt.Printf("Bytes: %8d\n", byteCount)
    }

    if countLines {
        lineCount := countLinesInData(data)
        fmt.Printf("Lines: %8d\n", lineCount)
    }

    if countWords {
        wordCount := countWordsInData(data)
        fmt.Printf("Words: %8d\n", wordCount)
    }

    if countCharacters {
        charCount := countCharactersInData(data)
        fmt.Printf("Characters: %8d\n", charCount)
    }
}

func ReadFile(filePath string) ([]byte, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func countLinesInData(data []byte) int {
    lineCount := 0
    for _, b := range data {
        if b == '\n' {
            lineCount++
        }
    }
    return lineCount
}

func countWordsInData(data []byte) int {
    wordCount := 0
    inWord := false
    for _, b := range data {
        if unicode.IsSpace(rune(b)) {
            inWord = false
        } else if !inWord {
            wordCount++
            inWord = true
        }
    }
    return wordCount
}

