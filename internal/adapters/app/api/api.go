package api

import "github.com/isuraem/hex/internal/ports"

type Adapter struct {
	db ports.DBport
}

func NewAdapter(db ports.DBport) *Adapter {
	return &Adapter{db: db}
}
