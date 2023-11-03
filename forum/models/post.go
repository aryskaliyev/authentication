package models

import (
	"time"
)

type Post struct {
	Id         int        `json:"id,omitempty"`
	Title      string     `json:"title"`
	Body       string     `json:"body"`
	Created    *time.Time `json:"created,omitempty"`
	Categories []PostCategory      `json:"category"`
}
