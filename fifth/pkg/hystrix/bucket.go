package hystrix

import (
	"sync"
	"time"
)

type bucket struct {
	sync.RWMutex
	// 请求总数
	Total int
	// 失败总数
	Failed int
	// 创建时间
	CreateTime time.Time
}

func NewBucket() *bucket {
	return &bucket{
		CreateTime: time.Now(),
	}
}

// 记录请求结果
func (b *bucket) Record(res bool) {
	b.Lock()
	defer b.Unlock()

	if !res {
		b.Failed++
	}
	b.Total++
}
