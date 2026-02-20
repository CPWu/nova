# Nova

A lightweight Go web service with HTML template rendering. Features multiple routes and Bootstrap-styled pages. Designed for containerized deployment on Kubernetes clusters, with optimized builds for ARM64 (Raspberry Pi) and AMD64 architectures.

## Features

- HTTP server with multiple routes (/, /about)
- HTML template rendering with Bootstrap 5 styling
- Template caching system for improved performance
- Base layout template with block inheritance
- Repository pattern for clean architecture
- Application configuration management
- Clean separation of concerns (config, handlers, rendering, templates)
- Multi-stage Docker builds for optimized image size
- Multi-architecture support (ARM64 + AMD64)

## Prerequisites

- Go 1.23 or later
- Docker (with buildx for multi-arch builds)
- Docker Hub account (optional, for registry push)

## Quick Start

### Local Development

Run directly with Go:
```bash
go run cmd/web/main.go
```

Or build and run:
```bash
go build -o nova cmd/web/main.go
./nova
```

Test the endpoints:
```bash
# Home page
curl http://localhost:8080

# About page
curl http://localhost:8080/about
```

Or open in browser: `http://localhost:8080`

### Docker

Build and run locally:
```bash
./build.sh
docker run -p 8080:8080 nova
```

Build and push to registry (multi-arch):
```bash
export DOCKER_USERNAME=your-dockerhub-username
docker login
./build.sh --push
```

## Project Structure

```
nova/
├── cmd/
│   └── web/
│       ├── main.go          # Application entry point and initialization
│       └── routes.go        # HTTP route definitions using Pat router
├── pkg/
│   ├── config/
│   │   └── config.go        # Application configuration (AppConfig)
│   ├── handlers/
│   │   └── handlers.go      # HTTP request handlers with Repository pattern
│   ├── models/
│   │   └── templatedata.go  # Data structures for template rendering
│   └── render/
│       └── render.go        # Template rendering with caching system
├── templates/               # HTML templates
│   ├── base.layout.tmpl     # Base layout with template blocks
│   ├── home.page.tmpl       # Home page content template
│   └── about.page.tmpl      # About page content template
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksums
├── Dockerfile               # Multi-stage Docker build
├── build.sh                 # Build script with multi-arch support
├── DOCKER.md                # Docker best practices guide
├── CLAUDE.md                # AI assistant guidance
├── CHANGELOG.md             # Version history
└── README.md                # This file
```

## Architecture

### Routing
The application uses the [Pat router](https://github.com/bmizerany/pat) for clean URL routing:
- Routes are defined in `cmd/web/routes.go`
- Supports pattern-based routing with simple syntax
- Integrated with the Repository pattern for handler access

### Template System
The application uses a template caching system with layout inheritance:
- **Base Layout** (`base.layout.tmpl`): Defines the HTML structure with named blocks (content, css, js)
- **Page Templates** (`*.page.tmpl`): Define content that fills the blocks in the base layout
- **Template Cache**: Parses all templates at startup and stores them in memory for fast rendering
- **Template Data**: Uses `TemplateData` struct from `pkg/models` to pass data to templates

### Repository Pattern
Handlers use the Repository pattern to access application configuration:
- **AppConfig**: Holds template cache, cache mode flag, and loggers
- **Repository**: Wraps AppConfig and provides it to handlers
- **Handlers**: Access configuration through the Repository
- **Template Data Flow**: Handlers create `TemplateData` instances and pass them to the render package

### Configuration
In `cmd/web/main.go`, set `app.UseCache` to control template caching:
- `UseCache = false`: Templates are parsed on every request (development mode)
- `UseCache = true`: Templates are cached at startup (production mode)

## Development

See [CLAUDE.md](CLAUDE.md) for detailed development instructions and architecture overview.

See [DOCKER.md](DOCKER.md) for Docker best practices and optimization techniques.

See [CHANGELOG.md](CHANGELOG.md) for version history and recent changes.

## Contributing

This is a personal project. Feel free to fork and modify for your own use.

## License

MIT
