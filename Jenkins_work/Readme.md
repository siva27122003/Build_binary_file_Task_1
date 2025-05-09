It looks like you're aiming to provide detailed steps for setting up a Jenkins Freestyle project without YAML or syntax errors. Here's the clean version of your instructions without the syntax issues:

---

### **COMMANDS USED**

| Command                                                                                        | Purpose                        |
| ---------------------------------------------------------------------------------------------- | ------------------------------ |
| `docker build -t jenkins-golang-bookapi .`                                                     | Build Docker image             |
| `docker run -d -p 8090:8080 -p 8091:8081 --name jenkins-golang-bookapi jenkins-golang-bookapi` | Start Jenkins container        |
| `docker exec -it jenkins-golang-bookapi /bin/bash`                                             | Access Jenkins container shell |
| `cat /var/jenkins_home/secrets/initialAdminPassword`                                           | Get admin login password       |
| `go mod tidy`                                                                                  | Install Go dependencies        |
| `go build -o bookapi_server .`                                                                 | Build the Go server binary     |
| `./bookapi_server`                                                                             | Run the Go server binary       |

---

### **Steps I followed to create this Jenkins Freestyle Project**

### **1. Dockerfile Configuration**

I created a `Dockerfile` in the project root to extend Jenkins, add Go, and configure dependencies for the **GoLang BookAPI** project.

#### **Dockerfile**:

```dockerfile
# Use official Jenkins LTS as base image
FROM jenkins/jenkins:lts

# Switch to root to install packages
USER root

# Install Go and other necessary dependencies
RUN apt-get update && \
    apt-get install -y golang-go git && \
    apt-get clean

# Set Go workspace environment variables
ENV GOPATH=/go
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

# Switch back to Jenkins user
USER jenkins
```

### **2. Build Docker Image**

I built the custom Jenkins image with Go and the required packages for the **GoLang BookAPI** project.

```bash
docker build -t jenkins-golang-bookapi .
```

### **3. Run Jenkins Container**

I ran the Jenkins container, mapping ports for the Jenkins UI and the Go application.

```bash
docker run -d -p 8090:8080 -p 8091:8081 --name jenkins-golang-bookapi jenkins-golang-bookapi
```

#### **Port Explanation**:

| Option                          | Purpose                                                      |
| ------------------------------- | ------------------------------------------------------------ |
| `-d`                            | Detached (background) mode                                   |
| `-p 8090:8080`                  | Jenkins UI on [http://localhost:8090](http://localhost:8090) |
| `-p 8091:8081`                  | Go server on [http://localhost:8091](http://localhost:8091)  |
| `--name jenkins-golang-bookapi` | Container name as `jenkins-golang-bookapi`                   |
| `jenkins-golang-bookapi`        | Uses the custom Jenkins image you built                      |

### **4. Get Jenkins Admin Password**

I retrieved the Jenkins admin password and pasted it into [http://localhost:8090](http://localhost:8090) to complete Jenkins setup (install plugins, create admin user).

```bash
docker exec -it jenkins-golang-bookapi cat /var/jenkins_home/secrets/initialAdminPassword
```

### **5. Configure Jenkins Job: BookAPI**

Followed these steps to configure the Jenkins job for the **GoLang BookAPI** project:

1. **Open Jenkins UI** at [http://localhost:8090](http://localhost:8090).

2. **Click** `New Item`.

3. **Enter name**: `GoBookAPI`.

4. **Choose** `Freestyle Project` → **Click OK**.

5. Under **Source Code Management**, select **Git** and provide the repository URL for the **GoLang BookAPI** project:

   ```bash
   https://github.com/your_username/bookapi.git
   ```

6. Under **Build Steps**, select **Execute shell** and add the following script:

```bash
cd /var/jenkins_home/workspace/GoBookAPI
go mod tidy
go build -o bookapi_server .
```

This script will:

* Install Go dependencies (`go mod tidy`)
* Build the Go binary (`go build`)

### **6. Run the Go Server in Jenkins Container**

After the build is complete, I ran the Go server binary from the Jenkins container:

```bash
docker exec -it jenkins-golang-bookapi /bin/bash
cd /var/jenkins_home/workspace/GoBookAPI
./bookapi_server
```

---

This version should work correctly without syntax errors! If you need any further clarification or adjustments, feel free to ask!
