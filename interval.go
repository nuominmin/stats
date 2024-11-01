package stats

import (
	"sync"
	"time"
)

// 记录调用间隔统计信息
type intervalStats struct {
	lastTime  time.Time     // 上一次调用时间
	interval  time.Duration // 最近一次间隔
	avg       time.Duration // 平均间隔
	totalTime time.Duration // 总调用时间
	count     int64         // 总调用次数
	mu        sync.Mutex
}

// 全局统计对象
var stats = &intervalStats{
	lastTime:  time.Now(),
	totalTime: 0,
	count:     0,
}

// UpdateStats 更新调用间隔统计信息
func UpdateStats() {
	stats.mu.Lock()
	defer stats.mu.Unlock()

	now := time.Now()
	if !stats.lastTime.IsZero() {
		stats.interval = now.Sub(stats.lastTime)
		stats.totalTime += stats.interval
		stats.count++

		// 计算平均间隔
		stats.avg = stats.totalTime / time.Duration(stats.count)
	}
	stats.lastTime = now
}

func ReadInterval() time.Duration {
	return stats.interval
}

func ReadAvg() time.Duration {
	return stats.avg
}

func ReadCount() int64 {
	return stats.count
}
