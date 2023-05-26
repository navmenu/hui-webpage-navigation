package utils_uuid

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func NewUUID() string {
	u := uuid.New()
	return hex.EncodeToString(u[:])
}
