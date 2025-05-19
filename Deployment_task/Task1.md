# Phone Number Validator

---

### 1. Jenkins Pipeline

I created a Jenkins pipeline to build a Go project. The pipeline compiles the Go code and stores the output as a zip artifact named `phone_number_validate.zip`.

### 2. Unzip Artifact

After Jenkins completes the build, I unzipped the artifact and navigated to the folder:

```bash
unzip phone_number_validate.zip
cd phone_app
```

At this point, the folder contained:

```bash
Dockerfile
main
```

---

## Docker Setup

### 3. Dockerfile Content

I created a Dockerfile with the following content:

```dockerfile
FROM alpine:latest
WORKDIR /app
COPY main .
RUN chmod +x main
CMD ["./main"]
```

This Dockerfile uses the Alpine base image, sets up a working directory, copies the Go binary, makes it executable, and defines the default command.

---

### 4. Build Docker Image

I built the Docker image using the following command:

```bash
docker build -t sivasankar123/phone-validator:latest .
```

---

### 5. Docker Login

I logged into Docker Hub:

```bash
docker login
```

I entered my Docker Hub username and password to authenticate.

---

### 6. Push Image to Docker Hub

After building the image, I pushed it to Docker Hub:

```bash
docker push sivasankar123/phone-validator:latest
```

---

## Usage

### Pull and Run on Any Machine

On any machine that has Docker installed, I can pull the image with:

```bash
docker pull sivasankar123/phone-validator:latest
```

### Run with Phone Number

To validate a phone number, I ran the following command:

```bash
docker run -it sivasankar123/phone-validator ./main <phone-number>
```

Example:

```bash
docker run -it sivasankar123/phone-validator ./main +919876543210
```

If the output is "Invalid phone number", I check if the input format matches what the Go program expects.

---

## Debug Inside the Container

To access the container shell and run the program manually, I used:

```bash
docker run -it sivasankar123/phone-validator /bin/sh
```

Then inside the shell, I executed:

```bash
./main <phone-number>
```
---
