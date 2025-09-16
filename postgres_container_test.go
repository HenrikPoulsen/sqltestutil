package sqltestutil_test

import (
	"context"
	"testing"

	"github.com/HenrikPoulsen/sqltestutil/v2"
	"github.com/stretchr/testify/assert"
)

func TestStartPostgresContainer(t *testing.T) {
	c, err := sqltestutil.StartPostgresContainer(context.Background(), sqltestutil.WithPort(5321))
	if err != nil {
		t.Fatal(err)
	}

	err = c.Shutdown(context.Background())
	assert.NoError(t, err)
}
