package loggers

import (
	"testing"

	. "github.com/hyperledger/burrow/util/slice"
	"github.com/stretchr/testify/assert"
)

func TestFilterLogger(t *testing.T) {
	testLogger := NewChannelLogger(100)
	filterLogger := NewFilterLogger(testLogger, func(keyvals []interface{}) bool {
		return len(keyvals) > 0 && keyvals[0] == "Spoon"
	})
	filterLogger.Log("Fish", "Present")
	filterLogger.Log("Spoon", "Present")
	assert.Equal(t, [][]interface{}{Slice("Fish", "Present")}, testLogger.FlushLogLines())
}
