# Raylib Fullscreen Example

- Example for a fullscreen app with raylib.
- Hides cursor for execution.
- Waits for 'q' to close
- targets 60 fps

For Arch Linux VM with Wayland on M2 Mac with QEMU
```shell
export CGO_CFLAGS="-DGRAPHICS_API_OPENGL_21"
go clean -cache
go run ./cmd/basictestvm/main.go
```


## General Project Setup
```shell
# init the module
go mod init github.com/mopore/go_playground/raylib/basictestvm

# search for dependencies in existing source code
go mod tidy

# pre download dependencies
go get -u ./...

```
