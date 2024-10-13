# Bencode to JSON Converter

This tool is a Go program that converts bencoded files to JSON format. It reads bencoded data from a specified file or standard input and outputs the parsed JSON to standard output.

## Features

- **Input Source**: Accepts a bencoded file as a command-line argument or reads from standard input if no file is provided.
- **JSON Output**: Outputs the parsed data in JSON format, properly indented for readability.
- **Verbose Mode**: Includes a `-v` flag to control whether large binary fields are output verbatim or truncated.

## Usage

### Prerequisites

- Go 1.16 or later
- The [anacrolix/torrent](https://github.com/anacrolix/torrent) package, which provides the bencode decoding functionality

To install the dependency, use:

```sh
go get github.com/anacrolix/torrent/bencode
```

### Running the Program

```sh
go run main.go [-v] [input-file]
```

- **`-v`**: Verbose mode flag. If set, large binary fields are output verbatim. By default, fields larger than 100 characters are replaced with `"..."`.
- **`input-file`**: Path to the bencoded file to be converted. If not provided, the program reads from standard input.

### Examples

#### Convert a Bencoded File to JSON

```sh
go run main.go example.bencode
```

This command reads from `example.bencode` and outputs the converted JSON to standard output.

#### Read from Standard Input

```sh
cat example.bencode | go run main.go
```

This command reads from `example.bencode` via a pipe and converts the data to JSON.

#### Verbose Mode

```sh
go run main.go -v example.bencode
```

This command reads from `example.bencode` and outputs the converted JSON with all fields included, without truncating large binary fields.

## Code Overview

- **`main()`**: Handles command-line arguments and input source selection, and uses `bencode` to decode the input.
- **`sanitize(data interface{}) interface{}`**: Sanitizes large binary fields by truncating them if they exceed 100 characters, unless verbose mode is enabled.
- **Verbose Mode**: When enabled (`-v` flag), the program outputs all fields without sanitizing large strings.

## Notes

- The program relies on the `anacrolix/torrent/bencode` library for decoding bencoded data.
- The `sanitize` function helps to reduce output size by truncating large binary fields when verbose mode is disabled.

## License

This project is licensed under the MIT License.
