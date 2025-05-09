COMMANDS USED
Command	Purpose
docker build -t jenkins-golang-bookapi .	Build Docker image
docker run -d -p 8090:8080 -p 8091:8081 --name jenkins-golang-bookapi jenkins-golang-bookapi	Start Jenkins container
docker exec -it jenkins-golang-bookapi /bin/bash	Access Jenkins container shell
cat /var/jenkins_home/secrets/initialAdminPassword	Get admin login password
go mod tidy	Install Go dependencies
go build -o bookapi_server .	Build the Go server binary
./bookapi_server	Run the Go server binary

Steps to Create Jenkins Freestyle Project
1. Dockerfile Setup
Created a Dockerfile to add Go to the Jenkins image.

dockerfile
Copy
Edit
FROM jenkins/jenkins:lts
USER root
RUN apt-get update && apt-get install -y golang-go git && apt-get clean
ENV GOPATH=/go
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH
USER jenkins
2. Build Docker Image
Built the Docker image:

bash
Copy
Edit
docker build -t jenkins-golang-bookapi .
3. Run Jenkins Container
Ran Jenkins container and exposed the necessary ports:

bash
Copy
Edit
docker run -d -p 8090:8080 -p 8091:8081 --name jenkins-golang-bookapi jenkins-golang-bookapi
4. Get Jenkins Admin Password
Retrieved the admin password:

bash
Copy
Edit
docker exec -it jenkins-golang-bookapi cat /var/jenkins_home/secrets/initialAdminPassword
5. Configure Jenkins Job
Open Jenkins UI at http://localhost:8090

Click New Item

Name it GoBookAPI

Choose Freestyle Project

In Source Code Management, add the repository URL.

Under Build Steps, use this script:

bash
Copy
Edit
cd /var/jenkins_home/workspace/GoBookAPI
go mod tidy
go build -o bookapi_server .
6. Run Go Server
After build, run the Go server:

bash
Copy
Edit
docker exec -it jenkins-golang-bookapi /bin/bash
cd /var/jenkins_home/workspace/GoBookAPI
./bookapi_server
