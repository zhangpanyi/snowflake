package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	workerIDBits int64 = 10
	sequenceBits int64 = 12
	maxWorkerID  int64 = -1 ^ (-1 << workerIDBits)
	maxSequence  int64 = -1 ^ (-1 << sequenceBits)
)

// Worker 工作节点
type Worker struct {
	id        int64
	sequence  int64
	offset    int64
	timeStamp int64
	mutex     sync.Mutex
}

// NewWorker 创建节点
func NewWorker(id, offset int64) (*Worker, error) {
	if id > maxWorkerID {
		return nil, errors.New("worker id out of range")
	}
	return &Worker{id: id}, nil
}

// Generate 生成ID
func (worker *Worker) Generate() int64 {
	worker.mutex.Lock()
	now := time.Now().UnixNano() / int64(time.Millisecond)
	if worker.timeStamp != now {
		worker.sequence = 0
		worker.timeStamp = now
	} else {
		worker.sequence++
		if worker.sequence > maxSequence {
			worker.mutex.Unlock()
			time.Sleep(time.Millisecond - time.Duration(now)%time.Millisecond)
			return worker.Generate()
		}
	}
	now = now - worker.offset
	id := int64(now<<(workerIDBits+sequenceBits) | worker.id<<sequenceBits | worker.sequence)
	worker.mutex.Unlock()
	return id
}
