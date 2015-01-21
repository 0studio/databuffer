package databuffer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMatchSize(t *testing.T) {
	assert.Equal(t, SLAB_GIP, getMatchSize(SLAB_GIP-1))
	assert.Equal(t, SLAB_GIP, getMatchSize(SLAB_GIP))
	assert.Equal(t, SLAB_GIP*2, getMatchSize(SLAB_GIP+1))

	assert.Equal(t, SLAB_INIT_SIZE, getMatchSize(SLAB_INIT_SIZE))
	assert.Equal(t, SLAB_INIT_SIZE, getMatchSize(SLAB_INIT_SIZE-1))

	assert.Equal(t, SLAB_INIT_SIZE*GROWTH_FACTOR/100, getMatchSize(SLAB_INIT_SIZE+1))
	assert.Equal(t, SLAB_INIT_SIZE*GROWTH_FACTOR/100, getMatchSize(SLAB_INIT_SIZE*GROWTH_FACTOR/100))

}
func TestGetBuffer(t *testing.T) {
	data := GetBuffer(SLAB_GIP)
	assert.Equal(t, len(data), SLAB_GIP)
	assert.True(t, cap(data) == SLAB_GIP)

	data2 := GetBuffer(SLAB_GIP - 1)
	assert.Equal(t, len(data2), SLAB_GIP-1)
	assert.True(t, cap(data2) == SLAB_GIP)
	fmt.Printf("这里打印的地址应该是不同的%p,%p\n", data, data2)

}
func TestPutBuffer(t *testing.T) {
	data := GetBuffer(SLAB_GIP)
	assert.Equal(t, len(data), SLAB_GIP)
	assert.True(t, cap(data) == SLAB_GIP)
	PutBuffer(data)

	data2 := GetBuffer(SLAB_GIP - 1)
	assert.Equal(t, len(data2), SLAB_GIP-1)
	assert.True(t, cap(data2) == SLAB_GIP)
	fmt.Printf("这里打印的地址应该是相同的%p,%p\n", data, data2)

}
