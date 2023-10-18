package sqltestutil_test

import (
	"context"
	"github.com/HenrikPoulsen/sqltestutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartPostgresContainer(t *testing.T) {
	c, err := sqltestutil.StartPostgresContainer(context.Background(), sqltestutil.WithPort(5321))
	if err != nil {
		t.Fatal(err)
	}

	err = c.Shutdown(context.Background())
	assert.NoError(t, err)
}
