```
    _    ____ ___             __  ____             _             
   / \  |  _ \_ _| __      __/ / |  _ \  ___   ___| | _____ _ __ 
  / _ \ | |_) | |  \ \ /\ / / /  | | | |/ _ \ / __| |/ / _ \ '__|
 / ___ \|  __/| |   \ V  V / /   | |_| | (_) | (__|   <  __/ |   
/_/   \_\_|  |___|   \_/\_/_/    |____/ \___/ \___|_|\_\___|_|   
                                                                 

```

This is the successor example project using Go's 1.22 advancements in the
net/http package.
Also the Docker image creation is optimized with a from scratch image.

# Howto
Start with creating the module in the current folder:

```bash
go mod init tut/apiwithdocker
```

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


# Testing with Curl
Use the following curl command to test the post of a user:
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"firstname":"Oliver","id":1,"lastname":"Queen"}' \
http://localhost:8080/api/v2/user
```
