package grpcTest

import (
	"github.com/romv7/blogs/internal/pb"
)

var (
	globalTestCases = grpcTestCases{
		{
			For: UserServiceTest,
			testValues: map[string]any{
				"email":     "david.homenick@email.com",
				"full_name": "David Homenick",
				"name":      "sickneck.david12",
				"type":      0,
			},
		},

		{
			For: BlogServiceTest_Comment,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "rommmms074@gmail.com",
					"full_name": "Rom Vales Villanueva",
					"name":      "romdevmod123",
					"type":      pb.User_T_AUTHOR,
				},
				"comment_text": "Cool post! I'm waiting to know what will be next!",
			},
		},

		{
			For: BlogServiceTest,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "user2@gmail.com",
					"full_name": "Christian C. Balunso",
					"name":      "username123",
					"type":      pb.User_T_AUTHOR,
				},

				"headline_text": "Upcoming Events",
				"summary_text":  "Mark your calendars for exciting events",
				"tags":          []string{"events", "community", "meetups"},
				"images":        []string{},
				"attachments":   []string{},
				"refs":          []string{"ref7", "ref8", "ref9"},
			},
		},

		{
			For: BlogServiceTest_Comment,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "user60@gmail.com",
					"full_name": "Andry Manlangit",
					"name":      "andry0123",
					"type":      pb.User_T_NORMAL,
				},
				"comment_text": "Okay okay okay",
			},
		},

		{
			For: BlogServiceTest,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "user1@gmail.com",
					"full_name": "John Mayer",
					"name":      "username123",
					"type":      pb.User_T_AUTHOR,
				},

				"headline_text": "Breaking News",
				"summary_text":  "Stay informed with the latest updates",
				"tags":          []string{"news", "current affairs"},
				"images":        []string{},
				"attachments":   []string{"attachment4.docx"},
				"refs":          []string{"ref5", "ref6"},
			},
		},

		{
			For: BlogServiceTest_Comment,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "emil@gmail.com",
					"full_name": "Emil Cioran",
					"name":      "existenstialdread",
					"type":      pb.User_T_AUTHOR,
				},
				"comment_text": "Fuck is this bro? You're writing a shitty essay again.",
			},
		},

		{
			For: BlogServiceTest,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "user1@gmail.com",
					"full_name": "John Mayer",
					"name":      "username123",
					"type":      pb.User_T_AUTHOR,
				},

				"headline_text": "Upcoming Events",
				"summary_text":  "Mark your calendars for exciting events",
				"tags":          []string{"events", "community", "meetups"},
				"images":        []string{},
				"attachments":   []string{},
				"refs":          []string{"ref7", "ref8", "ref9"},
			},
		},
		{
			For: UserServiceTest,
			testValues: map[string]any{
				"email":     "romdevmod@gmail.com",
				"full_name": "Rom Vales Villanueva",
				"name":      "rommthefox",
				"type":      0,
			},
		},

		{
			For: BlogServiceTest_Comment,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "ericblair000@gmail.com",
					"full_name": "George Orwell",
					"name":      "ericb0123",
					"type":      pb.User_T_AUTHOR,
				},
				"comment_text": "We're already living in our world now. Fuck the politicians and government.",
			},
		},

		{
			For: UserServiceTest,
			testValues: map[string]any{
				"email":     "dante@dmcheadquarters.com",
				"full_name": "Dante the Demon Slayer",
				"name":      "dantethejester",
				"type":      0,
			},
		},

		{
			For: BlogServiceTest,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "user1@gmail.com",
					"full_name": "John Mayer",
					"name":      "username123",
					"type":      pb.User_T_AUTHOR,
				},

				"headline_text": "Lorem ipsum dolor sit amet",
				"summary_text":  "Sed ut perspiciatis unde omnis iste natus error sit voluptatem",
				"tags":          []string{"technology", "programming", "web development"},
				"images":        []string{"image1.jpg", "image2.jpg", "image3.jpg"},
				"attachments":   []string{"attachment1.pdf", "attachment2.docx"},
				"refs":          []string{"ref1", "ref2", "ref3"},
			},
		},

		{
			For: UserServiceTest,
			testValues: map[string]any{
				"email":     "nero@dmcheadquarters.com",
				"full_name": "Nero (Son of Vergil)",
				"name":      "nerodmc12",
				"type":      0,
			},
		},

		{
			For: BlogServiceTest_Comment,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "user6@gmail.com",
					"full_name": "Shane Cheska Jane",
					"name":      "shane123",
					"type":      pb.User_T_NORMAL,
				},
				"comment_text": "This post needs to be revised more, there were a lot of typographical mistakes in the writing though the information is good.",
			},
		},

		{
			For: UserServiceTest,
			testValues: map[string]any{
				"email":     "ronnie.ernser@email.com",
				"full_name": "Ronnie Ernser",
				"name":      "ernser",
				"type":      0,
			},
		},

		{
			For: BlogServiceTest,
			testValues: map[string]any{
				"user": map[string]any{
					"email":     "user3@gmail.com",
					"full_name": "Peter D. Balana",
					"name":      "username123",
					"type":      pb.User_T_AUTHOR,
				},

				"headline_text": "Mock Blog Title",
				"summary_text":  "This is a mock blog post about various topics.",
				"tags":          []string{"technology", "programming", "web development"},
				"images":        []string{"image1.jpg", "image2.jpg", "image3.jpg"},
				"attachments":   []string{"attachment1.pdf", "attachment2.docx"},
				"refs":          []string{"ref1", "ref2", "ref3"},
			},
		},

		{
			For: UserServiceTest,
			testValues: map[string]any{
				"email":     "sharita.volkman@email.com",
				"full_name": "Sharita Bulkmann",
				"name":      "sharita.volkman",
				"type":      0,
			},
		},
	}
)
