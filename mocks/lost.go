package mocks

import "github.com/Kotyarich/find-your-pet/models"

func Generate() []models.Lost {
	return []models.Lost{
		{
			Id:          1,
			TypeId:      1,
			Sex:         "m",
			Breed:       "orchid card",
			Description: "Eligendi et rerum voluptatem magnam accusantium repudiandae.",
			StatusId:    1,
			Date:        "Thu Sep 12 2019 12:03:22 GMT+0300 (Moscow Standard Time)",
			Place:       "03262 Vandervort Villages",
		},
		{
			Id:     2,
			TypeId: 1,
			Sex:    "f",
			Breed:  "orchid card",
			Description: "Est consequuntur molestiae cupiditate vel.\n" +
				"Eligendi rem esse voluptatem.\n" +
				"Aut accusantium consequatur et.\n" +
				"Velit iure deserunt rerum.\n" +
				"Voluptas voluptas recusandae quisquam.\n",
			StatusId: 1,
			Date:     "Thu Sep 12 2019 12:03:22 GMT+0300 (Moscow Standard Time)",
			Place:    "03650 Vandervort Villages",
		},
		{
			Id:          3,
			TypeId:      2,
			Sex:         "m",
			Breed:       "Guinea-Bissau",
			Description: "None",
			StatusId:    1,
			Date:        "Thu Sep 12 2019 12:03:22 GMT+0300 (Moscow Standard Time)",
			Place:       "5364 Watson Courts",
		},
		{
			Id:     4,
			TypeId: 2,
			Sex:    "f",
			Breed:  "Guinea-Bissau",
			Description: "Ab molestiae aut accusantium et vel.\n" +
				"Rerum blanditiis unde nam vitae iusto.\n" +
				"Tempora dolorum exercitationem doloremque numquam qui officiis debitis reiciendis.\n" +
				"Iusto fugit earum pariatur totam.",
			StatusId: 1,
			Date:     "Thu Sep 12 2019 12:03:22 GMT+0300 (Moscow Standard Time)",
			Place:    "2313 Raynor Brooks",
		},
		{
			Id:     5,
			TypeId: 3,
			Sex:    "m",
			Breed:  "initiatives ROI",
			Description: "Aut optio voluptate impedit sit quisquam. Voluptatibus cum libero.\n" +
				"Ducimus odio vel soluta et dolores. Enim nostrum itaque.\n" +
				"Dolorum asperiores doloribus dolorem ut magnam repudiandae praesentium.\n" +
				"Doloribus labore cupiditate quaerat voluptas.",
			StatusId: 1,
			Date:     "Thu Sep 12 2019 12:03:22 GMT+0300 (Moscow Standard Time)",
			Place:    "03262 Vandervort Cities",
		},
	}
}
