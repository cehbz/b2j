package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/anacrolix/torrent/bencode"
)

func main() {
	// Parse command-line flags.
	verbose := flag.Bool("v", false, "output large binary fields verbatim")
	flag.Parse()

	// Determine input source: file or standard input.
	var input io.Reader
	if len(flag.Args()) > 0 {
		file, err := os.Open(flag.Args()[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		input = os.Stdin
	}

	decoder := bencode.NewDecoder(input)
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")

	var token interface{}
	for {
		// Decode bencoded data.
		err := decoder.Decode(&token)
		if err != nil {
			if err == io.EOF || err.Error() == "unexpected EOF" {
				break
			}
			fmt.Fprintf(os.Stderr, "Error decoding bencoded data: %v\n", err)
			os.Exit(1)
		}

		// Sanitize data if verbose mode is not enabled.
		if !*verbose {
			token = sanitize(token)
		}

		// Encode to JSON and output.
		err = encoder.Encode(token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding to JSON: %v\n", err)
			os.Exit(1)
		}
	}
}

// sanitize truncates large strings to "..." unless verbose mode is enabled.
func sanitize(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			v[key] = sanitize(value)
		}
	case []interface{}:
		for i, value := range v {
			v[i] = sanitize(value)
		}
	case string:
		if len(v) > 100 {
			return "..."
		}
	}
	return data
}
