package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/nabilulilalbab/nadia/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

// InitDatabase initializes the SQLite database and creates the necessary tables.
func InitDatabase(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Create transactions table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS transactions (
		id TEXT PRIMARY KEY,
		phone_number TEXT NOT NULL,
		package_code TEXT NOT NULL,
		package_name TEXT,
		payment_method TEXT NOT NULL,
		source TEXT,
		status TEXT NOT NULL,
		amount INTEGER DEFAULT 0,
		processing_fee INTEGER DEFAULT 0,
		trx_id TEXT,
		created_at DATETIME NOT NULL,
		completed_at DATETIME,
		error_message TEXT
	);
	
	CREATE INDEX IF NOT EXISTS idx_phone_number ON transactions(phone_number);
	CREATE INDEX IF NOT EXISTS idx_status ON transactions(status);
	CREATE INDEX IF NOT EXISTS idx_source ON transactions(source);
	CREATE INDEX IF NOT EXISTS idx_created_at ON transactions(created_at);
	`

	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	log.Println("Database initialized successfully")
	return db, nil
}

// SaveTransactionToDB saves or updates a transaction record in the database.
func SaveTransactionToDB(db *sql.DB, tx *models.TransactionRecord) error {
	query := `
	INSERT OR REPLACE INTO transactions 
	(id, phone_number, package_code, package_name, payment_method, source, status, 
	 amount, processing_fee, trx_id, created_at, completed_at, error_message)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	var completedAt *time.Time
	if tx.CompletedAt != nil {
		completedAt = tx.CompletedAt
	}

	_, err := db.Exec(query,
		tx.ID, tx.PhoneNumber, tx.PackageCode, tx.PackageName, tx.PaymentMethod,
		tx.Source, tx.Status, tx.Amount, tx.ProcessingFee, tx.TrxID,
		tx.CreatedAt, completedAt, tx.ErrorMessage)

	if err != nil {
		log.Printf("Failed to save transaction to DB: %v", err)
		return err
	}

	return nil
}

// LoadTransactionsFromDB loads all transactions from the database into a map.
func LoadTransactionsFromDB(db *sql.DB) (map[string]*models.TransactionRecord, error) {
	query := `
	SELECT id, phone_number, package_code, package_name, payment_method, source, 
	       status, amount, processing_fee, trx_id, created_at, completed_at, error_message
	FROM transactions
	ORDER BY created_at DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %v", err)
	}
	defer rows.Close()

	transactions := make(map[string]*models.TransactionRecord)
	count := 0
	for rows.Next() {
		tx := &models.TransactionRecord{}
		var completedAt sql.NullTime

		err := rows.Scan(
			&tx.ID, &tx.PhoneNumber, &tx.PackageCode, &tx.PackageName,
			&tx.PaymentMethod, &tx.Source, &tx.Status, &tx.Amount,
			&tx.ProcessingFee, &tx.TrxID, &tx.CreatedAt, &completedAt,
			&tx.ErrorMessage)

		if err != nil {
			log.Printf("Failed to scan transaction: %v", err)
			continue
		}

		if completedAt.Valid {
			tx.CompletedAt = &completedAt.Time
		}

		transactions[tx.ID] = tx
		count++
	}

	log.Printf("Loaded %d transactions from database", count)
	return transactions, nil
}

// GetAllTransactionsFromDB retrieves all transactions from the database as a slice.
func GetAllTransactionsFromDB(db *sql.DB, limit int) ([]*models.TransactionRecord, error) {
	query := `
	SELECT id, phone_number, package_code, package_name, payment_method, source, 
	       status, amount, processing_fee, trx_id, created_at, completed_at, error_message
	FROM transactions
	ORDER BY created_at DESC
	`

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %v", err)
	}
	defer rows.Close()

	var result []*models.TransactionRecord
	for rows.Next() {
		tx := &models.TransactionRecord{}
		var completedAt sql.NullTime

		err := rows.Scan(
			&tx.ID, &tx.PhoneNumber, &tx.PackageCode, &tx.PackageName,
			&tx.PaymentMethod, &tx.Source, &tx.Status, &tx.Amount,
			&tx.ProcessingFee, &tx.TrxID, &tx.CreatedAt, &completedAt,
			&tx.ErrorMessage)

		if err != nil {
			log.Printf("Failed to scan transaction: %v", err)
			continue
		}

		if completedAt.Valid {
			tx.CompletedAt = &completedAt.Time
		}

		result = append(result, tx)
	}

	return result, nil
}

// GetDatabaseStatus checks the current status of the database connection.
func GetDatabaseStatus(db *sql.DB) string {
	if db == nil {
		return "DISCONNECTED"
	}

	if err := db.Ping(); err != nil {
		return "ERROR"
	}

	return "CONNECTED"
}
