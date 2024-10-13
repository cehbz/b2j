package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/anacrolix/torrent/bencode"
)

func main() {
	var input io.Reader

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	var bencodedData interface{}
	err := bencode.NewDecoder(input).Decode(&bencodedData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding bencoded data: %v\n", err)
		os.Exit(1)
	}

	jsonData, err := json.MarshalIndent(bencodedData, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding to JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
