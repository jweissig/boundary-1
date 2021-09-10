package schema

import "github.com/hashicorp/boundary/internal/db/schema/internal/edition"

// PartialEditions is used by CreatePartialEditions. It is a map of edition
// names to the max version that should be included.
type PartialEditions map[string]int

// CreatePartialEditions is used by tests to create a subset of the Edition migrations.
func CreatePartialEditions(dialect Dialect, p PartialEditions) []edition.Edition {
	e := make([]edition.Edition, 0, len(p))
	for _, ee := range editions[dialect] {
		maxVer, ok := p[ee.Name]
		if ok {
			edition := edition.Edition{
				Name:          ee.Name,
				Dialect:       ee.Dialect,
				Priority:      ee.Priority,
				LatestVersion: nilVersion,
				Migrations:    make(map[int][]byte),
			}

			for k, b := range ee.Migrations {
				if k > maxVer {
					continue
				}

				edition.Migrations[k] = b
				if k > edition.LatestVersion {
					edition.LatestVersion = k
				}
			}
			e = append(e, edition)
		}
	}

	return e
}
