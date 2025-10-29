# Raylib Fullscreen Example

- Example for a fullscreen app with raylib.
- Hides cursor for execution.
- Waits for 'q' to close
- targets 60 fps

For Arch Linux VM with Wayland on M2 Mac with QEMU at the moment of writing this only OpenGL 2.1
is support. While 3.3 is standard you have to manually set the environment with:
```shell
export CGO_CFLAGS="-DGRAPHICS_API_OPENGL_21"
```

As an alternative add an entry to `/etc/environment`


## How the projet was setup up
```shell
# init the module
go mod init github.com/mopore/go_playground/raylib/basictestvm

# search for dependencies in existing source code
go mod tidy

# pre download dependencies
go get -u ./...

```
