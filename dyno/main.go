package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/icza/dyno"
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

	// decode content of file into a dynamic go structure
	// which can be read with the dyno library
	var res map[string]interface{}
	err = json.NewDecoder(file).Decode(&res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("all: %v\n", res)

	// since we know entries is an array, we can
	// simply get it as a slice
	entries, err := dyno.GetSlice(res, "entries")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("entries: %v\n", entries)

	// so we can easily iterate the slice
	fmt.Printf("splitted up:\n")
	for _, entry := range entries {
		fmt.Printf("- %v\n", entry)
		// and get
		for _, key := range []string{"int-value", "name"} {
			value, err := dyno.Get(entry, key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("  - %v\n", value)
		}
	}
}
