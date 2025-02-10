package main

import (
	"os"
	"strings"
	"fmt"
	"github.com/schollz/progressbar/v3"
)

func main() {
	files, _ := os.ReadDir("demos")


	// only read .dem files
	var demFiles = make([]string, 0)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".dem") {
			demFiles = append(demFiles, f.Name())
		}
	}

	c := make(chan *Data, len(demFiles))
	for _, f := range demFiles {
		go Read("demos/" + f, c)
	}

	bar := progressbar.Default(int64(len(demFiles)))

	ad := NewAggregatedData()
	i := 0
	for d := range c {
		i += 1
		ad.Add(d)
		bar.Add(1)
		if i == len(demFiles) - 1 {
			close(c)
		}
	}

	fmt.Printf("%20s | %s | %s | %s \n", "        Knife        ", "Games", "Deaths", "Ratio")
	for _, kt := range AllKnifeTypes {
		fmt.Printf("%-20s | %5d | %5d | %.3f \n", ToString(kt), ad.Games[kt], ad.Deaths[kt], float64(ad.Deaths[kt])/float64(ad.Games[kt]))
	}
}