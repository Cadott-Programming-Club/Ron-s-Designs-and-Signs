# Deploying to Vercel

This guide provides step-by-step instructions for deploying Ron's Designs and Signs to Vercel.

## Prerequisites

1. A Vercel account (sign up at https://vercel.com)
2. Vercel CLI installed: `npm install -g vercel`
3. A cloud database (Turso, Vercel Postgres, or Neon)

## Option 1: Manual Deployment via CLI

### Step 1: Prepare Your Project

```bash
# Install Node dependencies
npm install

# Generate Templ templates
go run github.com/a-h/templ/cmd/templ@latest generate

# Build Tailwind CSS
npm run css

# Copy CSS to public directory
cp static/css/output.css public/static/css/output.css
```

### Step 2: Login to Vercel

```bash
vercel login
```

This will open your browser to authenticate.

### Step 3: Link Your Project

```bash
vercel link
```

Answer the prompts:
- Set up and deploy? **Yes**
- Which scope? Select your account/team
- Link to existing project? **No** (or Yes if you've already created one)
- What's your project's name? `ronsdesigns` (or your preferred name)
- In which directory is your code located? `.` (current directory)

### Step 4: Configure Environment Variables

Set required environment variables in Vercel:

```bash
vercel env add DATABASE_URL production
vercel env add ENV production
vercel env add SITE_NAME production
vercel env add SITE_URL production
vercel env add CONTACT_EMAIL production
```

Enter values when prompted:
- `DATABASE_URL`: Your Turso/Postgres connection string
- `ENV`: `production`
- `SITE_NAME`: `Ron's Designs and Signs`
- `SITE_URL`: Your Vercel domain (e.g., `https://ronsdesigns.vercel.app`)
- `CONTACT_EMAIL`: `ronsdesigns@hotmail.com`

Optional:
```bash
vercel env add BREVO_API_KEY production
vercel env add DEFAULT_OG_IMAGE production
```

### Step 5: Deploy

**Preview deployment:**
```bash
vercel
```

**Production deployment:**
```bash
vercel --prod
```

### Step 6: View Your Deployment

```bash
vercel open
```

Or visit the URL shown in the deployment output.

## Option 2: Automated Deployment via GitHub Actions

### Step 1: Get Vercel Credentials

1. Go to https://vercel.com/account/tokens
2. Create a new token and copy it

3. Get your Org ID and Project ID:
```bash
# After linking your project locally
cat .vercel/project.json
```

### Step 2: Add GitHub Secrets

Go to your GitHub repository → Settings → Secrets and variables → Actions

Add these secrets:
- `VERCEL_TOKEN`: Your Vercel token
- `VERCEL_ORG_ID`: Your organization ID from `.vercel/project.json`
- `VERCEL_PROJECT_ID`: Your project ID from `.vercel/project.json`

### Step 3: Configure Environment Variables in Vercel

1. Go to your project in Vercel Dashboard
2. Navigate to Settings → Environment Variables
3. Add all required variables:
   - `DATABASE_URL`
   - `ENV` = `production`
   - `SITE_NAME` = `Ron's Designs and Signs`
   - `SITE_URL` = Your Vercel domain
   - `CONTACT_EMAIL` = `ronsdesigns@hotmail.com`
   - `BREVO_API_KEY` (optional)
   - `DEFAULT_OG_IMAGE` (optional)

### Step 4: Push to GitHub

The workflow file `.github/workflows/vercel-deploy.yml` is already set up.

Push to `main` branch to trigger deployment:
```bash
git push origin main
```

The GitHub Action will:
1. Generate Templ templates
2. Build Tailwind CSS
3. Deploy to Vercel

## Database Setup

### Recommended: Turso (SQLite-compatible)

Turso provides edge-distributed SQLite databases that work great with serverless:

1. Sign up at https://turso.tech
2. Install Turso CLI:
```bash
curl -sSfL https://get.tur.so/install.sh | bash
```

3. Login and create a database:
```bash
turso auth login
turso db create ronsdesigns
```

4. Get the database URL:
```bash
turso db show ronsdesigns --url
```

5. Create an auth token:
```bash
turso db tokens create ronsdesigns
```

6. Combine URL and token:
```
libsql://[database-url]?authToken=[your-token]
```

Set this as your `DATABASE_URL` in Vercel.

**Note:** You'll need to update your code to use `libsql` driver instead of SQLite:
```go
import _ "github.com/tursodatabase/libsql-client-go/libsql"
```

### Alternative: Vercel Postgres or Neon

1. Add Postgres from Vercel Dashboard or sign up at https://neon.tech
2. Get your connection string
3. Update your code to use `pgx` driver
4. Set `DATABASE_URL` in Vercel

## Troubleshooting

### Build Fails

Check Vercel build logs:
```bash
vercel logs
```

Common issues:
- Missing `go.mod` - ensure it's at project root
- Failed to generate templates - run `templ generate` before deploying
- Missing environment variables - check Vercel Dashboard → Settings → Environment Variables

### Static Files Not Loading

Verify in `vercel.json`:
- Static routes come before catch-all route
- Files exist in `public/` directory

### Database Connection Errors

- Verify `DATABASE_URL` is set in Vercel
- Check database allows connections from Vercel IPs
- For Turso, ensure auth token is included in URL
- Test connection locally first

### Function Timeout

- Free tier: 10s timeout
- Pro tier: 60s timeout  
- Optimize slow database queries
- Add database indexes
- Use caching where appropriate

## Viewing Logs

```bash
# Recent logs
vercel logs

# Follow logs in real-time
vercel logs --follow

# Logs for specific deployment
vercel logs [deployment-url]
```

## Rollback

If a deployment has issues:
```bash
# List deployments
vercel ls

# Rollback to previous
vercel rollback [deployment-url]
```

## Custom Domain

1. Go to Vercel Dashboard → Your Project → Settings → Domains
2. Add your custom domain
3. Update DNS records as instructed
4. Update `SITE_URL` environment variable

## Local Testing

Test locally before deploying:

```bash
# Set environment variables
export DATABASE_URL="./data/ronsdesigns.db"
export ENV="development"

# Run locally
make dev
```

## Next Steps

After successful deployment:

1. ✅ Test all pages work
2. ✅ Test contact form submission
3. ✅ Verify images load correctly
4. ✅ Check console for errors
5. ✅ Test on mobile devices
6. ✅ Configure custom domain (optional)
7. ✅ Set up monitoring/alerts
8. ✅ Add your domain to BREVO if using email

## Resources

- [Vercel Documentation](https://vercel.com/docs)
- [Vercel Go Runtime](https://vercel.com/docs/runtimes#official-runtimes/go)
- [Vercel CLI Reference](https://vercel.com/docs/cli)
- [Turso Documentation](https://docs.turso.tech)
- [Original Deployment Guide](https://gist.githubusercontent.com/corylanou/6f5a79948f5ddb1da656c18551b2edb8/raw/ff7142afe1854bd09ae000fa91be356138f2fa6b/vercel-golang-delpoy.md)
