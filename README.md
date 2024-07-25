# `wordc` Replica in Go

This is a simple Go implementation of the Unix `wc` (word count) command. It counts bytes, lines, words, and characters in a file or from standard input, based on the provided flags.

## Features

- **Byte Count** (`-c`): Count the number of bytes in the file.
- **Line Count** (`-l`): Count the number of lines in the file.
- **Word Count** (`-w`): Count the number of words in the file.
- **Character Count** (`-m`): Count the number of characters in the file.

## Usage

The program can be run with or without a file argument. If no file is specified, it reads from standard input.

### Command-Line Syntax

```bash
wordc [flags] [file]
```

### Flags

- `-c`: Count bytes in the file.
- `-l`: Count lines in the file.
- `-w` : Count words in the file.
- `-m`: Count characters in the file.

## Examples

Count bytes, lines, words, and characters in example.txt

```bash
wordc -c -l -w -m example.txt
```

Count bytes and words in example.txt

```bash
wordc -c -w example.txt
```

Read from standard input and count lines and words

```bash
echo -e "Hello\nWorld" | wordc -l -w
```

## Behavior

- If no file argument is provided, the program reads from standard input.
- If no flags are specified, the program defaults to counting bytes, lines, words, and characters.
- The output format is a space-separated list of counts, followed by the file name (or - for standard input).

## Implementation Details

- Bytes: Counted using len(content).
- Lines: Counted by counting newline characters.
- Words: Counted by splitting the content into fields using strings.Fields.
- Characters: Counted using utf8.RuneCountInString.

The program ensures compatibility with standard input by checking if the input is coming from a terminal or not. 
If it is from a terminal, the program exits with an error message indicating that no input was provided.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
