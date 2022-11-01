# tmbts-image-resizer

Cloud Function for resizing images in Cloud Storage

## Development

- Go 1.11+
- Dependencies managed with `go mod`

### Setup

These steps will describe how to setup this project for active development. Adjust paths to your desire.

1. Clone the repository: `git clone git@github.com:kyleshepherd/tmbts-image-resizer.git tmbts-image-resizer`
2. Build: `make build`
3. 🍻

### Dependencies

Dependencies are managed using `go mod` (introduced in 1.11), their versions
are tracked in `go.mod`.

To add a dependency:

```
go get url/to/origin
```

### Configuration

Configuration can be provided through a toml file, these are loaded
in order from:

- `/etc/tmbts-image-resizer/tmbts-image-resizer.toml`
- `$HOME/.config/tmbts-image-resizer.toml`

Alternatively a config file path can be provided through the
-c/--config CLI flag.

#### Example tmbts-image-resizer.toml

```toml
[log]
console = true
level = "debug"  # [debug|info|error]
```
