package domain

import "time"

type RefreshSession struct {
	ID        int64
	UseId     int64
	Token     string
	ExpiresAt time.Time
}
