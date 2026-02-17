# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- HTML template rendering with `html/template` package
- Home page route (/) with Bootstrap-styled template
- About page route (/about) with Bootstrap-styled template
- Proper Go project layout with `cmd/` and `pkg/` directories
- `pkg/handlers` package for HTTP request handlers with Repository pattern
- `pkg/render` package for template rendering logic with caching
- `pkg/config` package for application configuration management
- Template caching system with `CreateTemplateCache()` function
- Base layout template system (`base.layout.tmpl`) with template blocks
- Template block architecture for content, CSS, and JS injection
- Repository pattern for handlers to access application configuration
- AppConfig struct to hold template cache, UseCache flag, and InfoLog

### Changed
- Refactored from single-file to standard Go project layout
- Moved main application to `cmd/web/main.go`
- Organized code into reusable packages (handlers, render, config)
- Replaced plain text response with HTML template rendering
- Added Bootstrap 5.3.8 CSS framework for styling
- Updated Dockerfile to build from `cmd/web/` directory
- Templates now use layout inheritance with `{{template "base" .}}` and `{{define "content"}}` blocks
- Handlers now use Repository pattern for accessing app configuration
- Render package now supports both cached and non-cached template modes
- Main function now initializes AppConfig, template cache, and repository before routing

### Deprecated

### Removed
- Kubernetes manifests (`k8s/` directory) - deployment.yaml, service.yaml, kustomization.yaml, README.md
- Direct template parsing in favor of template caching system

### Fixed

### Security

## [0.1.0] - 2026-02-11

### Added
- Initial release
- Basic "Hello, World!" HTTP server on port 8080
- Docker support with multi-stage builds
- Kubernetes deployment configuration for k3s
- Support for CloudFlare Tunnel integration

[Unreleased]: https://github.com/cpwu/nova/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/cpwu/nova/releases/tag/v0.1.0
