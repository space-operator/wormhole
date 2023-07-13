package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeaderSet(t *testing.T) {
	assert.Equal(t, 7, calculateLeaderSetSize(19)) // 19 guardians = 13 for quorum, 7 in leaderset
	assert.Equal(t, 7, calculateLeaderSetSize(18)) // 18 guardians = 13 for quorum, 7 in leaderset
	assert.Equal(t, 6, calculateLeaderSetSize(17)) // 18 guardians = 12 for quorum, 6 in leaderset
}
