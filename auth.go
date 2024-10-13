package gokit

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

// GenerateAuthCode creates a base64-encoded, URL-safe random string of 20 bytes.
func GenerateAuthCode() (string, error) {
	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// GetUserIdFromCtx extracts a user ID from context metadata, returning uuid.Nil if not found.
func GetUserIdFromCtx(ctx context.Context) uuid.UUID {
	userId := uuid.Nil
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md["user-id"]) > 0 {
		parsed, err := uuid.Parse(md["user-id"][0])
		if err == nil {
			userId = parsed
		}
	}
	return userId
}
