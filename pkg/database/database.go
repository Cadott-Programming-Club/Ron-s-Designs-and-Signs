// This is a Go file for database operations
package database
}
package database

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	Conn *sql.DB
}

func New(ctx context.Context, databasePath string) (*DB, error) {
	dir := filepath.Dir(databasePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		// On some hosting providers (like Vercel) the filesystem is read-only.
		// If that's the case, fall back to a writable temporary path so the
		// app can still start (data will be ephemeral).
		if os.Getenv("VERCEL") != "" || strings.Contains(err.Error(), "read-only file system") {
			tmpPath := filepath.Join(os.TempDir(), filepath.Base(databasePath))
			slog.Warn("filesystem read-only, falling back to temp DB path", "original", databasePath, "fallback", tmpPath)
			databasePath = tmpPath
			dir = filepath.Dir(databasePath)
			// ensure tmp dir exists (usually does)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return nil, fmt.Errorf("unable to create fallback database directory: %w", err)
			}
		} else {
			return nil, fmt.Errorf("unable to create database directory: %w", err)
		}
	}

	conn, err := sql.Open("sqlite", databasePath+"?_foreign_keys=on&_journal_mode=WAL")
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	db := &DB{
		Conn: conn,
	}

	if err := db.migrate(); err != nil {
		return nil, fmt.Errorf("unable to run migrations: %w", err)
	}

	return db, nil
}

func (db *DB) Close() {
	db.Conn.Close()
}

package database

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

type DB struct {
	Conn *sql.DB
}

func New(ctx context.Context, databasePath string) (*DB, error) {
	dir := filepath.Dir(databasePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		// On some hosting providers (like Vercel) the filesystem is read-only.
		// If that's the case, fall back to a writable temporary path so the
		// app can still start (data will be ephemeral).
		if os.Getenv("VERCEL") != "" || strings.Contains(err.Error(), "read-only file system") {
			tmpPath := filepath.Join(os.TempDir(), filepath.Base(databasePath))
			slog.Warn("filesystem read-only, falling back to temp DB path", "original", databasePath, "fallback", tmpPath)
			databasePath = tmpPath
			dir = filepath.Dir(databasePath)
			// ensure tmp dir exists (usually does)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return nil, fmt.Errorf("unable to create fallback database directory: %w", err)
			}
		} else {
			return nil, fmt.Errorf("unable to create database directory: %w", err)
		}
	}

	conn, err := sql.Open("sqlite", databasePath+"?_foreign_keys=on&_journal_mode=WAL")
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	db := &DB{
		Conn: conn,
	}

	if err := db.migrate(); err != nil {
		return nil, fmt.Errorf("unable to run migrations: %w", err)
	}

	return db, nil
}

func (db *DB) Close() {
	db.Conn.Close()
}

func (db *DB) migrate() error {
	goose.SetBaseFS(migrationsFS)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	if err := goose.Up(db.Conn, "migrations"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
