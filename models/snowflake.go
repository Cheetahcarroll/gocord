package models

import "time"

// Snowflake is a uint64 alias for a snowflake in discord
type Snowflake uint64

// SnowStruct is an object to help build snowflakes
type SnowStruct struct {
	TimeStamp time.Time
	WorkerID  uint8
	ProcessID uint8
	Increment uint16
}

// BuildSnowflake uses a SnowStruct to build a snowflake
func (s *SnowStruct) BuildSnowflake() Snowflake {
	result := uint64(uint64((time.Now().UnixNano()/int64(time.Millisecond))-1420070400000) << 22)
	result += uint64(s.WorkerID) << 17
	result += uint64(s.ProcessID) << 12
	result += uint64(s.Increment)
	return Snowflake(result)
}

// GetTimeStamp gets the timestamp of a snowflake
func (s Snowflake) GetTimeStamp() time.Time {
	return time.Unix(0, int64(s>>22+1420070400000)*int64(time.Millisecond))
}

// GetWorkerID gets the worker id of a snowflake
func (s Snowflake) GetWorkerID() uint8 {
	return uint8(s << 42 >> 59)
}

// GetProcessID gets the process id of a snowflake
func (s Snowflake) GetProcessID() uint8 {
	return uint8(s << 47 >> 59)
}

// GetIncrement gets the increment value of a snowflake
func (s Snowflake) GetIncrement() uint16 {
	return uint16(s << 52 >> 52)
}
