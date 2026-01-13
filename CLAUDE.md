# Ron's Designs and Signs - Go Web Application

Custom signs, decals, and apparel business website built with Go + Templ + HTMX + Tailwind CSS.

## Critical: Check Build Logs

**ALWAYS check `./tmp/air-combined.log` after making code changes.** This log contains:
- Compilation errors
- Template generation errors
- Build failures

Never trust that a build succeeded without checking this log.

## Development Workflow

During development, `make dev` is always running and handles everything automatically:
- Kills any existing process on port 3000
- Regenerates Templ templates
- Runs `go mod tidy`
- Rebuilds and restarts the server

**You do NOT need to manually run:** `templ generate`, `go build`, or `air`

For Tailwind CSS changes, run in a separate terminal:
```bash
make css-watch
```

## Environment Variables

All configuration via `.envrc` (load with `direnv allow` or `source .envrc`):

| Variable | Description | Default |
|----------|-------------|---------|
| DATABASE_URL | SQLite database path | ./data/ronsdesigns.db |
| PORT | Server port | 3000 |
| ENV | Environment (development/production) | development |
| LOG_LEVEL | Logging level (DEBUG/INFO/WARN/ERROR) | DEBUG |
| SITE_NAME | Site name for meta tags | Ron's Designs and Signs |
| BREVO_API_KEY | Brevo API key for emails | (optional) |
| CONTACT_EMAIL | Email for contact form notifications | ronsdesigns@hotmail.com |

## Key Commands

| Command | Description |
|---------|-------------|
| `make dev` | Start with hot reload (main development workflow) |
| `make build` | Build production binary |
| `make test` | Run tests with race detection |
| `make lint` | Run golangci-lint and templ fmt |
| `make css` | Build Tailwind CSS (one-time) |
| `make css-watch` | Watch and rebuild Tailwind CSS |
| `make migrate` | Run database migrations |
| `make setup` | Install development tools |

## Project Structure

```
.
├── cmd/server/           # Application entry point
│   ├── main.go          # Server setup and startup
│   ├── slog.go          # Structured logging configuration
│   └── generate.go      # go:generate directives
├── internal/
│   ├── config/          # Environment configuration
│   ├── ctxkeys/         # Context key types
│   ├── database/        # SQLite database and migrations
│   ├── email/           # Brevo email service
│   ├── handler/         # HTTP handlers
│   ├── meta/            # SEO/meta tag helpers
│   └── middleware/      # Echo middleware
├── templates/
│   ├── layouts/         # Base layout and meta templates
│   └── pages/           # Page templates (home, services, gallery, contact)
├── static/
│   └── css/             # Tailwind CSS (input.css → output.css)
├── images/              # Existing images (preserved)
└── data/                # SQLite database files
```

## Code Patterns

### Logging
Always use `slog` for logging:
```go
slog.Info("message", "key", value)
slog.Error("failed to do thing", "error", err)
```

### Error Handling
Wrap errors with context:
```go
if err != nil {
    return fmt.Errorf("failed to process: %w", err)
}
```

### Templates
Templates own their meta - handlers don't pass it:
```go
// Handler
func (h *Handler) Page(c echo.Context) error {
    return pages.Page().Render(c.Request().Context(), c.Response().Writer)
}
```

```templ
// Template
templ Page() {
    @layouts.Base(meta.New("Title", "Description")) {
        // content
    }
}
```

### Database
Direct SQL via database/sql (no ORM):
```go
_, err := h.db.Conn.ExecContext(ctx, "INSERT INTO ...", args...)
```

## Routes

| Path | Method | Handler | Description |
|------|--------|---------|-------------|
| / | GET | Home | Home page |
| /services | GET | Services | Services page |
| /gallery | GET | Gallery | Gallery page |
| /contact | GET | Contact | Contact page |
| /contact | POST | ContactSubmit | Contact form submission |
| /health | GET | Health | Health check endpoint |
| /static/* | GET | Static | Static files (CSS, JS) |
| /images/* | GET | Static | Image files |

## Database

SQLite database stored at `./data/ronsdesigns.db`

### Tables
- `contact_submissions` - Stores contact form submissions

### Migrations
Migrations are embedded and run automatically on startup.

## Brevo Email Integration

Contact form submissions can optionally send email notifications via Brevo.
Set `BREVO_API_KEY` in `.envrc` to enable.

## Deployment (Self-Hosted)

1. Build the binary:
   ```bash
   make build
   ```

2. Copy to server:
   - `ronsdesigns` (binary)
   - `static/` directory
   - `images/` directory
   - `data/` directory (or create empty)

3. Set environment variables and run:
   ```bash
   export DATABASE_URL="./data/ronsdesigns.db"
   export PORT="3000"
   export ENV="production"
   ./ronsdesigns
   ```

Consider using systemd or similar for process management.
