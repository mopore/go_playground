```
 _____         _      ____            _        _                     
|_   _|__  ___| |_   / ___|___  _ __ | |_ __ _(_)_ __   ___ _ __ ___ 
  | |/ _ \/ __| __| | |   / _ \| '_ \| __/ _` | | '_ \ / _ \ '__/ __|
  | |  __/\__ \ |_  | |__| (_) | | | | || (_| | | | | |  __/ |  \__ \
  |_|\___||___/\__|  \____\___/|_| |_|\__\__,_|_|_| |_|\___|_|  |___/
                                                                     
```

This uses the test containers library (belongs to Docker) to utilize docker
containers for testing to easily spin up and tear down containers for testing.
In this example (see `main_test.go`), we spin up an Nginx container and test
if the container is available.
The container will be started on a random port we need get provided by the
testcontainers library.
The created Nginx container is accompanied by a cleanup container ("reaper").
Both containers will be stopped and removed after the test automatically.


# Pre-requisites
- Go
- Docker
- Existing Go project

# Setup
Use the testify assert package
```shell
go get github.com/stretchr/testify/assert
```

Get the test containers library
```shell
go get github.com/testcontainers/testcontainers-go
```

# Run
```shell
go test -v ./...
```
```
