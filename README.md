# kernel-message-reporter

Run a command and report kernel messages (`/dev/kmsg`) logged while it was running.

## Usage

```sh
kernel-message-reporter [-o output_file] <CMD> [ARGS...]
```

- `-o`: Write kernel messages to a file instead of stderr.

## Example

```sh
kernel-message-reporter -o kmsg.log dd if=/dev/zero of=/dev/null bs=1M count=1000
```

## Requirements

- Linux with access to `/dev/kmsg`

## Install

Download a prebuilt binary from the [Releases](https://github.com/appare45/kernel-message-reporter/releases) page, or build from source:

```sh
go build -o kernel-message-reporter .
```

## License

MIT
