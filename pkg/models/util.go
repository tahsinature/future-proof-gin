package models

import (
	"encoding/json"
)

// UserSessionInfo ...
type UserSessionInfo struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// JSONRaw ...
type JSONRaw json.RawMessage

// DataList ....
type DataList struct {
	Data JSONRaw `db:"data" json:"data"`
	Meta JSONRaw `db:"meta" json:"meta"`
}
