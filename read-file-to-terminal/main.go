package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fp := os.Args[1]
	pwd, err := os.Getwd()

	if err != nil {
		fmt.Println("An error ocurred: ", err)
		os.Exit(1)
	}

	f, err := os.Open(fmt.Sprintf("%s/%s", pwd, fp))
	defer f.Close()

	if err != nil {
		fmt.Println("An error ocurred: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
