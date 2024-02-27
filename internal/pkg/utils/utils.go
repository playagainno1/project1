package utils

import (
	"regexp"
	"strings"
	"time"

	uuid58 "github.com/AlexanderMatveev/go-uuid-base58"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

func NewULID() string {
	return strings.ToLower(ulid.Make().String())
}

func NewBase58UUID() string {
	encoded, err := uuid58.ToBase58(uuid.New())
	if err != nil {
		panic(err)
	}
	return encoded
}

func NewDeviceID() string {
	return uuid.New().String()
}

func UnixMilli() int64 {
	return time.Now().UnixMilli()
}

func IsEmail(s string) bool {
	re := regexp.MustCompile(`^\S+@[a-zA-Z0-9_-]+\.[a-zA-Z]{2,7}$`)
	return re.MatchString(s)
}
