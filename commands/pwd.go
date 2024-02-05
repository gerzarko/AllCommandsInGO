package commands

import (
	"fmt"
	"os"
)

func Pwd() {
	f, err := os.Open("./")
	if err != nil {
		fmt.Printf("can't enter file, error %v", err)
		return
	}

	fi, error := f.ReadDir(0)

	if error != nil {
		fmt.Printf("can't enter file, error %v", error)
		return
	}

	for i := range fi {
		fmt.Println(fi[i].Name())
	}
}
