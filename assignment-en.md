# Exam Assignment DevOps: NSS Thermometer
_March 2026_

The National Student Survey is a large-scale national satisfaction survey in which students share their opinions about their HBO program and educational institution. You have probably filled in this survey yourself at some point.

HBO-IT uses various systems to promote the NSS. One of these systems is the NSS Stretch Goal Thermometer. This is an application displayed on various screens within Saxion that shows the current number of votes and the achieved stretch goals.

In this exam assignment, your task is to get this application running properly in the cloud. Everything must be automated so that your application can easily be deployed from GitLab to a scalable cloud-based solution.

### Prerequisite Requirements

-   This course concludes with an assessment. You can only participate in the assessment once all your homework assignments have been signed off. You are of course welcome to make a start on this exam assignment in the meantime. The retake assignment is a different assignment, so there is no point in working on this final assignment if you are not going to complete your homework on time.
 
-   For each part of the assignment (see below), you explain how you solved the problem and why you solved it in this way. Write this in your own words in a way that is understandable to the lecturer assessing your work. You work out your documentation in a Markdown file in your repository. You may also make use of comments in the source code (in this case, refer to the locations of these comments in the Markdown file).

-   Your documentation is **at most** 1800 words (2 A4 pages).

-   Make sure you use a repository in the Saxion GitLab organisation, created via [https://repo.hboictlab.nl](https://repo.hboictlab.nl/).

-   Upload a zipped export of your GitLab project to Brightspace (Settings → General → Advanced → Export Project).

    **Also include the URL of your repository in the comments field on Brightspace.**

### Assessment Rules

-   The deadline for submitting the assignment is **Monday 6 April 23:59 (beginnen of week 3.9)**.

-   The assignment is completed in pairs of two students who have registered with the lecturer. Groups will be formed at the end of week 6, with the requirement that both group members have signed off an equal number of practice assignments (maximum 1 week difference).

-   Your grade is determined based on the assessment that takes place after you submit your assignment. Make sure you thoroughly understand all the parts you have completed. Theoretical questions about your solution will also be asked.

-   You only receive points for parts of your work that you can explain during the assessment. For example: if your code works but you cannot explain what it does or why it works that way, that part of the assignment will not count towards your final grade.

-   It is **not** permitted to receive help from anyone outside your assessment group, except from your lecturer. It is also **not** permitted to provide help to anyone else.

-   You may use online posts, articles, tutorials, books, and videos. You must [add references](https://saxionbibliotheek.libguides.com/c.php?g=707376&p=5098767) for all sources and code fragments used in your text/code.

-   You may use Generative AI transparently and responsibly in accordance with the [Saxion AI usage guide for students](https://saxion.sharepoint.com/sites/o365-kpf-000001/Gedeelde%20%20documenten/Forms/AllItems.aspx?id=%2Fsites%2Fo365%2Dkpf%2D000001%2FGedeelde%20%20documenten%2FGenAI%20student%20guide%2Epdf&parent=%2Fsites%2Fo365%2Dkpf%2D000001%2FGedeelde%20%20documenten). Your work must include a justification of your AI usage in your process according to APA guidelines. You remain solely responsible for the content of your work. You must clearly indicate which prompts you used and which parts of the generated response you actually adopted or used.

## Exam Assignment

### Part 1: Development and Quality

#### 1. Setting up the Git process
During the assignment, we expect you to follow a solid software development process, as covered in week 1. This includes things such as tracking tasks using GitLab issues, using Git to collaborate on code, and using branches and merge requests to conduct code reviews. All of this should be clearly evident from your GitLab repository.

Additionally, you may consider professionalising your workflow further by making use of:

- Rebasing to combine unnecessary commits (commit squashing).
- Git hooks (for security checks or other components you want to verify locally, for example [GitLeaks](https://github.com/gitleaks/gitleaks)).
- Standardised commit messages, for example using [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/).

#### 2. Studying the applications and getting them running locally
In the repository you will find a folder containing a backend application. This is a Go application provided by us that returns data in JSON format. The `README.md` in the backend folder contains more information about the project. Install the dependencies and see if you can get the application running locally. Also run the other commands from the `README.md` for testing, linting, security checking, etc.

To verify that the application is working correctly, you can call the API endpoints using Postman (or another tool of your choice). The file `endpoints.postman_collection.json` can be imported into Postman and contains all the API endpoints.

**Tip:** The current implementation of the backend contains a security vulnerability (SQL injection). Fix this issue, but keep the old code as well. Later in the pipeline phase, we also want to be able to test that the pipeline fails when a security vulnerability is present.

In the frontend folder you will find our frontend application, built with Angular. Study the `README.md` and see if you can install and run the application. Also check whether you can run the various `npm run` commands.

Take the time to study the backend and frontend code thoroughly. Understand which languages are used. This will help you make the right choices later when working with Docker, CI/CD, and deployment. For example:

- Do you need a development environment for all actions?
- Is the language compiled or interpreted?
- How do you add environment variables?
- etc.

#### 3. Seed script
When you run the backend application and navigate to [http://localhost:8000/api/goals](http://localhost:8000/api/goals), you will see a list of stretch goals currently present in the database. Initially, the database is empty.

The `seed-application` folder contains data files that need to be read and added to the backend via a POST request. The data format is explained in the `README.md` in the `seed-application` folder.

Update the Bash script `add_data.sh` in the `seed-application` folder, which reads the data and uses `curl` to add the stretch goals to the backend (also known as "seeding").

**Tips:**
- Make sure you set the content type to `application/json` in your curl command.
- Make your script dynamic so that files with different names can easily be read in, by passing the filename as a parameter to your script.
- Make your script configurable by using environment variables.
- Ensure proper error handling.

#### 4. Containerising the applications

The next step is to containerise the entire application using Docker and Docker Compose. The goal is for all components (frontend, backend, database, and seed script) to run in containers, so that you do not need to install any language, runtime, or framework on your own machine.

The `README.md` files for the frontend and backend describe the available configuration options.

We want to achieve the following:
- Separate containers for the backend, frontend, database, and seed script.
- Use Microsoft SQL Server as the database. The 2022 version is sufficient for your solution. The backend's `README.md` explains which variables you need to configure the database connection.
- Use Nginx or Apache HTTP Server as the web server for your frontend. You do **not** need to add a reverse proxy, load balancer, or HTTPS.
- The seed script should run after the backend and database have started.
- Keep the images as small as possible. Choose appropriate base images and explore other ways to reduce image size.
- Keep the configuration as flexible as possible so that parts of the Docker Compose setup can be reused later when deploying to the cloud. Think about environment variables, .env files, dockerignore, multi-stage builds, Docker secrets, etc.

The intention is to let Docker do all the work, but if you need to perform multiple steps locally before you can start the composition (`docker compose up`), elaborate these steps in a shell script and include the script in your repository.

#### 5. Continuous Integration

Implement a development pipeline with Continuous Integration for the entire application. We want the CI portion to run after every **push** or **merge request** on the **main branch**. Your CI pipeline should consist of several steps:

- Backend:
    - Install dependencies, test, lint, security check, and code coverage.
    - The code coverage website must be served as GitLab Pages.
- Frontend:
    - Install dependencies, build the frontend, and run tests.
    - The build must be available as a downloadable artefact in GitLab.
    - The test report must be visible via GitLab Pages.

Try to make the pipeline as efficient as possible by running jobs in parallel and using caching. Also test whether the security check works correctly (see step 1 of the exam assignment).

### Part 2: Deployment
Now we will make the application available using AWS. We recommend building this up step by step, starting simple (without load balancers, auto scaling, etc.). In steps 6 and 7, you will build this simple architecture using Terraform and GitLab CI/CD. Then in step 8, you will expand your architecture to make it *highly available*.

Make sure you create a diagram of your AWS architecture and include it in your documentation. Use [draw.io](https://draw.io/) for this and use the AWS 2026 shapes (click *+ More Shapes > AWS 2026*).

#### 6. Create an infrastructure to run the application on
Create a suitable infrastructure for the application in AWS using Terraform. The infrastructure does not (yet) need to be *highly available* or scalable. If you cannot manage this with Terraform, you may create the infrastructure manually (you will of course lose points for this). Your Terraform configuration must be placed in the provided `infra` folder.

In this step, make sure the application runs on the AWS architecture. The application does not yet need to be deployed with Terraform to the newly created infrastructure. You will do this in step 7 using Continuous Deployment.

#### 7. Automate the deployment of the applications
Automate the deployment of the backend so that when we add a new tag to the main branch in GitLab, the GitLab CI:

-   Builds the necessary Docker images for the backend
-   Uploads the image(s) to the [private container registry](https://docs.gitlab.com/ee/user/packages/container_registry/index.html) of your GitLab repository
-   Uses the images in AWS so that the new version of the backend runs in the cloud
-   Uploads the built frontend web files as a static website (for example in S3 or on an EC2 instance)

There are various ways to deploy the backend to your cloud environment, for example via SSH. Make sure that secrets (such as SSH private keys) are stored securely.

**Tips:**

- If you use SSH to automatically deploy the backend to a server, the client will check whether the server is a known host. To automate this, you can store the **public key** of the server as an environment variable in your CI/CD environment and add it to the `known_hosts` file:

    ```bash
    echo "$SSH_KNOWN_HOST" >> ~/.ssh/known_hosts
    ```

- Retrieve the public key once by running the following command locally (can be done from any machine with `ssh-keyscan`) and save the output as a variable:

    ```bash
    ssh-keyscan -H <your-server-ip>
    ```

**Additional ideas:**
- You can include manual steps in your pipeline to create or tear down the infrastructure.

### 8. Make the backend highly available

Extend your Terraform configuration with a highly available and scalable backend. If your frontend runs on EC2, extend the scalability there as well. Don't forget to update the frontend to connect to your new infrastructure. To validate whether your infrastructure is correctly configured, verify that your frontend displays the stretch goals.

Requirements:
- The number of backend machines in use scales automatically under increased load.
- There are machines in multiple availability zones handling the requests.

Don't forget to update your architecture diagram so that the final architecture is included in your documentation.
