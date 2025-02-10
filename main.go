package main

import (
	"os"
	"fmt"
	"github.com/schollz/progressbar/v3"
)

func main() {
	d := NewData()

	files, _ := os.ReadDir("demos")

	bar := progressbar.Default(int64(len(files)))
	for _, f := range files {
		d.Read("demos/" + f.Name())
		bar.Add(1)
	}

	fmt.Println(d.GetDeaths())
	fmt.Println(d.GetOwners())
}