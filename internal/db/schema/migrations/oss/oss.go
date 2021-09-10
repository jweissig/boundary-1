package oss

import (
	"embed"

	"github.com/hashicorp/boundary/internal/db/schema"
)

// postgres contains the migrations sql files for postgres oss edition
//go:embed postgres
var postgres embed.FS

func init() {
	schema.RegisterEdition("oss", schema.Postgres, postgres, 0)
}
