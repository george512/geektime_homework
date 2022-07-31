package hystrix

import (
	"log"
	"sync"
	"time"
)

type RollWindow struct {
	sync.RWMutex
	// 是否处于熔断状态
	fusing bool
	// 滑动窗口数量
	size int
	// 桶队列
	buckets []*bucket
	// 触发熔断的请求总量
	rate int
	// 触发熔断的失败率
	failedRate float64
	// 上次熔断发生时间
	lastFusingTime time.Time
	// 熔断恢复的时间间隔
	fusingBackTime time.Duration
}

func NewRollWindow(size int, rate int, failedRate float64, fusingBackTime time.Duration) *RollWindow {
	return &RollWindow{
		size:           size,
		buckets:        make([]*bucket, 0),
		rate:           rate,
		failedRate:     failedRate,
		fusingBackTime: fusingBackTime,
	}
}

// 追加新桶
func (r *RollWindow) AppendBucket() {
	r.Lock()
	defer r.Unlock()

	r.buckets = append(r.buckets, NewBucket())
	// 如果桶的数量大于size, 将第一个桶出队
	if len(r.buckets) > r.size+1 {
		r.buckets = r.buckets[1:]
	}
}

// 获取队尾桶
func (r *RollWindow) GetLastBucket() *bucket {
	if len(r.buckets) == 0 {
		r.AppendBucket()
	}
	return r.buckets[len(r.buckets)-1]
}

// 记录请求结果
func (r *RollWindow) RecordRes(res bool) {
	r.GetLastBucket().Record(res)
}

// 启动滑动窗口, 每0.1秒在队尾增加一个桶
func (r *RollWindow) Launch() {
	go func() {
		r.AppendBucket()
		time.Sleep(time.Millisecond * 100)
	}()
}

// 判断是否需要触发熔断
func (r *RollWindow) FusingJudge() bool {
	r.RLock()
	defer r.RUnlock()
	var total, failed int
	for _, v := range r.buckets {
		total += v.Total
		failed += v.Failed
	}
	if float64(failed)/float64(total) >= r.failedRate && total > r.rate {
		return true
	}
	return false
}

// 是否到了熔断恢复时间
func (r *RollWindow) OverFusingBackTime() bool {
	return time.Since(r.lastFusingTime) > r.fusingBackTime

}

// 监控请求失败率, 触发/恢复熔断
func (r *RollWindow) Monitor() {
	go func() {
		for {
			// 已经触发熔断，判断是否到了恢复时间
			if r.fusing {
				if r.OverFusingBackTime() {
					r.fusing = false
				}
			} else {
				// 触发熔断
				if r.FusingJudge() {
					r.Lock()
					r.fusing = true
					r.lastFusingTime = time.Now()
					r.Unlock()
				}
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()
}

// 获取当前熔断状态
func (r *RollWindow) FusingStatus() bool {
	r.RLock()
	defer r.RUnlock()
	return r.fusing
}

// 每隔一秒展示当前熔断状态
func (r *RollWindow) ShowStatus() {
	go func() {
		for {
			log.Print("当前熔断状态: ", r.fusing)
			time.Sleep(1 * time.Second)
		}
	}()
}

// 每隔一秒展示当前total和failed
func (r *RollWindow) ShowTotalStatus() {
	go func() {
		for {
			r.RLock()
			var total, failed int
			for _, v := range r.buckets {
				total += v.Total
				failed += v.Failed
			}
			log.Print("total: ", total, ", failed: ", failed)
			r.RUnlock()
			time.Sleep(1 * time.Second)
		}
	}()
}
