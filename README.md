# Ron's Designs and Signs

Custom signs, decals, and apparel business website built with Go + Templ + HTMX + Tailwind CSS.

## Prerequisites

- [Go 1.21+](https://go.dev/doc/install)
- [Node.js 18+](https://nodejs.org/) (for Tailwind CSS)
- [Make](https://www.gnu.org/software/make/)
- [direnv](https://direnv.net/) (optional, but recommended)

### Install Development Tools

```bash
make setup
```

This installs:
- `templ` - Template generation
- `air` - Hot reload
- `golangci-lint` - Linting
- Tailwind CSS CLI

## Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs.git
cd Ron-s-Designs-and-Signs
```

### 2. Configure Environment Variables

Copy the environment template:

```bash
cp .envrc.example .envrc  # If exists, otherwise create .envrc
```

Or create `.envrc` with the following content:

```bash
# Database (SQLite)
export DATABASE_URL="./data/ronsdesigns.db"

# Server
export PORT="8000"
export ENV="development"
export LOG_LEVEL="DEBUG"

# Site / SEO
export SITE_NAME="Ron's Designs and Signs"
export SITE_URL="http://localhost:8000"
export DEFAULT_OG_IMAGE="/static/images/og-default.png"

# Brevo (email service) - Optional
export BREVO_API_KEY=""
export CONTACT_EMAIL="ronsdesigns@hotmail.com"
```

Load environment variables:

```bash
direnv allow  # If using direnv
# OR
source .envrc  # Manual loading
```

### 3. Initialize Database

The database will be created automatically on first run, but you can initialize it manually:

```bash
make migrate
```

### 4. Build CSS

Generate Tailwind CSS:

```bash
make css
```

### 5. Run Development Server

Start the server with hot reload:

```bash
make dev
```

The site will be available at [http://localhost:8000](http://localhost:8000)

### 6. Watch CSS Changes (Optional)

In a separate terminal, watch for Tailwind CSS changes:

```bash
make css-watch
```

## Development Workflow

### Primary Development Command

```bash
make dev
```

This command:
- Kills any existing process on port 8000
- Regenerates Templ templates
- Runs `go mod tidy`
- Rebuilds and restarts the server automatically on file changes
- **You don't need to manually run:** `templ generate`, `go build`, or `air`

### Available Make Commands

| Command | Description |
|---------|-------------|
| `make dev` | Start with hot reload (main development workflow) |
| `make build` | Build production binary |
| `make run` | Run the server without hot reload |
| `make test` | Run tests with race detection |
| `make lint` | Run golangci-lint and templ fmt |
| `make css` | Build Tailwind CSS (one-time) |
| `make css-watch` | Watch and rebuild Tailwind CSS |
| `make migrate` | Run database migrations |
| `make setup` | Install development tools |
| `make clean` | Remove build artifacts |

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
│   ├── css/             # Tailwind CSS (input.css → output.css)
│   └── fonts/           # Web fonts
├── images/              # Image assets
│   └── work/            # Gallery work samples
├── data/                # SQLite database files
└── tmp/                 # Build artifacts and logs
```

## Routes

| Path | Method | Description |
|------|--------|-------------|
| `/` | GET | Home page |
| `/services` | GET | Services page |
| `/gallery` | GET | Portfolio gallery |
| `/contact` | GET | Contact form |
| `/contact` | POST | Contact form submission |
| `/health` | GET | Health check endpoint |
| `/static/*` | GET | Static files (CSS, JS) |
| `/images/*` | GET | Image files |

## Technology Stack

- **Backend:** Go 1.21+ with Echo framework
- **Templates:** Templ (type-safe HTML templates)
- **Styling:** Tailwind CSS v4
- **Interactivity:** HTMX
- **Database:** SQLite
- **Email:** Brevo API (optional)
- **Hot Reload:** Air

## Configuration

All configuration is done via environment variables in `.envrc`:

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | SQLite database path | `./data/ronsdesigns.db` |
| `PORT` | Server port | `8000` |
| `ENV` | Environment (development/production) | `development` |
| `LOG_LEVEL` | Logging level (DEBUG/INFO/WARN/ERROR) | `DEBUG` |
| `SITE_NAME` | Site name for meta tags | `Ron's Designs and Signs` |
| `SITE_URL` | Base URL for the site | `http://localhost:8000` |
| `BREVO_API_KEY` | Brevo API key for emails | (optional) |
| `CONTACT_EMAIL` | Email for contact form notifications | `ronsdesigns@hotmail.com` |

## Database

SQLite database stored at `./data/ronsdesigns.db`

### Tables
- `contact_submissions` - Stores contact form submissions

### Migrations
Migrations are embedded and run automatically on startup.

## Troubleshooting

### Build Errors

Check build logs:
```bash
cat ./tmp/air-combined.log
```

### Port Already in Use

Kill existing processes on port 8000:
```bash
lsof -ti:8000 | xargs kill -9
```

### Template Generation Errors

Regenerate templates manually:
```bash
templ generate
```

### CSS Not Updating

Rebuild CSS:
```bash
make css
```

Or use watch mode:
```bash
make css-watch
```

## Production Deployment

### 1. Build Production Binary

```bash
make build
```

### 2. Copy Files to Server

- `ronsdesigns` (binary)
- `static/` directory
- `images/` directory
- `data/` directory (or create empty)

### 3. Set Production Environment

```bash
export DATABASE_URL="./data/ronsdesigns.db"
export PORT="8000"
export ENV="production"
export LOG_LEVEL="INFO"
export SITE_URL="https://yourdomain.com"
```

### 4. Run the Binary

```bash
./ronsdesigns
```

Consider using systemd or similar for process management.

## Contributing

See [CLAUDE.md](./CLAUDE.md) for detailed development patterns and code conventions.

## License

Copyright © 2025 Ron's Designs and Signs