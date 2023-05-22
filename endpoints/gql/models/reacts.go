package models

import "github.com/romv7/blogs/internal/pb"

type Reacts struct {
	LikeCount     uint64 `json:"likeCount"`
	ConfusedCount uint64 `json:"confusedCount"`
	LoveCount     uint64 `json:"loveCount"`
	LaughCount    uint64 `json:"laughCount"`
	SadCount      uint64 `json:"sadCount"`
	CareCount     uint64 `json:"careCount"`
}

func Proto_GQLModelReacts(r *pb.Reacts) *Reacts {
	return &Reacts{
		LikeCount:     r.LikeCount,
		ConfusedCount: r.ConfusedCount,
		LoveCount:     r.LoveCount,
		LaughCount:    r.LaughCount,
		SadCount:      r.SadCount,
		CareCount:     r.CareCount,
	}
}

func (r *Reacts) Proto() *pb.Reacts {
	return &pb.Reacts{
		LikeCount:     r.LikeCount,
		ConfusedCount: r.ConfusedCount,
		LoveCount:     r.LoveCount,
		LaughCount:    r.LaughCount,
		SadCount:      r.SadCount,
		CareCount:     r.CareCount,
	}
}
