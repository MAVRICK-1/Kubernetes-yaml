Jenkins provides a rich set of commands and steps for declarative and scripted pipelines. Below is a categorized list of commonly used commands with explanations:

---

### **Pipeline Syntax**
1. **`pipeline {}`**  
   Defines the declarative pipeline structure.

2. **`agent {}`**  
   Specifies where the pipeline or stage runs. Example:
   - `agent any` - Runs on any available agent.
   - `agent { label 'my-node' }` - Runs on an agent with a specific label.

3. **`stages {}`**  
   Groups multiple stages of the pipeline.

4. **`stage('Name') {}`**  
   Defines a stage within the pipeline.

5. **`steps {}`**  
   Contains the steps to execute in a stage.

---

### **General Steps**
1. **`echo 'message'`**  
   Prints a message to the console output.

2. **`git`**  
   Clones a Git repository. Example:
   ```groovy
   git branch: 'main', url: 'https://github.com/user/repo.git'
   ```

3. **`sh 'command'`**  
   Runs a shell command (for Linux/macOS). Example:
   ```groovy
   sh 'ls -la'
   ```

4. **`bat 'command'`**  
   Runs a batch command (for Windows). Example:
   ```groovy
   bat 'dir'
   ```

5. **`input`**  
   Pauses the pipeline for manual input. Example:
   ```groovy
   input 'Proceed to deploy?'
   ```

6. **`timeout(time: X, unit: 'UNIT') {}`**  
   Sets a timeout for the enclosed block. Example:
   ```groovy
   timeout(time: 10, unit: 'MINUTES') {
       sh 'some-long-task'
   }
   ```

---

### **Post Actions**
1. **`post {}`**  
   Defines actions to run after a stage or pipeline, like:
   - `always` - Runs regardless of success or failure.
   - `success` - Runs only on success.
   - `failure` - Runs only on failure.

   Example:
   ```groovy
   post {
       always {
           echo 'Pipeline completed'
       }
       failure {
           mail to: 'team@example.com', subject: 'Build Failed', body: 'Check Jenkins!'
       }
   }
   ```

---

### **Environment Management**
1. **`environment {}`**  
   Defines environment variables. Example:
   ```groovy
   environment {
       VAR_NAME = 'value'
   }
   ```

2. **`withEnv(['VAR=value']) {}`**  
   Temporarily sets environment variables for a block. Example:
   ```groovy
   withEnv(['PATH+EXTRA=/usr/local/bin']) {
       sh 'run-my-script'
   }
   ```

---

### **Tools**
1. **`tool 'tool-name'`**  
   Configures a tool installed in Jenkins (e.g., Maven, JDK). Example:
   ```groovy
   def mvnHome = tool name: 'Maven 3.6.3'
   sh "${mvnHome}/bin/mvn clean install"
   ```

---

### **Archiving and Reporting**
1. **`archiveArtifacts 'pattern'`**  
   Archives files for later retrieval. Example:
   ```groovy
   archiveArtifacts artifacts: '**/target/*.jar'
   ```

2. **`junit 'pattern'`**  
   Publishes JUnit test results. Example:
   ```groovy
   junit '**/test-results.xml'
   ```

3. **`stash` and `unstash`**  
   Saves and restores files between stages. Example:
   ```groovy
   stash name: 'source', includes: '**/*'
   unstash 'source'
   ```

---

### **Parallel Execution**
1. **`parallel`**  
   Runs multiple branches concurrently. Example:
   ```groovy
   parallel(
       branch1: {
           echo 'Running branch 1'
       },
       branch2: {
           echo 'Running branch 2'
       }
   )
   ```

---

### **Build Control**
1. **`build`**  
   Triggers another job. Example:
   ```groovy
   build job: 'another-job', parameters: [string(name: 'param', value: 'value')]
   ```

2. **`retry`**  
   Retries a block on failure. Example:
   ```groovy
   retry(3) {
       sh 'unstable-command'
   }
   ```

---

### **Notifications**
1. **`mail`**  
   Sends an email. Example:
   ```groovy
   mail to: 'team@example.com', subject: 'Build Success', body: 'Check the build logs.'
   ```

2. **`slackSend`**  
   Sends a message to Slack (requires plugin). Example:
   ```groovy
   slackSend channel: '#alerts', message: 'Build failed!'
   ```

---

### **Credentials**
1. **`withCredentials`**  
   Uses Jenkins credentials in a block. Example:
   ```groovy
   withCredentials([usernamePassword(credentialsId: 'my-creds', usernameVariable: 'USER', passwordVariable: 'PASS')]) {
       sh 'curl -u $USER:$PASS https://example.com'
   }
   ```

---

### **Clean Workspace**
1. **`deleteDir()`**  
   Deletes the current workspace.

2. **`cleanWs()`**  
   Cleans the workspace (requires the Workspace Cleanup plugin).

---

### **Conditional Execution**
1. **`when {}`**  
   Adds conditions for stages. Example:
   ```groovy
   stage('Deploy') {
       when {
           branch 'main'
       }
       steps {
           echo 'Deploying to production'
       }
   }
   ```

2. **`if/else`**  
   Scripted pipelines support traditional `if/else`. Example:
   ```groovy
   if (env.BRANCH_NAME == 'main') {
       sh 'deploy-production.sh'
   } else {
       sh 'deploy-staging.sh'
   }
   ```

---

### **Debugging**
1. **`currentBuild.result`**  
   Gets or sets the build result (`SUCCESS`, `FAILURE`, etc.).

2. **`error 'message'`**  
   Fails the build with a message.

3. **`script {}`**  
   Runs Groovy code for advanced logic. Example:
   ```groovy
   script {
       def myVar = 'value'
       echo "Value is: ${myVar}"
   }
   ```

---

Let me know if you'd like a deeper explanation of any specific command or feature!
