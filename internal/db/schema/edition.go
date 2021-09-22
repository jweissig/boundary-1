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

// RegisterEdition registers an edition for use by the Manager. It will panic if:
// - An unsupported dialect is provided.
// - The same (dialect, name) is registered.
// - The same (dialect, priority) is registered.
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

	for _, ee := range e {
		if ee.Name == name {
			panic(fmt.Sprintf("edition %s with dialect %s already registered", name, dialect))
		}

		if ee.Priority == priority {
			panic(fmt.Sprintf("edition %s with dialect %s and priority %d has same priority as edition %s", name, dialect, priority, ee.Name))
		}
	}

	e = append(e, edition.New(name, dialect, fs, priority))
	e.Sort()

	editions[dialect] = e
}
