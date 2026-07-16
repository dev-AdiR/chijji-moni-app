## Prerequisites

- Go
- Docker
- Task (`go install github.com/go-task/task/v3/cmd/task@latest`)

## Local Development

### 1. Clone the repository

```bash
git clone https://github.com/dev-AdiR/chijji-moni-app.git
cd chijji-moni-app
```

### 2. Set up environment variables

Create a `.env` file:

```env
SUPABASE_URL=your_supabase_url
SUPABASE_SECRET_KEY=your_supabase_secret_key
JWT_SECRET=your_jwt_secret
```

### 3. Run locally with Docker Compose

```bash
docker compose up
```

## Deployment

Deployment is handled via GitHub Actions. To deploy a new version:

1. Go to **Actions** → **Build and Deploy to EKS**
2. Click **Run workflow**
3. Enter the version (e.g. `v1.0.3`)
4. Click **Run workflow**

This will:
- Build and push the Docker image to Docker Hub as `adityarawat1999/chijji-moni-app:v1.0.3`
- Trigger the infra repo to deploy it to EKS

## Required GitHub Secrets

| Secret | Description |
|---|---|
| `DOCKERHUB_USERNAME` | Docker Hub username |
| `DOCKERHUB_TOKEN` | Docker Hub access token |
| `INFRA_REPO_TOKEN` | GitHub token with access to the infra repo |

## Related

- [chijji-moni-infra](https://github.com/dev-AdiR/chijji-moni-infra) — Infrastructure and Kubernetes manifests