package databuffer

import (
	"sync"
)

// 客户端发过来的消息长度 基本上是稳定的某几个数值之间，
// 使用sync.Pool 对make([]byte,n) 进行内存重用
// 尽量减少gc时间
var (
	bufMap map[int]*sync.Pool = make(map[int]*sync.Pool)
)

func getMatchSize(size int) (ret int) {
	// 找到合适的大小
	if size <= SLOT_INIT_SIZE {
		if size%SLOT_GIP == 0 {
			return size
		}
		return (size/SLOT_GIP + 1) * SLOT_GIP
	}
	return getMatchSlotSize(size)
}

const (
	SLOT_GIP = 8 // 当大小小于SLOT_INIT_SIZE byte时，所需内存按每SLOT_GIP byte 递增

	SLOT_RATIO     = 120 // 120 means 1.2,must >100
	SLOT_INIT_SIZE = 512
	// SLOT_TRY_CNT   = 16
)

func getMatchSlotSize(size int) (slotSize int) {
	slotSize = SLOT_INIT_SIZE
	for {
		// for i := 0; i < SLOT_TRY_CNT; i++ {
		slotSize = slotSize * SLOT_RATIO / 100
		if size > slotSize {
			continue
		}
		return slotSize
	}
	// return 0 // 0 means ,doesnot find match slot
}

func getPool(size int) (pool *sync.Pool) {
	// for key, _ := range bufMap {
	// 	log.Debug("buf cache size=", key)

	// }

	if pool = bufMap[size]; pool == nil {
		pool = &sync.Pool{}
		bufMap[size] = pool
	}
	return
}
func GetBuffer(size int) (byteSlic []byte) {
	matchSize := getMatchSize(size)
	if matchSize < size {
	}

	if v := getPool(matchSize).Get(); v != nil {
		byteSlic = v.([]byte)
		byteSlic = byteSlic[0:size]
		return
	}
	byteSlic = make([]byte, size, matchSize)
	return
}
func PutBuffer(byteSlic []byte) {
	getPool(getMatchSize(cap(byteSlic))).Put(byteSlic)
}
