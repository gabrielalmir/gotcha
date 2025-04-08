package utils

import (
	"fmt"
	"sync"
	"time"
)

const (
	nodeIDBits   = 10
	sequenceBits = 12
	nodeIDMax    = -1 ^ (-1 << nodeIDBits)
	sequenceMask = -1 ^ (-1 << sequenceBits)
	timeShift    = nodeIDBits + sequenceBits
	nodeIDShift  = sequenceBits
	epoch        = 1288834974657 // Twitter epoch (Nov 04 2010 01:42:54.657)
)

type Snowflake struct {
	mu        sync.Mutex
	lastStamp int64
	nodeID    int64
	sequence  int64
}

func NewSnowflake(nodeID int64) (*Snowflake, error) {
	if nodeID < 0 || nodeID > nodeIDMax {
		return nil, fmt.Errorf("node ID must be between 0 and %d", nodeIDMax)
	}

	return &Snowflake{
		nodeID: nodeID,
	}, nil
}

func (s *Snowflake) Generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	timestamp := time.Now().UnixMilli()

	if timestamp < s.lastStamp {
		timestamp = s.lastStamp
	}

	if s.lastStamp == timestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			timestamp = s.waitNextMilli(timestamp)
		}
	} else {
		s.sequence = 0
	}

	s.lastStamp = timestamp

	id := ((timestamp - epoch) << timeShift) |
		(s.nodeID << nodeIDShift) |
		s.sequence

	return id
}

func (s *Snowflake) waitNextMilli(timestamp int64) int64 {
	for timestamp <= s.lastStamp {
		timestamp = time.Now().UnixMilli()
	}
	return timestamp
}

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func ToBase62(n int64) string {
	if n == 0 {
		return string(base62Chars[0])
	}

	var result []byte
	for n > 0 {
		result = append(result, base62Chars[n%62])
		n /= 62
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
