# SRT Offset

A simple command-line tool to offset SRT subtitle files.
This tool allows you to adjust the timing of subtitles in SRT files, which can be useful when subtitles are out of sync with the video.

## Features

- Offset subtitles by a specified time (positive or negative).
- Supports input and output file paths.
- Can be used in scripts or manually from the command line.
- Cross-platform compatibility (Windows, Linux, macOS).

## Installation

You can install the tool by downloading the pre-built binaries from the [releases page](https://github.com/zambrinf/srt-offset/releases) or by cloning the repository and building it from source.

```bash
git clone https://github.com/zambrinf/srt-offset.git
cd srt-offset
go build -o srt-offset main.go
sudo mv srt-offset /usr/local/bin/
sudo chmod +x /usr/local/bin/srt-offset
```

## Usage

To use the tool, run one of the following commands:

```bash
srt-offset -i input.srt -offset <time>
srt-offset -i input.srt -offset <time> -o output.srt
```

Where:

- `-i` specifies the input SRT file.
- `-offset` specifies the time to offset the subtitles (e.g., `-0.5` to make the subtitles appear 0.5 seconds earlier).
- `-o` specifies the output SRT file. If not provided, it will overwrite the input file.

## Example

To offset subtitles by 0.5 seconds, you can run:

```bash
srt-offset -i input.srt -offset -0.5 -o output.srt
```

This will adjust the timing of all subtitles in `input.srt` by -0.5 seconds and save the result to `output.srt`.

## Testing

To run tests, you can use the following command:

```bash
go test -v ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find a bug or have a feature request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details
