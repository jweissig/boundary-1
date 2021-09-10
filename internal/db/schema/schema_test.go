package schema_test

import (
	"context"
	"testing"

	"github.com/hashicorp/boundary/internal/db/schema"
	"github.com/hashicorp/boundary/testing/dbtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMigrateStore(t *testing.T) {
	dialect := dbtest.Postgres
	ctx := context.Background()

	c, u, _, err := dbtest.StartUsingTemplate(dialect, dbtest.WithTemplate(dbtest.Template1))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, c())
	})

	ran, err := schema.MigrateStore(ctx, dialect, u, schema.WithEditions(
		schema.CreatePartialEditions(schema.Dialect(dialect), schema.PartialEditions{"oss": 1}),
	))
	assert.NoError(t, err)
	assert.True(t, ran)

	ran, err = schema.MigrateStore(ctx, dialect, u, schema.WithEditions(
		schema.CreatePartialEditions(schema.Dialect(dialect), schema.PartialEditions{"oss": 1}),
	))
	assert.NoError(t, err)
	assert.False(t, ran)

	ran, err = schema.MigrateStore(ctx, dialect, u, schema.WithEditions(
		schema.CreatePartialEditions(schema.Dialect(dialect), schema.PartialEditions{"oss": 2}),
	))
	assert.NoError(t, err)
	assert.True(t, ran)
	ran, err = schema.MigrateStore(ctx, dialect, u, schema.WithEditions(
		schema.CreatePartialEditions(schema.Dialect(dialect), schema.PartialEditions{"oss": 2}),
	))
	assert.NoError(t, err)
	assert.False(t, ran)
}
