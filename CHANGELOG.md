# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Static file serving route (`/static/*`) for serving images and other static assets
- Static directory structure with `static/images/` for image assets
- Image display on home page template

### Changed
- Updated `cmd/web/routes.go` to include file server handler for static assets

### Deprecated

### Removed

### Fixed

### Security

## [0.3.0] - 2026-02-23

### Added
- Session management using `alexedwards/scs/v2` package
- Session configuration with 24-hour lifetime and secure cookie settings
- CSRF protection middleware using `justinas/nosurf` package
- Custom request logging middleware (`WriteToConsole`)
- Middleware pipeline with panic recovery, session loading, and CSRF protection
- Session-based remote IP tracking (stored on home page visit, displayed on about page)
- Conditional display of remote IP address in about page template
- `cmd/web/middleware.go` file for custom middleware functions

### Changed
- Replaced Pat router with Chi router (`github.com/go-chi/chi/v5`)
- Simplified handler registration (removed need for `http.HandlerFunc` wrapper)
- Updated dependencies: removed Pat, added Chi v5, nosurf, and scs/v2
- Enhanced AppConfig with InProduction flag and Session manager
- Home handler now stores remote IP in session
- About handler now retrieves and displays remote IP from session

### Removed
- `github.com/bmizerany/pat` router dependency

### Fixed
- Template syntax error in about.page.tmpl (missing space between `index` and `.StringMap`)
- Routing bug where About handler was incorrectly mapped to "/" instead of "/about"

### Security
- Added CSRF protection to all routes with secure cookie configuration
- Session cookies configured with SameSite=Lax and Secure flag for production

## [0.2.0] - 2026-02-19

### Added
- HTML template rendering with `html/template` package
- Home page route (/) with Bootstrap-styled template
- About page route (/about) with Bootstrap-styled template
- Proper Go project layout with `cmd/` and `pkg/` directories
- `pkg/handlers` package for HTTP request handlers with Repository pattern
- `pkg/render` package for template rendering logic with caching
- `pkg/config` package for application configuration management
- `pkg/models` package with TemplateData struct for passing data to templates
- `cmd/web/routes.go` for centralized route definitions
- Pat router (`github.com/bmizerany/pat`) for cleaner URL routing
- Template caching system with `CreateTemplateCache()` function
- Base layout template system (`base.layout.tmpl`) with template blocks
- Template block architecture for content, CSS, and JS injection
- Repository pattern for handlers to access application configuration
- AppConfig struct to hold template cache, UseCache flag, and InfoLog
- TemplateData struct with StringMap, IntMap, FloatMap, Data, CSRFToken, Flash, Error, and Warning fields
- Data passing from handlers to templates via TemplateData
- Enhanced build.sh with auto-detection of Docker Hub username
- Build script validation for Docker daemon and username requirements

### Changed
- Refactored from single-file to standard Go project layout
- Moved main application to `cmd/web/main.go`
- Organized code into reusable packages (handlers, render, config, models)
- Replaced plain text response with HTML template rendering
- Added Bootstrap 5.3.8 CSS framework for styling
- Updated Dockerfile to build from `cmd/web/` directory
- Templates now use layout inheritance with `{{template "base" .}}` and `{{define "content"}}` blocks
- Handlers now use Repository pattern for accessing app configuration
- Render package now supports both cached and non-cached template modes
- Main function now initializes AppConfig, template cache, and repository before routing
- Separated routing logic from main.go into routes.go for better organization
- About handler now demonstrates data passing with StringMap
- About template now displays dynamic data from template context
- Improved build.sh with better error messages and user guidance

### Removed
- Kubernetes manifests (`k8s/` directory) - deployment.yaml, service.yaml, kustomization.yaml, README.md
- Direct template parsing in favor of template caching system

## [0.1.0] - 2026-02-11

### Added
- Initial release
- Basic "Hello, World!" HTTP server on port 8080
- Docker support with multi-stage builds
- Kubernetes deployment configuration for k3s
- Support for CloudFlare Tunnel integration

[Unreleased]: https://github.com/cpwu/nova/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/cpwu/nova/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/cpwu/nova/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/cpwu/nova/releases/tag/v0.1.0
