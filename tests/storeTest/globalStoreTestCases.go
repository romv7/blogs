package storeTest

import (
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userStoreTestCase struct {
	u *pb.User
}

type postStoreTestCase struct {
	u *pb.User
	p *pb.Post
}

type userStoreTestCases []*userStoreTestCase
type postStoreTestCases []*postStoreTestCase

var (
	globalUserTestCases = userStoreTestCases{
		{
			u: &pb.User{
				Id:       utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "julie09",
				Email:    "juliekelly@gmail.com",
				FullName: "Julie Kelly",
				Type:     pb.User_T_NORMAL,
				State: &pb.UserState{
					CreatedAt: timestamppb.Now(),
					UpdatedAt: nil,
					Disabled:  false,
					UVerified: false,
				},
			},
		},

		{
			u: &pb.User{
				Id:       utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "samuel35",
				Email:    "samueljohnson@yahoo.com",
				FullName: "Samuel Johnson",
				Type:     pb.User_T_NORMAL,
				State: &pb.UserState{
					CreatedAt: timestamppb.Now(),
					UpdatedAt: nil,
					Disabled:  false,
					UVerified: false,
				},
			},
		},
		{
			u: &pb.User{
				Id:       utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "laurad42",
				Email:    "lauradavis@hotmail.com",
				FullName: "Laura Davis",
				Type:     pb.User_T_NORMAL,
				State: &pb.UserState{
					CreatedAt: timestamppb.Now(),
					UpdatedAt: nil,
					Disabled:  false,
					UVerified: false,
				},
			},
		},
		{
			u: &pb.User{
				Id:       utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "adam25",
				Email:    "adamsmith@gmail.com",
				FullName: "Adam Smith",
				Type:     pb.User_T_NORMAL,
				State: &pb.UserState{
					CreatedAt: timestamppb.Now(),
					UpdatedAt: nil,
					Disabled:  false,
					UVerified: false,
				},
			},
		},
		{
			u: &pb.User{
				Id:       utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "sophiar29",
				Email:    "sophiarodriguez@yahoo.com",
				FullName: "Sophia Rodriguez",
				Type:     pb.User_T_NORMAL,
				State: &pb.UserState{
					CreatedAt: timestamppb.Now(),
					UpdatedAt: nil,
					Disabled:  false,
					UVerified: false,
				},
			},
		},
	}

	globalPostTestCases = postStoreTestCases{

		{
			u: globalUserTestCases[getRandIndex()].u,
			p: &pb.Post{
				Id:           utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:         uuid.NewString(),
				HeadlineText: "How to Choose the Right Fitness Program for Your Goals",
				SummaryText:  "",
				Tags:         &pb.Tags{Data: []string{"fitness", "exercise", "workout"}},
				State: &pb.PostState{
					Stage:       pb.PostState_S_WIP,
					Status:      pb.PostState_S_DRAFT,
					CreatedAt:   timestamppb.Now(),
					PublishedAt: nil,
					RevisedAt:   nil,
					ArchivedAt:  nil,
					Reacts:      &pb.Reacts{},
				},
			},
		},
		{
			u: globalUserTestCases[getRandIndex()].u,
			p: &pb.Post{
				Id:           utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:         uuid.NewString(),
				HeadlineText: "The Top 5 Benefits of Meditation for Your Mental Health",
				SummaryText:  "",
				Tags:         &pb.Tags{Data: []string{"meditation", "mindfulness", "mental health"}},
				State: &pb.PostState{
					Stage:       pb.PostState_S_PUB,
					Status:      pb.PostState_S_DRAFT,
					CreatedAt:   timestamppb.Now(),
					PublishedAt: timestamppb.Now(),
					RevisedAt:   nil,
					ArchivedAt:  nil,
					Reacts:      &pb.Reacts{},
				},
			},
		},
		{
			u: globalUserTestCases[getRandIndex()].u,
			p: &pb.Post{
				Id:           utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:         uuid.NewString(),
				HeadlineText: "The Benefits of Working Remotely: Why More Companies are Embracing Remote Work",
				SummaryText:  "",
				Tags:         &pb.Tags{Data: []string{"remote work", "telecommuting", "work-life balance"}},
				State: &pb.PostState{
					Stage:       pb.PostState_S_WIP,
					Status:      pb.PostState_S_DRAFT,
					CreatedAt:   timestamppb.Now(),
					PublishedAt: timestamppb.Now(),
					RevisedAt:   nil,
					ArchivedAt:  nil,
					Reacts:      &pb.Reacts{},
				},
			},
		},
		{
			u: globalUserTestCases[getRandIndex()].u,
			p: &pb.Post{
				Id:           utils.RandomUniqueId() + uint32(time.Now().Unix()),
				Uuid:         uuid.NewString(),
				HeadlineText: "The Future of Artificial Intelligence: Trends to Watch in 2023 and Beyond",
				SummaryText:  "",
				Tags:         &pb.Tags{Data: []string{"artificial intelligence", "machine learning", "automation"}},
				State: &pb.PostState{
					Stage:       pb.PostState_S_PUB,
					Status:      pb.PostState_S_DRAFT,
					CreatedAt:   timestamppb.Now(),
					PublishedAt: timestamppb.Now(),
					RevisedAt:   nil,
					ArchivedAt:  nil,
					Reacts:      &pb.Reacts{},
				},
			},
		},
	}
)

var (
	ErrPropNotMatched = errors.New("property not matched")
)

func getRandIndex() int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(len(globalUserTestCases) + 1 - 1)
}
