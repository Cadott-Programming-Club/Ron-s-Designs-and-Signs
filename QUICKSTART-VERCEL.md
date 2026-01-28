# Quick Deploy to Vercel

## 🚀 Deploy Now (5 minutes)

### Step 1: Install Vercel CLI
```bash
npm install -g vercel
```

### Step 2: Login
```bash
vercel login
```

### Step 3: Deploy
```bash
cd /path/to/Ron-s-Designs-and-Signs
vercel
```

## ⚡ Required Environment Variables

Set in Vercel Dashboard → Settings → Environment Variables:

```
DATABASE_URL=<your-turso-or-postgres-url>
ENV=production
SITE_NAME=Ron's Designs and Signs
SITE_URL=<your-vercel-url>
CONTACT_EMAIL=ronsdesigns@hotmail.com
```

## 🗄️ Database Options

### Turso (Recommended - SQLite Compatible)
```bash
# Install Turso CLI
curl -sSfL https://get.tur.so/install.sh | bash

# Create database
turso auth login
turso db create ronsdesigns
turso db show ronsdesigns --url
turso db tokens create ronsdesigns

# Set DATABASE_URL in Vercel:
# libsql://[url]?authToken=[token]
```

### Vercel Postgres
1. Add from Vercel Dashboard
2. Copy `DATABASE_URL`
3. Update code to use pgx driver

## 📝 Full Documentation

See [DEPLOYMENT.md](./DEPLOYMENT.md) for:
- GitHub Actions setup
- Troubleshooting
- Custom domains
- Advanced configuration

## ✅ Pre-deployment Checklist

- [ ] Vercel account created
- [ ] Database created and URL obtained
- [ ] Environment variables set in Vercel
- [ ] Code pushed to GitHub (for CI/CD)
- [ ] `vercel login` completed
- [ ] Ready to run `vercel`

## 🆘 Quick Help

```bash
# View logs
vercel logs

# List deployments
vercel ls

# Rollback if needed
vercel rollback <deployment-url>
```

---

**Need help?** See [DEPLOYMENT.md](./DEPLOYMENT.md) for detailed instructions.
