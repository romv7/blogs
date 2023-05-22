package models

import "time"

type UserOpsType uint

const (
	UserOpsType_UO_CREATE_POST = iota
	UserOpsType_UO_UPDATE_POST
	UserOpsType_UO_DELETE_POST
	UserOpsType_UO_CREATE_COMMENT
	UserOpsType_UO_UPDATE_COMMENT
	UserOpsType_UO_DELETE_COMMENT
)

func (so UserOpsType) String() string {
	var enums = []string{
		"UO_CREATE_POST", "UO_UPDATE_POST", "UO_DELETE_POST",
		"UO_CREATE_COMMENT", "UO_UPDATE_COMMENT", "UO_DELETE_COMMENT",
	}

	return enums[so]
}

type UserOpsResults struct {
	Op        SuperOpsType `json:"op"`
	Message   *string      `json:"message"`
	Code      int32        `json:"code"`
	StartTime time.Time    `json:"startTime"`
	EndTime   time.Time    `json:"endTime"`
}
