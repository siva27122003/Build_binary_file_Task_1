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

