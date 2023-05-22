package models

import "time"

type SuperOpsType uint

const (
	SuperOpsType_SO_NEW SuperOpsType = iota
	SuperOpsType_SO_UPDATE
	SuperOpsType_SO_DELETE
)

func (so SuperOpsType) String() string {
	var enums = []string{
		"SO_NEW", "SO_UPDATE", "SO_DELETE",
	}

	return enums[so]
}

type SuperOpsResults struct {
	Op        SuperOpsType `json:"op"`
	Message   *string      `json:"message"`
	Code      int32        `json:"code"`
	StartTime time.Time    `json:"startTime"`
	EndTime   time.Time    `json:"endTime"`
	Uuid      string       `json:"resource_uuid"`
}
