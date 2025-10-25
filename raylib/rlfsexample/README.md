# Raylib Fullscreen Example

- Example for a fullscreen app with raylib.
- Hides cursor for execution.
- Waits for 'q' to close
- targets 60 fps


```shell
# init the module
go mod init github.com/mopore/go_playground/raylib/rlfsexample

# search for dependencies in existing source code
go mod tidy

# pre download dependencies
go get -u ./...

# Create the executable
go build -o bin/rlfsexample_darwin_arm64  cmd/rlfsexample/main.go
```
