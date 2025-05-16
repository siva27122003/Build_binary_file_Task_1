###  Faced Issues and Fixes

####  **Error 1: Outdated Code Used in Jenkins Build**

* **Issue:** I noticed that Jenkins was building an older version of my code, even though I had recently committed updates.
* **What I did:** I realized I had provided the GitHub repository in both the Jenkins SCM configuration and inside the `Jenkinsfile` (`git` command in the Checkout stage).
* **Fix:** I removed the `git` checkout step from the `Jenkinsfile` and allowed Jenkins to handle source code checkout using its SCM configuration only. This ensured Jenkins always pulled the latest commit.

---

####  **Error 2: Workspace Was Empty After `cleanWs()`**

* **Issue:** After running the pipeline, I found that all my files were missing — the workspace was completely empty when I used the `ls` command.
* **What I did:** I investigated and saw that `cleanWs()` was called inside the same stage as `git`, causing it to delete files after pulling the repo.
* **Fix:** I moved `cleanWs()` into a separate stage before the checkout. This ensured the workspace was cleaned first, then the latest code was checked out properly.

---

####  **Error 3: Changes to `main.go` Weren’t Reflected in Workspace**

* **Issue:** After building the project, I noticed that the `main.go` file still showed old code, not the latest version I had pushed.
* **What I did:** I checked and confirmed the `Jenkinsfile` had its own Git logic along with the Jenkins SCM setup.
* **Fix:** I removed the Git command from the pipeline script and trusted Jenkins SCM entirely. After this, the latest version of `main.go` was pulled and reflected correctly in the build.

---
###  Pipeline Task with golang

```markdown
#  Go Binary Build and CI/CD Pipeline with Jenkins

This repository presents an end-to-end Continuous Integration and Continuous Deployment (CI/CD) pipeline for building and packaging a Go-based application. The pipeline is implemented using Jenkins and includes source code management, module initialization, compilation, testing, and artifact archiving.

---

##  Repository Structure

```

Build\_binary\_file\_Task\_1/
├── Docker\_Artifact/            # (Optional: Artifact management)
├── Jenkins\_pipeline/
│   └── src/
│       ├── Jenkinsfile         # Jenkins declarative pipeline definition
│       ├── main.go             # Go source code (entry point)
│       ├── Testing.go          # Unit tests
│       └── go.mod              # Go module file

````

---

##  CI/CD Pipeline Overview

The Jenkinsfile implements a multi-stage declarative pipeline to automate the build and deployment workflow:

### 1. **Prepare Modules**
- Initializes the Go module if `go.mod` does not exist.
- Executes `go mod tidy` to clean and synchronize module dependencies.

### 2. **Build**
- Compiles the Go source code into a standalone binary named `main`.

### 3. **Run Binary**
- Executes the compiled binary with a sample argument (phone number) to validate runtime behavior.

### 4. **Unit Testing**
- Runs Go unit tests defined in `Testing.go` to verify functionality.

### 5. **Zip Build Output**
- Installs the `zip` package (if required).
- Compresses the `main` binary into a zip archive for distribution.

### 6. **Archive Artifacts**
- Archives the zipped binary (`phone_number_validate.zip`) as a Jenkins build artifact.

---

##  Technologies Used

- **Go (Golang)**: For application development and testing.
- **Jenkins**: For CI/CD automation.
- **Shell scripting**: Used within pipeline stages for setup and execution.
- **Zip utility**: For packaging build outputs.

---

##  How to Run Locally

To replicate the build process manually:

```bash
cd Jenkins_pipeline/src
go mod init buildbinarytask     # Skip if go.mod exists
go mod tidy
go build -o main main.go
chmod +x main
./main 9876543210               # Sample input
go test ./...                   # Run all tests
````

To package:

```bash
zip phone_number_validate.zip main
```

---

## 📦 Output Artifact

After successful execution, Jenkins will archive:

```
artifact_output/phone_number_validate.zip
```

This archive contains the compiled Go binary ready for deployment or distribution.

---






