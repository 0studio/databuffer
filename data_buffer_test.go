package databuffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMatchSize(t *testing.T) {
	assert.Equal(t, SLOT_GIP, getMatchSize(SLOT_GIP-1))
	assert.Equal(t, SLOT_GIP, getMatchSize(SLOT_GIP))
	assert.Equal(t, SLOT_GIP*2, getMatchSize(SLOT_GIP+1))

	assert.Equal(t, SLOT_INIT_SIZE, getMatchSize(SLOT_INIT_SIZE))
	assert.Equal(t, SLOT_INIT_SIZE, getMatchSize(SLOT_INIT_SIZE-1))

	assert.Equal(t, SLOT_INIT_SIZE*SLOT_RATIO/100, getMatchSize(SLOT_INIT_SIZE+1))
	assert.Equal(t, SLOT_INIT_SIZE*SLOT_RATIO/100, getMatchSize(SLOT_INIT_SIZE*SLOT_RATIO/100))

}
