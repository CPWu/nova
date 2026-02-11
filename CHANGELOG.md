# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial Go web service with HTTP endpoint
- Multi-stage Dockerfile for optimized builds
- Multi-architecture build support (ARM64 + AMD64)
- Build script with registry push capability
- Kubernetes manifests (Deployment, Service, Kustomization)
- Health checks and readiness probes
- Resource limits optimized for Raspberry Pi
- Docker best practices documentation
- Development guide for AI assistants
- Comprehensive README

### Changed
- Updated Go version to 1.23 in go.mod and Dockerfile
- Updated k8s deployment to use chunw208/nova:latest image
- Made build.sh executable with proper permissions
- Improved build script Docker login check

### Deprecated

### Removed

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
