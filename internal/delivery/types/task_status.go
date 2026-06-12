package types

import (
	"database/sql/driver"
	"fmt"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "PENDING"
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusCompleted  TaskStatus = "COMPLETED"
)

// For Database with GORM
func (t TaskStatus) Valuer() (driver.Value, error) {
	switch t {
	case StatusPending, StatusInProgress, StatusCompleted:
		return string(t), nil
	default:
		return nil, fmt.Errorf("invalid taks status: %s", t)
	}
}

func (t *TaskStatus) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf("cannot scan nil into TaskStatus")
	}

	var statusStr string
	switch v := value.(type) {
	case []byte:
		statusStr = string(v)
	case string:
		statusStr = v
	default:
		return fmt.Errorf("unsupported data type for TaskStatus: %T", v)
	}

	switch TaskStatus(statusStr) {
	case StatusPending, StatusInProgress, StatusCompleted:
		*t = TaskStatus(statusStr)
	default:
		return fmt.Errorf("invalid task status from database: %s", statusStr)
	}

	return nil
}

// For Validator
func (t TaskStatus) IsValid() bool {
	switch t {
	case StatusPending, StatusInProgress, StatusCompleted:
		return true
	default:
		return false
	}
}
