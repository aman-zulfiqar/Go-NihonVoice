# Security Policy — Google Text-to-Speech usage

This document describes recommended security practices for using Google Cloud Text-to-Speech (TTS) in this project.

## Summary
- **Never commit** service account keys, API keys, or any credentials to git.
- Use a **service account** with least privilege (only the permissions required to synthesize audio).
- Store credentials in a secure secret store (GitHub Secrets, Google Secret Manager, Vault).
- Prefer Google Cloud IAM + Workload Identity where possible (avoid long-lived JSON keys).
- Rotate credentials regularly and revoke unused/compromised keys immediately.

## Secrets & Credentials
- **Service account JSON**:
  - Do **not** commit to repo.
  - Store in: Google Secret Manager, HashiCorp Vault, or GitHub Actions Secrets.
  - In local dev, use `GOOGLE_APPLICATION_CREDENTIALS` to point to the JSON file (kept outside repo).
- **API keys** (if used): restrict by IP, referrer, and enable only Text-to-Speech API.
- **GitHub**: use repository or organization-level Secrets for CI (e.g. `GCP_SA_KEY`).

## IAM & Permissions (least privilege)
- Create a dedicated service account for this project.
- Grant only the minimal permission required to synthesize speech. If possible, create a custom role containing only `texttospeech.*synthesize*` permission(s) rather than broad roles.
- Avoid using owner/editor roles on the service account.

## Local development
- Keep local service account JSON files in a protected directory (not in git).
- Use `gcloud auth application-default login` for short-lived dev credentials where appropriate.

## CI / CD (example)
- Store service account JSON as an encrypted secret (e.g. `GCP_SA_KEY`).
- In GitHub Actions, write the JSON to a file at runtime and set `GOOGLE_APPLICATION_CREDENTIALS` before running tests or builds.

Example snippet for GitHub Actions:
```yaml
- name: Set up GCP credentials
  env:
    GCP_SA_KEY: ${{ secrets.GCP_SA_KEY }}
  run: |
    echo "$GCP_SA_KEY" > $HOME/gcp-sa.json
    export GOOGLE_APPLICATION_CREDENTIALS=$HOME/gcp-sa.json
    # run build/test steps here
