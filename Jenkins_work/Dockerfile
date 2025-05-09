# Use Jenkins base image
FROM jenkins/jenkins:lts

USER root

# Install Go
RUN apt-get update && \
    apt-get install -y golang-go && \
    apt-get clean

# Set Go environment
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

USER jenkins
