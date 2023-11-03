package models

import (
	"time"
)

type Category struct {
	Id      int        `json:"id,omitempty"`
	Name    string     `json:"name"`
	Created *time.Time `json:"created,omitempty"`
}
