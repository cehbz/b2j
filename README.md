# Bencode to JSON Converter

b2j is a go program to bencoded files to JSON. It reads bencoded data from a named file or stdin and outputs JSON to stdout.

## Features

- **Input Source**: Accepts a bencoded file name argument or reads from standard input if no name is provided.
- **JSON Output**: Outputs pretty printed JSON.
- **Verbose Mode**: By default large (>100 char) fields are elided to "...". The `-v` flag outputs them verbatim.

## Usage

### Prerequisites

- Go 1.16 or later
- [anacrolix/torrent](https://github.com/anacrolix/torrent), which provides the bencode decoding functionality

To install the dependency, use:

```sh
go get github.com/anacrolix/torrent/bencode
```

### Building the Program
```sh
go build
```
or
```sh
go install
```
like any other go program

### Running the Program

```sh
b2j [-v] [input-file]
```

- **`-v`**: Verbose mode flag. If set, large binary fields are output verbatim. By default, fields larger than 100 characters are replaced with `"..."`.
- **`input-file`**: Path to the bencoded file to be converted. If not provided, it uses standard input.

### Examples

#### Convert a Bencoded File to JSON

```sh
b2j example.bencode
```

Read from `example.bencode` and output the converted JSON to standard output.

#### Read from Standard Input

```sh
cat example.bencode | b2j
```

Read from `example.bencode` via stdin and convert the data to JSON.

#### Verbose Mode

```sh
b2j -v example.bencode
```

This command read from `example.bencode` and output the converted JSON without truncating large fields.

## License

This project is licensed under the MIT License.
