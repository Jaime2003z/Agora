package storage

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

func NewDB(filename string) *DB {
	db, err := sql.Open("sqlite", filename)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
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
		votes_yes INTEGER DEFAULT 0,
		votes_no INTEGER DEFAULT 0,
		executed INTEGER DEFAULT 0  -- 0 = no, 1 = sí
	);

	-- Tabla para votos (opcional, para rastrear quién votó)
	CREATE TABLE IF NOT EXISTS votes (
		proposal_id TEXT,
		user_id TEXT,
		vote BOOLEAN,  -- 1 = sí, 0 = no
		PRIMARY KEY (proposal_id, user_id)
	);
	`

	if _, err = db.Exec(createTables); err != nil {
		panic(err)
	}

	return &DB{db}
}	

func (db *DB) Close() error {
	return db.DB.Close()
}