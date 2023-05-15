package storeTest

import (
	"errors"
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
	p []*pb.Post
}

type userStoreTestCases []*userStoreTestCase
type postStoreTestCases []*postStoreTestCase

var (
	globalUserTestCases = userStoreTestCases{

		{
			u: &pb.User{
				Id:       utils.RandomUniqueId() + uint64(time.Now().Unix()),
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
				Id:       utils.RandomUniqueId() + uint64(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "samuel35",
				Email:    "samueljohnson@yahoo.com",
				FullName: "Samuel Johnson",
				Type:     pb.User_T_AUTHOR,
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
				Id:       utils.RandomUniqueId() + uint64(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "laurad42",
				Email:    "lauradavis@hotmail.com",
				FullName: "Laura Davis",
				Type:     pb.User_T_AUTHOR,
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
				Id:       utils.RandomUniqueId() + uint64(time.Now().Unix()),
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
				Id:       utils.RandomUniqueId() + uint64(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "sophiar29",
				Email:    "sophiarodriguez@yahoo.com",
				FullName: "Sophia Rodriguez",
				Type:     pb.User_T_AUTHOR,
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
			u: &pb.User{
				Id:       utils.RandomUniqueId() + uint64(time.Now().Unix()),
				Uuid:     uuid.NewString(),
				Name:     "rommms99",
				Email:    "rommms@gmail.com",
				FullName: "Rom Vales Villanueva",
				Type:     pb.User_T_AUTHOR,
				State: &pb.UserState{
					CreatedAt: timestamppb.Now(),
					UpdatedAt: nil,
					Disabled:  false,
					UVerified: true,
				},
			},
			p: []*pb.Post{
				{
					Id:           utils.RandomUniqueId() + uint64(time.Now().Unix()),
					Uuid:         uuid.NewString(),
					HeadlineText: "The Rise of Sustainable Fashion: How Brands are Embracing Eco-Friendly Practices",
					SummaryText:  "Learn about the growing trend of sustainable fashion and how fashion brands are adopting eco-friendly practices.",
					Tags:         &pb.Tags{Data: []string{"sustainability", "ethical", "eco-conscious"}},
				},
				{
					Id:           utils.RandomUniqueId() + uint64(time.Now().Unix()),
					Uuid:         uuid.NewString(),
					HeadlineText: "Maximizing Your Workout: Tips for Getting the Most Out of Your Gym Session",
					SummaryText:  "Discover effective tips for optimizing your gym workouts to achieve your fitness goals.",
					Tags:         &pb.Tags{Data: []string{"fitness", "exercise", "strength"}},
				},
				{
					Id:           utils.RandomUniqueId() + uint64(time.Now().Unix()),
					Uuid:         uuid.NewString(),
					HeadlineText: "Traveling on a Budget: How to See the World Without Breaking the Bank",
					SummaryText:  "Learn how to travel on a budget with practical tips and tricks for saving money while exploring the world.",
					Tags:         &pb.Tags{Data: []string{"budget", "adventure", "affordable"}},
				},
				{
					Id:           utils.RandomUniqueId() + uint64(time.Now().Unix()),
					Uuid:         uuid.NewString(),
					HeadlineText: "The Benefits of Mindfulness Meditation: How It Can Improve Your Mental Health",
					SummaryText:  "Explore the benefits of mindfulness meditation, including its ability to reduce stress, improve focus, and boost mental well-being.",
					Tags:         &pb.Tags{Data: []string{"mindfulness", "meditation", "wellness"}},
				},
				{
					Id:           utils.RandomUniqueId() + uint64(time.Now().Unix()),
					Uuid:         uuid.NewString(),
					HeadlineText: "Mastering the Art of Cooking: Essential Kitchen Skills for Every Home Chef",
					SummaryText:  "Learn essential kitchen skills for mastering the art of cooking and taking your culinary game to the next level.",
					Tags:         &pb.Tags{Data: []string{"cooking", "culinary", "recipes"}},
				},
				{
					Id:           utils.RandomUniqueId() + uint64(time.Now().Unix()),
					Uuid:         uuid.NewString(),
					HeadlineText: "The Importance of Sleep: How It Affects Your Physical and Mental Health",
					SummaryText:  "Discover the importance of sleep for maintaining good physical and mental health, and learn practical tips for improving your sleep quality.",
					Tags:         &pb.Tags{Data: []string{"sleep", "rest", "recovery"}},
				},
			},
		},
	}
)

var (
	ErrPropNotMatched = errors.New("property not matched")
)
