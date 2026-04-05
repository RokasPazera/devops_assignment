# NSS Thermometer – DevOps Documentation

## Part 1: Development and Quality

### 1. Git Process

I used GitLab issues to track tasks and created a separate branch for each feature or fix. When a feature was done, I opened a merge request to merge it into `main`. I did this because it lets me review code before it goes into the main branch, which prevents mistakes from being pushed directly.

I also added a *pre-push* git hook (`.git/hooks/pre-push`) that runs *gosec* automatically before every push. I used this because the assignment asked for security checks before pushing. If gosec finds problems, the push is blocked. I left the known SQL injection in the code on purpose (as the assignment required) and used `--no-verify` when I needed to push it anyway.

### 2. Applications
The frontend is Angular and is served by nginx. I noticed the API URL was hardcoded to `localhost:3000`, so I refactored it to use Angular environment files. The local `environment.ts` keeps `localhost:3000` so Docker Compose still works, and the CI/CD pipeline generates `environment.aws.ts` with the actual ALB URL at build time.

### 3. Seed Script

I updated `add_data.sh` to be more robust. I added `set -euo pipefail` at the top so the script stops immediately if something goes wrong instead of silently continuing. I also added checks at the start to make sure the data files exist before trying to read them — if they don't exist the script prints a clear error and exits.

I made the script check the HTTP status code from each curl request. If a POST fails, it prints an error message and counts the failure. At the end it prints how many goals succeeded and failed, and exits with code 1 if anything failed.

I did all of this because a good script should tell you clearly when something goes wrong, rather than silently failing.

### 4. Docker

I containerised all four components: backend, frontend, database, and seed script.

I used multi-stage builds for the backend and frontend to keep the images small:
- **Backend**: I compile the Go binary in a `golang:alpine` image, then copy only the binary into a clean `alpine` image. This way the final image has no Go compiler or source code in it.
- **Frontend**: I build the Angular app in a `node:alpine` image, then copy only the built files into an `nginx:alpine` image.
- **Seed**: A simple `alpine` image with only bash and curl installed.
- **Database**: I used the official Microsoft SQL Server 2022 image.

I added `.dockerignore` files to each service so Docker doesn't include unnecessary files in the build context (like `node_modules`, test files, or git history). This makes builds faster and images smaller.

I added health checks to `docker-compose.yml` so containers start in the right order. The database health check runs a simple SQL query to confirm MSSQL is ready. The backend only starts after the database is healthy, and the frontend and seed only start after the backend is healthy. I did this because without health checks, the backend would try to connect to the database before it was ready and crash.

### 5. Continuous Integration

I built a CI pipeline in `.gitlab-ci.yml` with four stages: `test`, `pages`, `build`, and `deploy`.

All test jobs run in parallel to save time:
- `backend:test` — runs Go unit tests and generates a coverage report
- `backend:lint` — runs golangci-lint
- `backend:security` — runs gosec; I set `allow_failure: true` because I kept the SQL injection vulnerability in the code intentionally (as the assignment required), so the pipeline can continue even when this job fails
- `frontend:build` — builds the Angular app; if `API_URL` is set as a GitLab variable, it injects the ALB URL into the production environment file before building
- `frontend:test` — runs Karma unit tests headlessly in Chrome

The `pages` stage publishes the coverage report and test report to GitLab Pages.

I added caching for Go modules and npm `node_modules` so repeated pipeline runs are faster.

The build and deploy stages only run on git tags, because the assignment requires deployment to be triggered by a tag.

---

## Part 2: Deployment

### 6 & 7. Terraform and CI/CD Deployment

I wrote the infrastructure in Terraform (`infra/main.tf`). I chose `us-east-1` as the region.

The infrastructure includes:
- A VPC with two public subnets in different availability zones
- An Internet Gateway and route table
- Security groups for the ALB, backend, and database
- An EC2 instance running MSSQL in Docker (the database)
- An Application Load Balancer that routes traffic to the backend
- An Auto Scaling Group for the backend EC2 instances
- An S3 bucket for the frontend static website

I kept sensitive values like the MSSQL password and AWS credentials out of git by storing them in `terraform.tfvars` (which is in `.gitignore`) and as masked GitLab CI variables.

For deployment, the pipeline uses AWS CLI to get the current list of backend instance IPs from the ASG, then SSH's into each one to pull the latest Docker image and restart the container. I stored the SSH private key as a protected, masked GitLab variable so it never appears in logs.

The database had to be created manually the first time with `CREATE DATABASE thermometer` because MSSQL doesn't create it automatically.

### 8. High Availability

I made the backend highly available by spreading it across two availability zones (`us-east-1a` and `us-east-1b`). The ALB routes traffic between the instances and only sends requests to healthy ones. I configured ELB health checks on the ASG so if an instance fails the `/api/goals` health check, it gets replaced automatically.

---

### 9. AI

I used AI to check for examples on how to connect pipelines with terraform and which variables I need exactly to store in gitlab, so that it would work.