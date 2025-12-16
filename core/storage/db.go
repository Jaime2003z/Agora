package storage

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func NewDB(filename string) (*DB, error) {
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	createTables := `
	-- Tabla para bloques de la cadena (ledger)
	CREATE TABLE IF NOT EXISTS blocks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		hash TEXT NOT NULL,
		prev_hash TEXT,
		data TEXT,  -- Puedes serializar JSON o lo que uses
		timestamp INTEGER
	);

	-- Tabla para propuestas de gobernanza
	CREATE TABLE IF NOT EXISTS proposals (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT,
		creator_id TEXT NOT NULL,
		created_at INTEGER NOT NULL,
		closed BOOLEAN DEFAULT 0,
		votes_yes INTEGER DEFAULT 0,
		votes_no INTEGER DEFAULT 0,
		executed INTEGER DEFAULT 0,  -- 0 = no, 1 = sí
		result_approved BOOLEAN,
		result_yes_votes INTEGER,
		result_no_votes INTEGER,
		result_total_votes INTEGER,
		result_decided_at INTEGER
	);

	-- Tabla para votos (opcional, para rastrear quién votó)
	CREATE TABLE IF NOT EXISTS votes (
		id TEXT PRIMARY KEY,
		proposal_id TEXT NOT NULL,
		voter_id TEXT NOT NULL,
		approved BOOLEAN NOT NULL,
		timestamp INTEGER NOT NULL,
		FOREIGN KEY (proposal_id) REFERENCES proposals(id),
		UNIQUE(proposal_id, voter_id)
	);
	`

	if _, err = db.Exec(createTables); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}
