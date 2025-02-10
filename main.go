package main

import (
	"os"
	"strings"
	"fmt"
	"github.com/schollz/progressbar/v3"
)

func main() {
	files, _ := os.ReadDir("demos")

	bar := progressbar.Default(int64(len(files)))

	c := make(chan *Data, len(files))
	for _, f := range files {
		// only read .dem files
		if strings.HasSuffix(f.Name(), ".dem") {
			go Read("demos/" + f.Name(), c)
		}
	}

	ad := NewAggregatedData()
	for d := range c {
		ad.Add(d)
		bar.Add(1)
	}

	for _, kt := range AllKnifeTypes {
		fmt.Printf("%-20s | %-2d | %-2d \n", ToString(kt), ad.Owners[kt], ad.Deaths[kt])
	}
}