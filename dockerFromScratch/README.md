This will create a very small Docker image by using the "from scratch" base image.

Run the following command to build the Docker image:

```bash
docker buildx -t from-scratch-image .
```

After the build is complete, you can run the Docker image with the following command:

```bash
docker run --rm from-scratch-image
```
