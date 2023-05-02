package sql

import (
	"encoding/json"

	"github.com/romv7/blogs/pb"
)

// This auxiliary function will marshall a blogs.Reacts to a JSON by
// using a simple map.
func ReactsMarshall(reacts *pb.Reacts) string {

	uuids := make([]string, 0)

	for _, user := range reacts.Users {
		uuids = append(uuids, user.Uuid)
	}

	p, _ := json.Marshal(map[string]any{
		"likeCount":     reacts.LikeCount,
		"confusedCount": reacts.ConfusedCount,
		"loveCount":     reacts.LoveCount,
		"laughCount":    reacts.LaughCount,
		"sadCount":      reacts.SadCount,
		"careCount":     reacts.CareCount,
		"users":         uuids,
	})

	return string(p)
}
