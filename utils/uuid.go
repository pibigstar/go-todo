package utils

import (
	"strings"

	"github.com/pborman/uuid"
)

// GetUUID 获取uuid
func GetUUID() string {
	uid := uuid.NewRandom().String()
	return strings.Replace(uid, "-", "", -1)
}
