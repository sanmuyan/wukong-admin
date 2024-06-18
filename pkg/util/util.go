package util

import "github.com/google/uuid"

func GetRandomID() string {
	return uuid.NewString()
}
