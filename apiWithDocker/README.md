```
    _    ____ ___             __  ____             _             
   / \  |  _ \_ _| __      __/ / |  _ \  ___   ___| | _____ _ __ 
  / _ \ | |_) | |  \ \ /\ / / /  | | | |/ _ \ / __| |/ / _ \ '__|
 / ___ \|  __/| |   \ V  V / /   | |_| | (_) | (__|   <  __/ |   
/_/   \_\_|  |___|   \_/\_/_/    |____/ \___/ \___|_|\_\___|_|   
                                                                 

```

# Howto

Start with creating the module in the current folder:

```bash
go mod init apiWithDocker
```

Install the dependency to `github.com/gin-gonic/gin`:
`-u` means update.

```bash
go get -u github.com/gin-gonic/gin
```

In case you added the depencies (esp. when removing one) run `go mod tidy`.

Create your main go file see `main.go`.

Provide a `Dockerfile` (see `Dockerfile`).

Create the Docker image:

```bash
docker buildx -t api-with-docker .
```

Run the Docker image:

```bash
docker run -p 8080:8080 api-with-docker
```

