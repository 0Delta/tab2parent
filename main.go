package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"fmt"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	fileName := os.Args[1]
	lines, err := ReadFile(fileName)
	if err != nil {
		log.Print("[FATAL] File open failed.")
		return
	}

	parent := []string{}
	ret := ""
	sep := "\t"
	for _, l := range lines {
		retl := []string{}
		for i, w := range(strings.Split(l, sep)){
			if w == "" {
				retl = append(retl, parent[i])
			} else {
				retl = append(retl, w)
			}
		}
		ret += strings.Join(retl, ".") + "\n"
		parent = retl
	}
	fmt.Println(ret)
}

func ReadFile(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("[ERROR] File %s could not read: %v\n", filePath, err)
		return nil, err
	}
	defer f.Close()
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("[ERROR] File %s scan error: %v\n", filePath, err)
		return nil, err
	}
	return lines, nil
}
