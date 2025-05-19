### Docker Image Deployment in Docker Hub via jenkins pipeline
---

### Docker-Based Jenkins Pipeline for Go Binary Deployment

In this task, I implemented a CI/CD workflow using Jenkins to build and push a Go-based binary Docker image to Docker Hub.

#### Step 1: Created Dockerfile for Application

I used Alpine Linux as the base image to keep the image lightweight. The pre-built Go binary was copied to the container, given executable permissions, and executed as the entry point.

**Dockerfile (Deployment\_task/Dockerfile):**

```dockerfile
FROM alpine:latest

WORKDIR /app

COPY Deployment_task/main /app/main 

RUN chmod +x main

ENTRYPOINT ["./main"]
```

#### Step 2: Built a Custom Jenkins Image with Go and Docker CLI

Since the default Jenkins image does not include Docker CLI or Go, I created a custom Jenkins image with both tools pre-installed. I installed Go (v1.21.4) and Docker CLI, set up the Go environment, and granted the Jenkins user access to Docker.

**Custom Jenkins Dockerfile:**

```dockerfile
FROM jenkins/jenkins:lts

USER root

RUN apt-get update && \
    apt-get install -y curl gnupg lsb-release zip tar wget && \
    apt-get clean

ENV GO_VERSION=1.21.4
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/opt/go
ENV PATH=$PATH:$GOPATH/bin

RUN mkdir -p /opt/go && chown -R jenkins:jenkins /opt/go

RUN curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" \
    > /etc/apt/sources.list.d/docker.list && \
    apt-get update && \
    apt-get install -y docker-ce-cli

RUN groupadd docker && usermod -aG docker jenkins

USER jenkins
```

#### Step 3: Configured Jenkinsfile for Pipeline

I created a Jenkinsfile to automate the build and push process. The pipeline has the following stages:

* Check Docker installation and container access
* Build the Docker image from the binary
* Log in to Docker Hub
* Push the image to the Docker Hub repository

**Jenkinsfile:**

```groovy
pipeline {
    agent any

    environment {
        IMAGE_NAME = "sivasankar123/phone-number-validator"
        TAG = "latest"
    }

    stages {
        stage('Check Docker Access') {
            steps {
                sh 'docker --version'
                sh 'docker ps'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh "docker build -t ${IMAGE_NAME}:${TAG} -f Deployment_task/Dockerfile ."
            }
        }

        stage('Login to Docker Hub') {
            steps {
                sh 'docker login -u sivasankar123 -p Siva@2003'
            }
        }

        stage('Push Docker Image') {
            steps {
                sh "docker push ${IMAGE_NAME}:${TAG}"
            }
        }
    }
}
```

#### Step 4: Verified the Build

After the pipeline completed, I verified that the image was successfully pushed to Docker Hub and can be pulled and run using:

```bash
docker pull sivasankar123/phone-number-validator
docker run sivasankar123/phone-number-validator
```

---

