package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"flag"
)

func main() {
	var open = flag.String("open", ".", "please specify -name file")
	flag.Parse()
	files, err:= ioutil.ReadDir(*open)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)

	for i, file := range files {
	    fmt.Print(i)
		fmt.Println(file.Name())
	}
}
