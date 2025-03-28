// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package queries

import (
	"time"
)

type Device struct {
	ID           int32     `db:"id" json:"id"`
	Name         string    `db:"name" json:"name,omitempty"`
	Brand        string    `db:"brand" json:"brand,omitempty"`
	State        string    `db:"state" json:"state,omitempty"`
	CreationTime time.Time `db:"creation_time" json:"creation_time"`
}
