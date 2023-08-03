package key_manager

import (
	"encoding/gob"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"time"
)

type CopyData struct {
	Timestamp time.Time
	Path      string
	Username  string
}

type KeyData struct {
	PrivateKey string
	PublicKey  string
	Copies     []CopyData
}

type DatabaseManager struct {
	dbPath string
	km     *KeyManager
	data   map[string]KeyData
}

func getDataDir() string {
	var dataDir string

	switch runtime.GOOS {
	case "windows":
		dataDir = os.Getenv("APPDATA")
	case "darwin":
		dataDir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support")
	default: // Unix, Linux, etc.
		dataDir = os.Getenv("HOME")
	}

	dataDir = filepath.Join(dataDir, "KeyManager")

	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return dataDir
}

func NewDatabaseManager(km *KeyManager) (*DatabaseManager, error) {
	dbPath := filepath.Join(getDataDir(), "keys.db")

	dm := &DatabaseManager{
		dbPath: dbPath,
		km:     km,
		data:   make(map[string]KeyData),
	}

	err := dm.loadDB()
	if err != nil {
		return nil, err
	}

	return dm, nil
}

func (dm *DatabaseManager) loadDB() error {
	file, err := os.Open(dm.dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // Ignore if file doesn't exist
		}
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&dm.data)
	if err != nil {
		return err
	}

	return nil
}

func (dm *DatabaseManager) saveDB() error {
	file, err := os.Create(dm.dbPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(dm.data)
	if err != nil {
		return err
	}

	return nil
}

func (dm *DatabaseManager) InsertKey(keyName string) error {
	_, ok := dm.data[keyName]
	if ok {
		return fmt.Errorf("key already exists: %s", keyName)
	}

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

	data := KeyData{
		PrivateKey: privateKeyString,
		PublicKey:  publicKeyString,
		Copies:     []CopyData{},
	}

	dm.data[keyName] = data

	return dm.saveDB()
}

func (dm *DatabaseManager) GetKey(keyName string) (string, string, error) {
	data, ok := dm.data[keyName]

	if !ok {
		return "", "", fmt.Errorf("key not found: %s", keyName)
	}

	return data.PrivateKey, data.PublicKey, nil
}

func (dm *DatabaseManager) AddCopyRecord(keyName string, outputPath string) error {
	// Get the current user
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return err
	}

	data, ok := dm.data[keyName]
	if !ok {
		return fmt.Errorf("key not found: %s", keyName)
	}

	data.Copies = append(data.Copies, CopyData{
		Timestamp: time.Now(),
		Path:      outputPath,
		Username:  usr.Username,
	})

	dm.data[keyName] = data

	return dm.saveDB()
}

func (dm *DatabaseManager) GetCopyData(keyName string) ([]CopyData, error) {
	data, ok := dm.data[keyName]
	if !ok {
		return nil, fmt.Errorf("key not found: %s", keyName)
	}

	return data.Copies, nil
}

func (dm *DatabaseManager) DeleteKey(keyName string) error {
	_, ok := dm.data[keyName]
	if !ok {
		return fmt.Errorf("key not found: %s", keyName)
	}

	delete(dm.data, keyName)

	return dm.saveDB()
}
