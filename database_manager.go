package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DatabaseManager struct {
	db *sql.DB
	km *KeyManager
}

func NewDatabaseManager(km *KeyManager) (*DatabaseManager, error) {
	db, err := sql.Open("sqlite3", "./keys.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &DatabaseManager{db: db, km: km}, nil
}

func (dm *DatabaseManager) Close() {
	dm.db.Close()
}

func (dm *DatabaseManager) InitializeDB() error {
	createTableQuery := `CREATE TABLE IF NOT EXISTS keys (name TEXT PRIMARY KEY, private_key TEXT, public_key TEXT);`

	_, err := dm.db.Exec(createTableQuery)
	if err != nil {
		return err
	}
	return nil
}

func (dm *DatabaseManager) InsertKey(keyName string) error {
	// Generate new RSA keys
	privateKey, publicKey, err := dm.km.GenerateRSAKeys()
	if err != nil {
		return err
	}

	// Convert keys to PEM string
	privateKeyString, publicKeyString, err := dm.km.ExportRSAKeysToString(privateKey, publicKey)
	if err != nil {
		return err
	}

	insertKeyQuery := `INSERT INTO keys (name, private_key, public_key) VALUES (?, ?, ?);`

	_, err = dm.db.Exec(insertKeyQuery, keyName, privateKeyString, publicKeyString)
	if err != nil {
		return err
	}

	return nil
}

func (dm *DatabaseManager) GetKey(keyName string) (string, string, error) {
	getKeyQuery := `SELECT private_key, public_key FROM keys WHERE name = ?;`

	row := dm.db.QueryRow(getKeyQuery, keyName)

	var privateKeyString string
	var publicKeyString string
	err := row.Scan(&privateKeyString, &publicKeyString)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", "", nil
		} else {
			return "", "", err
		}
	}

	return privateKeyString, publicKeyString, nil
}
