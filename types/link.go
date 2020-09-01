package types

import (
	"time"
)

type Link struct {
	ID      int       `storm:"increment" json:"id"`
	Name    string    `json:"name"`
	Serial  string    `json:"serial" storm:"index"`
	Addr    string    `json:"addr"`
	Channel int       `json:"channel"`
	Created time.Time `json:"created"`
}