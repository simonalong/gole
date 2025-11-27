package test

import (
	"github.com/simonalong/gole/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt64ToBase62(t *testing.T) {
	assert.Equal(t, "uRzw7B9", util.Int64ToBase62(1753086935083))
}
