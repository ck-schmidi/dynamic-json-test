package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/valyala/fastjson"
)

func main() {
	var fileName = flag.String("file", "", "filename for file to test")
	flag.Parse()

	// open file for test
	file, err := os.Open(*fileName)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	var p fastjson.Parser
	v, err := p.ParseBytes(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all: %v\n", v)

	// since we know entries is an array, we can
	// simply get it as a slice
	entries := v.GetArray("entries")
	fmt.Printf("entries: %v\n", entries)

	// so we can easily iterate the slice
	for _, entry := range entries {
		fmt.Printf("- %v\n", entry)
		// and get some values by a key
		for _, key := range []string{"int-value", "name", "xxx"} {
			value := entry.Get(key)
			fmt.Printf("  - %v\n", value)
		}
	}
}
