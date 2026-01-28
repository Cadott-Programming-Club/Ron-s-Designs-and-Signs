# Vercel Deployment Guide

This document provides instructions for deploying Ron's Designs and Signs to Vercel.

## Prerequisites

1. Install Vercel CLI:
```bash
npm i -g vercel
```

2. Login to Vercel:
```bash
vercel login
```

## Environment Variables

Configure these in the Vercel Dashboard (Settings → Environment Variables):

| Variable | Description | Example |
|----------|-------------|---------|
| `DATABASE_URL` | Database connection URL (Postgres/Turso) | `postgres://user:pass@host/db` |
| `ENV` | Environment (production) | `production` |
| `PORT` | Server port (optional, Vercel provides) | `3000` |
| `SITE_NAME` | Site name for SEO | `Ron's Designs and Signs` |
| `SITE_URL` | Full site URL | `https://your-domain.vercel.app` |
| `DEFAULT_OG_IMAGE` | Default OG image path | `/static/images/og-default.png` |
| `BREVO_API_KEY` | Brevo API key for emails (optional) | `xkeysib-...` |
| `CONTACT_EMAIL` | Email for contact form | `ronsdesigns@hotmail.com` |

## Database Setup

**Important:** Vercel serverless functions do not support SQLite with CGO. You need to use a cloud database:

### Option 1: Turso (SQLite-compatible, edge-distributed)
1. Sign up at https://turso.tech
2. Create a database
3. Set `DATABASE_URL` to your Turso connection string

### Option 2: Vercel Postgres or Neon
1. Add Postgres from Vercel Dashboard
2. Update your code to use `pgx` driver instead of SQLite
3. Run migrations on the new database

## Deployment

### First-time Setup

1. Install dependencies and build locally to ensure everything works:
```bash
# Generate Templ templates
go run github.com/a-h/templ/cmd/templ@latest generate

# Build Tailwind CSS
npm run build:css

# Test the build
go build -o /tmp/test ./api/index.go
```

2. Link to Vercel:
```bash
vercel link
```

### Preview Deployment
```bash
vercel
```

### Production Deployment
```bash
vercel --prod
```

### View Logs
```bash
vercel logs
```

## Project Structure

```
.
├── api/
│   └── index.go          # Vercel serverless handler
├── public/                # Static assets served by Vercel
│   ├── static/           # CSS, JS files
│   └── images/           # Image files
├── internal/             # Application code
├── templates/            # Templ templates
├── vercel.json           # Vercel configuration
└── go.mod
```

## How It Works

1. **api/index.go**: Single serverless function that wraps the entire Echo application
2. **vercel.json**: Routes configuration
   - Static files (`/static/*`, `/images/*`) served directly from `public/`
   - All other requests routed to the Go handler
   - **Important:** Route order matters! Static routes must be defined before the catch-all `/(.*)` route
3. **sync.Once**: Ensures the app initializes only once per container (fast warm starts)

## Build Configuration

Vercel automatically detects the Go version from `go.mod` (currently 1.24). 

**Important Notes:**
- Templ templates are generated and committed to the repository (included in `*_templ.go` files)
- Tailwind CSS output should be pre-built and committed, or built using a Vercel build command
- The `@vercel/go` builder compiles the Go code during deployment

### Optional: Build Command for Tailwind

If you want to build Tailwind CSS during deployment, add this to `vercel.json`:

```json
{
  "buildCommand": "npm install && npm run build:css"
}
```

## Local Development

Local development continues to work as before:

```bash
# Start development server with hot reload
make dev

# Watch Tailwind CSS (in separate terminal)
make css-watch
```

The same codebase works both:
- Locally: Full HTTP server with SQLite
- Vercel: Serverless function with cloud database

## Troubleshooting

### Build fails
- Ensure `go.mod` is at project root
- Run `go mod tidy` before deploying
- Check Vercel build logs for specific errors

### Static files not loading
- Verify `vercel.json` routes are correct
- Check that files exist in `public/` directory
- Static routes must come before catch-all route

### Database connection fails
- Verify `DATABASE_URL` is set in Vercel dashboard
- Ensure you're using a cloud database (not local SQLite)
- Check database allows connections from Vercel IPs

### Function timeout
- Free tier: 10s timeout
- Pro tier: 60s timeout
- Optimize slow queries and database operations

## Additional Resources

- [Vercel Go Runtime Documentation](https://vercel.com/docs/runtimes#official-runtimes/go)
- [Vercel Configuration Reference](https://vercel.com/docs/project-configuration)
- [Original Deployment Guide](https://gist.githubusercontent.com/corylanou/6f5a79948f5ddb1da656c18551b2edb8/raw/ff7142afe1854bd09ae000fa91be356138f2fa6b/vercel-golang-delpoy.md)
