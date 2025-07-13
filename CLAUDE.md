# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Hugo-based static site for the book companion website "A Pythonic Adventure" (pythonicadventure.com). The site uses the Hugo Relearn theme and is built with Go modules.

## Development Commands

### Building the Site
```bash
hugo
```

### Running Development Server
```bash
hugo server
```

### Building for Production
```bash
hugo --minify
```

## Architecture

### Site Structure
- `content/` - Markdown content files organized by sections
  - `_index.md` - Homepage content
  - `about/` - About page
  - `contact/` - Contact form page
  - `make-it-better/` - Feedback section
  - `projects/` - Projects showcase
  - `troubleshooting/` - Help section
- `layouts/partials/` - Custom HTML partials
  - `analytics.html` - Analytics tracking
  - `logo.html` - Site logo
  - `meta.html` - Meta tags
- `static/` - Static assets (images, icons)
- `config.toml` - Hugo configuration
- `public/` - Generated site output (auto-generated)

### Theme Configuration
- Uses Hugo Relearn theme via Go modules
- Configured with light/dark theme variants
- Theme path: `github.com/McShelby/hugo-theme-relearn`

### Content Management
- Content is written in Markdown
- Uses Hugo's content organization with sections
- Supports both light and dark themes
- Includes contact form functionality with reCAPTCHA v3

### Key Features
- Responsive design with Relearn theme
- Contact form with reCAPTCHA integration
- Multi-theme support (light/dark)
- Static site generation for fast performance
- Book companion content organization

## Development Notes

- Go 1.19+ required for module support
- Hugo theme is managed via Go modules
- Site builds to `public/` directory
- reCAPTCHA keys stored in `recaptcha/` directory (not committed)
- Base URL configured for pythonicadventure.com

## Form Processor Deployment (Cloud Run)

The `formprocessor/` directory contains a Go application that handles contact form submissions with reCAPTCHA v3 verification and email delivery via SMTP.

### Prerequisites

1. **Google Cloud CLI installed and authenticated**
2. **Google Cloud Project with billing enabled**
3. **reCAPTCHA v3 secret key** (stored as environment variable)
4. **SMTP credentials** for email delivery

### Deployment Steps

1. **Set the Google Cloud project:**
   ```bash
   gcloud config set project cloud-run-test-383922
   ```

2. **Enable required Google Cloud services:**
   ```bash
   gcloud services enable run.googleapis.com cloudbuild.googleapis.com
   ```

3. **Deploy the application to Cloud Run:**
   ```bash
   cd formprocessor
   gcloud run deploy --source .
   ```
   - Confirm service name: `formprocessor`
   - Select region: `us-east1`
   - Allow unauthenticated invocations: `y`

4. **Configure environment variables in Google Cloud Console:**
   - Navigate to Cloud Run > formprocessor service
   - Go to "Variables & Secrets" tab
   - Add environment variables:
     - `RECAPTCHA_SECRET`: Your reCAPTCHA v3 secret key
     - `SMTP_PASSWORD`: Gmail app password for email delivery

5. **Verify deployment:**
   - Service should show as "Protected" (green) in Cloud Console
   - Test the form submission functionality
   - Monitor logs for any errors

### Additional Deployment Files (Recommended)

Consider adding these files to `formprocessor/` for better deployment control:

**Dockerfile:**
```dockerfile
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/form.html .
EXPOSE 8080
CMD ["./main"]
```

**.gcloudignore:**
```
.git
.gitignore
README.md
Dockerfile
.dockerignore
```

### Security Notes

- reCAPTCHA secret key is stored as environment variable (not in code)
- SMTP password uses Gmail app password (not account password)
- Service runs with minimal privileges on Cloud Run
- HTTPS is enforced by Cloud Run