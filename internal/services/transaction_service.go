package services

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/nabilulilalbab/nadia/internal/database"
	"github.com/nabilulilalbab/nadia/internal/models"
	"github.com/nabilulilalbab/nadia/internal/utils"
)

// TransactionService handles business logic related to transactions and OTP sessions.
type TransactionService struct {
	db               *sql.DB
	otpSessions      map[string]*models.OTPSession
	sessionMutex     sync.RWMutex
	transactions     map[string]*models.TransactionRecord
	transactionMutex sync.RWMutex
}

// NewTransactionService creates a new TransactionService.
func NewTransactionService(db *sql.DB) (*TransactionService, error) {
	ts := &TransactionService{
		db:          db,
		otpSessions: make(map[string]*models.OTPSession),
	}

	// Load existing transactions from the database
	transactions, err := database.LoadTransactionsFromDB(db)
	if err != nil {
		return nil, err
	}
	ts.transactions = transactions

	return ts, nil
}

// OTP Session Management

// CreateOTPSession creates and stores a new OTP session.
func (s *TransactionService) CreateOTPSession(phone, authID string) {
	s.sessionMutex.Lock()
	defer s.sessionMutex.Unlock()

	normalizedPhone := utils.NormalizePhoneNumber(phone)
	s.otpSessions[normalizedPhone] = &models.OTPSession{
		PhoneNumber: normalizedPhone,
		AuthID:      authID,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(5 * time.Minute),
	}
}

// GetOTPSession retrieves an OTP session if it exists and is not expired.
func (s *TransactionService) GetOTPSession(phone string) (*models.OTPSession, error) {
	s.sessionMutex.RLock()
	defer s.sessionMutex.RUnlock()

	normalizedPhone := utils.NormalizePhoneNumber(phone)
	session, exists := s.otpSessions[normalizedPhone]
	if !exists {
		return nil, log.Output(1, "OTP session not found. Please request OTP first.")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, log.Output(1, "OTP session expired. Please request a new OTP.")
	}

	return session, nil
}

// DeleteOTPSession removes an OTP session.
func (s *TransactionService) DeleteOTPSession(phone string) {
	s.sessionMutex.Lock()
	defer s.sessionMutex.Unlock()

	normalizedPhone := utils.NormalizePhoneNumber(phone)
	delete(s.otpSessions, normalizedPhone)
}

// Transaction Management

// RecordTransaction creates a new transaction record and saves it.
func (s *TransactionService) RecordTransaction(phoneNumber, packageCode, packageName, paymentMethod, source string) *models.TransactionRecord {
	s.transactionMutex.Lock()
	defer s.transactionMutex.Unlock()

	record := &models.TransactionRecord{
		ID:            utils.GenerateTransactionID(),
		PhoneNumber:   phoneNumber,
		PackageCode:   packageCode,
		PackageName:   packageName,
		PaymentMethod: paymentMethod,
		Source:        source,
		Status:        "PENDING",
		CreatedAt:     time.Now(),
	}

	s.transactions[record.ID] = record

	if err := database.SaveTransactionToDB(s.db, record); err != nil {
		log.Printf("Failed to save new transaction to database: %v", err)
	}

	return record
}

// UpdateTransactionStatus updates the status of an existing transaction.
func (s *TransactionService) UpdateTransactionStatus(id, status, trxID string, amount, processingFee int, errorMsg string) {
	s.transactionMutex.Lock()
	defer s.transactionMutex.Unlock()

	if record, exists := s.transactions[id]; exists {
		record.Status = status
		record.TrxID = trxID
		record.Amount = amount
		record.ProcessingFee = processingFee
		record.ErrorMessage = errorMsg
		now := time.Now()
		record.CompletedAt = &now

		if err := database.SaveTransactionToDB(s.db, record); err != nil {
			log.Printf("Failed to update transaction in database: %v", err)
		}
	}
}

// UpdateTransactionPackageName updates the package name of a transaction.
func (s *TransactionService) UpdateTransactionPackageName(id, packageName string) {
	s.transactionMutex.Lock()
	defer s.transactionMutex.Unlock()

	if record, exists := s.transactions[id]; exists {
		record.PackageName = packageName
		if err := database.SaveTransactionToDB(s.db, record); err != nil {
			log.Printf("Failed to update transaction package name in database: %v", err)
		}
	}
}

// GetSystemStats calculates and returns system-wide transaction statistics.
func (s *TransactionService) GetSystemStats() models.SystemStats {
	s.transactionMutex.RLock()
	defer s.transactionMutex.RUnlock()

	var total, successful, failed, revenue int

	for _, tx := range s.transactions {
		total++
		if tx.Status == "SUCCESS" {
			successful++
			revenue += tx.Amount
		} else if tx.Status == "FAILED" {
			failed++
		}
	}

	successRate := 0.0
	if total > 0 {
		successRate = float64(successful) / float64(total) * 100
	}

	return models.SystemStats{
		TotalTransactions:      total,
		SuccessfulTransactions: successful,
		FailedTransactions:     failed,
		TotalRevenue:           revenue,
		SuccessRate:            successRate,
		LastUpdated:            time.Now(),
	}
}

// GetSourceStats calculates and returns transaction statistics grouped by source.
func (s *TransactionService) GetSourceStats() []models.SourceStats {
	s.transactionMutex.RLock()
	defer s.transactionMutex.RUnlock()

	sourceMap := make(map[string]*models.SourceStats)

	for _, tx := range s.transactions {
		source := tx.Source
		if source == "" {
			source = "unknown"
		}

		if _, exists := sourceMap[source]; !exists {
			sourceMap[source] = &models.SourceStats{Source: source}
		}

		sourceMap[source].Count++
		if tx.Status == "SUCCESS" {
			sourceMap[source].Revenue += tx.Amount
		}
	}

	var result []models.SourceStats
	for _, stats := range sourceMap {
		if stats.Count > 0 {
			successCount := 0
			for _, tx := range s.transactions {
				txSource := tx.Source
				if txSource == "" {
					txSource = "unknown"
				}
				if txSource == stats.Source && tx.Status == "SUCCESS" {
					successCount++
				}
			}
			stats.SuccessRate = float64(successCount) / float64(stats.Count) * 100
		}
		result = append(result, *stats)
	}

	return result
}

// GetRecentTransactions returns the most recent transactions.
func (s *TransactionService) GetRecentTransactions(limit int) []models.TransactionRecord {
	txList, err := database.GetAllTransactionsFromDB(s.db, limit)
	if err != nil {
		log.Printf("Failed to get transactions from DB for recent list: %v", err)
		// Fallback to in-memory data if DB fails
		s.transactionMutex.RLock()
		defer s.transactionMutex.RUnlock()

		var result []models.TransactionRecord
		for _, tx := range s.transactions {
			result = append(result, *tx)
		}

		// Sort by created time (newest first)
		for i := 0; i < len(result)-1; i++ {
			for j := i + 1; j < len(result); j++ {
				if result[i].CreatedAt.Before(result[j].CreatedAt) {
					result[i], result[j] = result[j], result[i]
			}
			}
		}

		if len(result) > limit {
			result = result[:limit]
		}
		return result
	}

	var result []models.TransactionRecord
	for _, tx := range txList {
		result = append(result, *tx)
	}
	return result
}

// GetTransactionDetail retrieves a single transaction by its ID.
func (s *TransactionService) GetTransactionDetail(id string) (*models.TransactionRecord, bool) {
	s.transactionMutex.RLock()
	defer s.transactionMutex.RUnlock()
	tx, exists := s.transactions[id]
	return tx, exists
}
