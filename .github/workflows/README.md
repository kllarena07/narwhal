# GitHub Actions Workflows

This directory contains GitHub Actions workflows for the Narwhal project.

## Workflows

### `ci.yml`
Main CI workflow that runs on every push and pull request. Runs both backend and frontend tests in parallel.

### `backend.yml`
Go backend-specific CI workflow. Runs tests, builds the orchestrator, checks code formatting, and runs `go vet`.

### `frontend.yml`
Frontend-specific CI workflow. Runs linting, TypeScript checks, and builds the Next.js application.

### `python-backend.yml`
Python backend CI workflow (for server.py). Runs code formatting checks, import sorting, linting, and type checking.

## Dependabot

Dependabot is configured via `.github/dependabot.yml` to automatically create pull requests for dependency updates:
- Go modules (weekly)
- npm/pnpm packages (weekly)
- GitHub Actions (weekly)

## Running Locally

To test workflows locally, you can use [act](https://github.com/nektos/act):

```bash
# Install act
brew install act  # macOS
# or
curl https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Run a workflow
act -j backend
act -j frontend
```
