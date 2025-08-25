package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// NormalizePhoneNumber removes common characters and country codes from a phone number.
func NormalizePhoneNumber(phone string) string {
	// Remove spaces and special characters
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	// Handle different formats
	if strings.HasPrefix(phone, "+62") {
		phone = phone[3:]
	} else if strings.HasPrefix(phone, "62") {
		phone = phone[2:]
	}

	// Remove leading zero
	if strings.HasPrefix(phone, "0") {
		phone = phone[1:]
	}

	return phone
}

// FormatPhoneNumber ensures the phone number is in the correct format for the API.
func FormatPhoneNumber(phone string, withCountryCode bool) string {
	normalized := NormalizePhoneNumber(phone)
	if withCountryCode {
		return "62" + normalized
	}
	return normalized
}

// GenerateTransactionID creates a unique transaction ID.
func GenerateTransactionID() string {
	return fmt.Sprintf("TXN_%d", time.Now().UnixNano())
}

// ParseTimestamp converts a string timestamp to an int64.
func ParseTimestamp(timestampStr string) int64 {
	if timestamp, err := strconv.ParseInt(timestampStr, 10, 64); err == nil {
		return timestamp
	}
	return 0
}
