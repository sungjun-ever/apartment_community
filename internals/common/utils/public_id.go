package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GeneratePublicId() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}
