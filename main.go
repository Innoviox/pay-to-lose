package main

import (
	"os"
)

func main() {
	files, _ := os.ReadDir("demos")
	for _, f := range files {
		read("demos/" + f.Name())
	}
}