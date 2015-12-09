package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-soundplayer"
)

func main() {
	for _, arg := range os.Args[1:] {
		if err := soundplayer.Play(arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
