# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Nova is a Go web service with HTML template rendering using a template caching system and base layout inheritance. It serves multiple routes (/, /about) with Bootstrap-styled pages. The application follows the Repository pattern for clean architecture and uses application configuration management. The project is containerized using Docker with multi-stage builds for optimized image size.

## Build and Run Commands

### Local Development

Build and run locally:
```bash
go build -o nova cmd/web/main.go
./nova
```

Run without building:
```bash
go run cmd/web/main.go
```

### Docker

Build Docker image locally:
```bash
./build.sh
```

Build and push multi-architecture image (ARM64 + AMD64) to registry:
```bash
export DOCKER_USERNAME=your-dockerhub-username
./build.sh --push
```

Or manually:
```bash
docker build -t nova .
```

Run the container:
```bash
docker run -p 8080:8080 nova
```

### Testing the Service

Once running (locally or in Docker), test the endpoints:
```bash
# Home page
curl http://localhost:8080

# About page
curl http://localhost:8080/about
```

Or open in browser: `http://localhost:8080`

## Architecture

This follows the standard Go project layout with clean separation of concerns:

**Application Structure:**
- **cmd/web/main.go**: Application entry point with initialization and server setup
  - Initializes AppConfig with template cache, session manager, and production flag
  - Configures session with 24-hour lifetime and secure cookie settings
  - Creates Repository with app config
  - Sets up handlers and render package
  - Calls routes() function to get configured HTTP handler
  - Starts HTTP server
- **cmd/web/routes.go**: HTTP route definitions using Chi router
  - Uses `github.com/go-chi/chi/v5` for URL routing
  - Configures middleware pipeline (Recoverer, NoSurf CSRF, SessionLoad)
  - Defines GET routes for / and /about
  - Returns configured http.Handler with middleware
- **cmd/web/middleware.go**: Custom middleware functions
  - NoSurf: CSRF protection using `justinas/nosurf`
  - SessionLoad: Loads and saves session data
  - WriteToConsole: Custom request logging middleware
- **pkg/config/**: Application configuration package
  - **config.go**: AppConfig struct holding template cache, UseCache flag, InfoLog, InProduction flag, and Session manager
- **pkg/handlers/**: HTTP handlers package with Repository pattern
  - **handlers.go**: Repository struct and handler methods (Home, About)
  - Repository pattern provides handlers access to application configuration including session manager
  - Home handler stores remote IP address in session
  - About handler retrieves remote IP from session and displays it
  - Handlers create TemplateData instances and pass them to render package
- **pkg/models/**: Data structures for template rendering
  - **templatedata.go**: TemplateData struct with fields for StringMap, IntMap, FloatMap, Data, CSRFToken, Flash, Error, Warning
  - Used to pass data from handlers to templates
- **pkg/render/**: Template rendering package with caching system
  - **render.go**: RenderTemplate function and CreateTemplateCache function
  - Supports both cached (production) and non-cached (development) modes
  - Parses *.page.tmpl and *.layout.tmpl files into template cache
  - AddDefaultData helper function for adding default data to all templates
- **templates/**: HTML templates with Bootstrap 5 styling and layout inheritance
  - **base.layout.tmpl**: Base layout defining HTML structure with named blocks (content, css, js)
  - **home.page.tmpl**: Home page content that extends base layout
  - **about.page.tmpl**: About page content that extends base layout, displays data from StringMap

**Routes (defined in cmd/web/routes.go):**
- `/` - Home page with Bootstrap styling
- `/about` - About page with data passed via TemplateData

**Technical Details:**
- Server runs on port 8080
- Uses Go standard library (`net/http`, `html/template`)
- Uses Chi router (`github.com/go-chi/chi/v5`) for HTTP routing
- Middleware pipeline with panic recovery, CSRF protection, and session management
- Session management with `alexedwards/scs/v2` for state persistence across requests
- CSRF protection using `justinas/nosurf` on all routes
- Follows Go's standard project layout (cmd, pkg structure)
- Exported packages allow for easy testing and reusability
- Template caching system for improved performance
- Repository pattern for clean architecture and dependency injection
- Base layout template with block inheritance (content, css, js blocks)
- TemplateData struct for passing typed data to templates

**Template System:**
- **Layout Inheritance**: Page templates extend `base.layout.tmpl` using `{{template "base" .}}`
- **Named Blocks**: Define content blocks with `{{define "content"}}...{{end}}`
- **Caching**: Templates parsed at startup and stored in map[string]*template.Template
- **Cache Modes**:
  - `UseCache = false`: Parse templates on every request (development)
  - `UseCache = true`: Use pre-parsed cached templates (production)

**Repository Pattern and Data Flow:**
1. AppConfig created in main with template cache and session manager
2. Session configured with 24-hour lifetime and secure cookie settings
3. Repository struct wraps AppConfig
4. Handlers receive Repository to access app configuration and session
5. Render package initialized with AppConfig reference
6. Middleware pipeline processes requests (recovery, CSRF, session loading)
7. Handlers use session to store/retrieve data across requests
8. Handlers create TemplateData instances with data to pass to templates
9. TemplateData flows from handlers → render.RenderTemplate → templates
10. Templates access data via dot notation (e.g., `{{index .StringMap "test"}}`)

The Dockerfile implements a multi-stage build pattern:
1. **Build stage**: Uses `golang:1.23-alpine` to compile the Go binary from cmd/web/
2. **Runtime stage**: Uses `alpine:latest` with only the compiled binary for minimal image size

The build script supports multi-architecture builds (AMD64 + ARM64) for broad compatibility.

## Docker Best Practices

When modifying Dockerfiles in this repository, refer to `DOCKER.md` for detailed guidance on:
- Multi-stage builds for size optimization
- Base image selection (prefer Alpine variants)
- Layer caching strategies
- Security practices (non-root users, vulnerability scanning)
- Go-specific optimizations (CGO_ENABLED=0, build flags)
- Health checks for container orchestration

Key optimization for Go builds:
```dockerfile
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o app
```

## Dependencies

Go version: 1.23 (specified in go.mod)
Module path: github.com/cpwu/nova
Docker Hub: chunw208/nova

**External Dependencies:**
- `github.com/go-chi/chi/v5` v5.2.5 - HTTP routing with middleware support
- `github.com/alexedwards/scs/v2` v2.9.0 - Session management
- `github.com/justinas/nosurf` v1.2.0 - CSRF protection middleware
