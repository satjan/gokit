package gokit

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"strings"
)

// Contains checks if a slice contains a given item.
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// StrToInt converts a string to an integer. Returns 0 on failure.
func StrToInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

// StrToFloat converts a string to a float64. Returns 0.0 on failure.
func StrToFloat(str string) float64 {
	val, _ := strconv.ParseFloat(str, 64)
	return val
}

// IntToStr converts an integer to a string.
func IntToStr(val int) string {
	return strconv.Itoa(val)
}

// RandSeq generates a random hex sequence of n bytes.
func RandSeq(n int) string {
	key := make([]byte, n)
	if _, err := rand.Read(key); err != nil {
		return ""
	}
	return hex.EncodeToString(key)
}

// Lower trims whitespace and converts a string to lowercase.
func Lower(username string) string {
	return strings.ToLower(strings.TrimSpace(username))
}

// Upper trims whitespace and converts a string to uppercase.
func Upper(username string) string {
	return strings.ToUpper(strings.TrimSpace(username))
}

// Capitalize capitalizes the first letter of each word in a string.
func Capitalize(str string) string {
	return strings.Title(strings.ToLower(strings.TrimSpace(str)))
}

// MaskEmail partially masks an email's local part, keeping the first and last characters.
func MaskEmail(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 {
		return email // Return unmasked if invalid email format
	}
	localPart, domain := parts[0], parts[1]

	if len(localPart) <= 2 {
		return localPart + "@" + domain
	}

	maskedLocalPart := string(localPart[0]) + strings.Repeat("*", len(localPart)-2) + string(localPart[len(localPart)-1:])
	return maskedLocalPart + "@" + domain
}

// MaskPhone masks the middle part of a phone number, showing the first and last characters.
func MaskPhone(phone string) string {
	if len(phone) <= 2 {
		return phone
	}
	return phone[:1] + strings.Repeat("*", len(phone)-2) + phone[len(phone)-1:]
}
