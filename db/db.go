package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const dbFile = "store.db"
const tableSchema = `CREATE TABLE IF NOT EXISTS kv_store (key TEXT PRIMARY KEY, value TEXT);`

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbFile)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(tableSchema)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func SetValue(db *sql.DB, key, value string) error {
	_, err := db.Exec(`INSERT INTO kv_store (key, value) VALUES (?, ?) ON CONFLICT(key) DO UPDATE SET value=excluded.value;`, key, value)
	return err
}

func GetValue(db *sql.DB, key string) (string, error) {
	var value string
	err := db.QueryRow(`SELECT value FROM kv_store WHERE key = ?;`, key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("key not found")
	}
	return value, err
}

func DeleteKey(db *sql.DB, key string) error {
	_, err := db.Exec(`DELETE FROM kv_store WHERE key = ?;`, key)
	return err
}

func ListKeys(db *sql.DB) (map[string]string, error) {
	rows, err := db.Query(`SELECT key, value FROM kv_store;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make(map[string]string)
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return nil, err
		}
		result[key] = value
	}
	return result, nil
}
