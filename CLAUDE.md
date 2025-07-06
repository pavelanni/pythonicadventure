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