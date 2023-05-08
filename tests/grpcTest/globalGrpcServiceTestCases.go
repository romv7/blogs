package grpcTest

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
			For:        BlogServiceTest,
			testValues: map[string]any{},
		},

		{
			For:        BlogServiceTest,
			testValues: map[string]any{},
		},

		{
			For:        BlogServiceTest,
			testValues: map[string]any{},
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
			For: UserServiceTest,
			testValues: map[string]any{
				"email":     "dante@dmcheadquarters.com",
				"full_name": "Dante the Demon Slayer",
				"name":      "dantethejester",
				"type":      0,
			},
		},

		{
			For:        BlogServiceTest,
			testValues: map[string]any{},
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
			For: UserServiceTest,
			testValues: map[string]any{
				"email":     "ronnie.ernser@email.com",
				"full_name": "Ronnie Ernser",
				"name":      "ernser",
				"type":      0,
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
