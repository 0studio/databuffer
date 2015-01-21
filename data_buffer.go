package databuffer

import (
	"sync"
)

// 客户端发过来的消息长度 基本上是稳定的某几个数值之间，
// 使用sync.Pool 对make([]byte,n) 进行内存重用
// 尽量减少gc时间
var (
	bufMap   map[int]*sync.Pool = make(map[int]*sync.Pool)
	poolLock sync.RWMutex       = sync.RWMutex{}
)

func getMatchSize(size int) (ret int) {

	// 找到合适的大小
	if size <= SLAB_INIT_SIZE {
		if size%SLAB_GIP == 0 {
			return size
		}
		return (size/SLAB_GIP + 1) * SLAB_GIP
	}
	return getMatchSlabSize(size)
}

const (
	SLAB_GIP = 8 // 当大小小于SLAB_INIT_SIZE byte时，所需内存按每SLAB_GIP byte 递增

	GROWTH_FACTOR  = 125 // 125 means 1.25,must >100
	SLAB_INIT_SIZE = 512
	// SLAB_TRY_CNT   = 16
)

func getMatchSlabSize(size int) (slabSize int) {
	slabSize = SLAB_INIT_SIZE
	for {
		// for i := 0; i < SLAB_TRY_CNT; i++ {
		slabSize = slabSize * GROWTH_FACTOR / 100
		if size > slabSize {
			continue
		}
		return slabSize
	}
	// return 0 // 0 means ,doesnot find match slab
}

func getPool(size int) (pool *sync.Pool) {
	// for key, _ := range bufMap {
	// 	log.Debug("buf cache size=", key)

	// }
	var tmpPool *sync.Pool
	if pool = bufMap[size]; pool == nil {
		poolLock.Lock()
		defer poolLock.Unlock()
		pool = &sync.Pool{}
		if tmpPool = bufMap[size]; tmpPool == nil {
			// 再次判断一下map 里是否已经有了，
			bufMap[size] = pool
		} else {
			pool = tmpPool
		}
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
