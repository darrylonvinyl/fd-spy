# fd-spy

**fd-spy** is a lightweight Linux CLI utility written in Go that visualizes the open file descriptors of a specific process.

It parses the `/proc` filesystem to resolve file descriptors to their actual targets (files, sockets, pipes, or devices), giving you a quick "under the hood" look at what a process is holding onto.

## Usage

```bash
# Build (or run directly)
go run main.go <PID>

Example Output:

/dev/pts/0   (STDIN)
/dev/pts/0   (STDOUT)
/dev/pts/0   (STDERR)
SOCKET       socket:[12345]
PIPE         pipe:[67890]
FILE         /usr/share/code/resources.pak
