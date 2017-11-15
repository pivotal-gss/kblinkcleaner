package main


import (
	"flag"
	"fmt"
	"os"
	"encoding/json"
)

var (
	kbData = flag.String("f", "", "Json file containing the kb articles")
)

func parseJSON(f string) (*RootObj, error) {
	robj := new(RootObj)
	fh, err := os.Open(f)
	if err != nil {
		return robj, err
	}

	decoder := json.NewDecoder(fh)
	err = decoder.Decode(&robj)
	if err != nil {
		return robj, err
	}
	return robj, nil
}

func main() {
	flag.Parse()

	if *kbData == "" {
		fmt.Println("You must specify -f <file.json>")
		os.Exit(1)
	}

	a, err := parseJSON(*kbData)
	if err != nil {
		fmt.Println(err)
	}

	a.Clean()
	enc := json.NewEncoder(os.Stdout)
	enc.SetEscapeHTML(false)
	enc.Encode(a)
}