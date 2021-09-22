package schema

import (
	"embed"
	"fmt"

	"github.com/hashicorp/boundary/internal/db/schema/internal/edition"
)

// Dialect is same as edition.Dialect
type Dialect = edition.Dialect

// Supported dialects.
const (
	Postgres Dialect = "postgres"
)

var supportedDialects = map[Dialect]struct{}{
	Postgres: struct{}{},
}

type dialects map[Dialect]edition.Editions

var editions = make(dialects)

// RegisterEdition registers an edition.Edition for use by the Manager.
func RegisterEdition(name string, dialect Dialect, fs embed.FS, priority int) {
	if _, ok := supportedDialects[dialect]; !ok {
		panic(fmt.Sprintf("unsupported dialect: %s", dialect))
	}
	var e edition.Editions
	var ok bool

	e, ok = editions[dialect]
	if !ok {
		e = make(edition.Editions, 0)
	}

	e = append(e, edition.New(name, dialect, fs, priority))
	e.Sort()

	editions[dialect] = e
}
