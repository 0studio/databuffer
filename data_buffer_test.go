package databuffer

import (
	"fmt"
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
func TestGetBuffer(t *testing.T) {
	data := GetBuffer(SLOT_GIP)
	assert.Equal(t, len(data), SLOT_GIP)
	assert.True(t, cap(data) == SLOT_GIP)

	data2 := GetBuffer(SLOT_GIP - 1)
	assert.Equal(t, len(data2), SLOT_GIP-1)
	assert.True(t, cap(data2) == SLOT_GIP)
	fmt.Printf("这里打印的地址应该是不同的%p,%p\n", data, data2)

}
func TestPutBuffer(t *testing.T) {
	data := GetBuffer(SLOT_GIP)
	assert.Equal(t, len(data), SLOT_GIP)
	assert.True(t, cap(data) == SLOT_GIP)
	PutBuffer(data)

	data2 := GetBuffer(SLOT_GIP - 1)
	assert.Equal(t, len(data2), SLOT_GIP-1)
	assert.True(t, cap(data2) == SLOT_GIP)
	fmt.Printf("这里打印的地址应该是相同的%p,%p\n", data, data2)

}
