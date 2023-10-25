package test

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// these are used for testing when a persistent value is needed
const buildNum = "CIRCLE_BUILD_NUM"
const PersistentWorkspace = "persistent"
const PersistentCollection = "snp"
const PersistentAlias = "getalias"

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//nolint:gosec
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset creates a random string with length and charset
func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString creates a random string with a specific length
func RandomString(length int) string {
	return stringWithCharset(length, charset)
}

const randomNameLength = 6

func RandomName(prefix string) string {
	num, found := os.LookupEnv(buildNum)
	if !found {
		if user, found := os.LookupEnv("USER"); found {
			num = strings.ToLower(user)
		} else {
			num = "dev"
		}
	}

	return fmt.Sprintf("go_%s_%s_%s", num, prefix, RandomString(randomNameLength))
}

func Description() string {
	num, found := os.LookupEnv(buildNum)
	if !found {
		num = "dev"
	}
	return fmt.Sprintf("created by terraform integration test run %s", num)
}
