package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/mocks"
	"github.com/Kotyarich/find-your-pet/models"
	"github.com/brianvoe/gofakeit/v4"
	"github.com/bxcodec/faker"
	"github.com/gavv/httpexpect/v2"
)

func GetStandardLost(count int) []models.Lost {
	var losts []models.Lost
	for i := 0; i < count; i++ {
		lost := models.Lost{}
		faker.FakeData(&lost)
		lost.Id = (i + 1)
		lost.Location = ""
		lost.StatusId = 1
		losts = append(losts, lost)
	}
	return losts
}

func GetStandardFound(count int) []models.Found {
	var founds []models.Found
	for i := 0; i < count; i++ {
		found := models.Found{}
		faker.FakeData(&found)
		found.Id = (i + 1)
		found.Location = ""
		founds = append(founds, found)
	}
	return founds
}

func TestHandlerData_ProfileLostHandler(t *testing.T) {
	t.Parallel()
	standardHandlerData := &HandlerData{
		DebugMode: false,
	}
	server := httptest.NewServer(http.HandlerFunc(standardHandlerData.ProfileLostHandler))
	e := httpexpect.New(t, server.URL)

	type test struct {
		name     string
		hd       *HandlerData
		authorId string
		wantErr  bool
		testType string
	}
	tests := []test{
		{
			name:     "1",
			hd:       standardHandlerData,
			authorId: "13940",
			wantErr:  false,
			testType: "usual",
		},
		{
			name:     "wrong vk_id 1",
			hd:       standardHandlerData,
			authorId: "7857hggvhgvvv",
			wantErr:  true,
			testType: "wrong-id",
		},
	}
	for i := 1; i < 15; i++ {
		test := test{
			name:     strconv.Itoa(i + 1),
			hd:       standardHandlerData,
			authorId: strconv.Itoa(gofakeit.Number(1000, 100000)),
			wantErr:  false,
			testType: "usual",
		}
		tests = append(tests, test)
	}

	for i := 1; i < 15; i++ {
		test := test{
			name:     fmt.Sprintf("wrong vk_id %d", i+1),
			hd:       standardHandlerData,
			authorId: gofakeit.Word(),
			wantErr:  true,
			testType: "wrong-id",
		}
		tests = append(tests, test)
	}

	for i := 1; i < 15; i++ {
		test := test{
			name:     fmt.Sprintf("some-error %d", i+1),
			hd:       standardHandlerData,
			authorId: strconv.Itoa(gofakeit.Number(1000, 100000)),
			wantErr:  true,
			testType: "internal-error",
		}
		tests = append(tests, test)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.testType {
			case "usual":
				losts := GetStandardLost(10)
				ownCount := gofakeit.Number(1, 3)
				authorId, _ := strconv.Atoi(tt.authorId)
				var ownLosts []models.Lost
				for i := 0; i < ownCount; i++ {
					losts[i].AuthorId = authorId
					ownLosts = append(ownLosts, losts[i])
				}
				tt.hd.ProfileController = mocks.NewMockProfileController(losts, []models.Found{}, nil)

				resultLosts := e.GET("/profile/lost").WithQuery("vk_id", tt.authorId).Expect().
					Status(http.StatusOK).JSON().Object().
					ContainsKey("payload").Value("payload").
					Array()
				resultLosts.Equal(ownLosts)
			case "wrong-id":
				losts := []models.Lost{}
				tt.hd.ProfileController = mocks.NewMockProfileController(losts, []models.Found{}, errors.New("Wrong arg"))
				e.GET("/profile/lost").WithQuery("vk_id", tt.authorId).Expect().
					Status(http.StatusBadRequest).NoContent()
			case "internal-error":
				losts := GetStandardLost(10)
				ownCount := gofakeit.Number(1, 3)
				authorId, _ := strconv.Atoi(tt.authorId)
				var ownLosts []models.Lost
				for i := 0; i < ownCount; i++ {
					losts[i].AuthorId = authorId
					ownLosts = append(ownLosts, losts[i])
				}
				tt.hd.ProfileController = mocks.NewMockProfileController(losts, []models.Found{}, errors.New("Internal error"))
				e.GET("/profile/lost").WithQuery("vk_id", tt.authorId).Expect().
					Status(http.StatusInternalServerError).NoContent()
			}

		})
	}
}

func TestHandlerData_ProfileLostOpeningHandler(t *testing.T) {
	t.Parallel()
	standardHandlerData := &HandlerData{
		DebugMode: false,
	}
	server := httptest.NewServer(http.HandlerFunc(standardHandlerData.ProfileLostOpeningHandler))
	e := httpexpect.New(t, server.URL)

	type Form struct {
		LostId   string `form:"lost_id"`
		StatusId string `form:"status_id"`
	}

	type test struct {
		name     string
		hd       *HandlerData
		form     *Form
		testType string
	}
	tests := []test{
		{
			name: "empty id",
			hd:   standardHandlerData,
			form: &Form{
				LostId:   "",
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "wrong-arg",
		},
		{
			name: "empty status",
			hd:   standardHandlerData,
			form: &Form{
				LostId:   strconv.Itoa(gofakeit.Number(1, 20)),
				StatusId: "",
			},
			testType: "wrong-arg",
		},
		{
			name: "wrong special id",
			hd:   standardHandlerData,
			form: &Form{
				LostId:   "454930hkefwkjkwefw",
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "wrong-arg",
		},
		{
			name: "wrong special status",
			hd:   standardHandlerData,
			form: &Form{
				LostId:   strconv.Itoa(gofakeit.Number(1, 20)),
				StatusId: "mfkrel90903333",
			},
			testType: "wrong-arg",
		},
	}
	for i := 1; i <= 20; i++ {
		tst := test{
			name: strconv.Itoa(i),
			hd:   standardHandlerData,
			form: &Form{
				LostId:   strconv.Itoa(i),
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "usual",
		}
		tests = append(tests, tst)

		// For id

		tst = test{
			name: fmt.Sprintf("wrong id %d", i),
			hd:   standardHandlerData,
			form: &Form{
				LostId:   "iuniugtyr",
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "wrong-arg",
		}
		tests = append(tests, tst)

		// For status

		tst = test{
			name: fmt.Sprintf("wrong status %d", i),
			hd:   standardHandlerData,
			form: &Form{
				LostId:   strconv.Itoa(i),
				StatusId: "moibutfrrty",
			},
			testType: "wrong-arg",
		}
		tests = append(tests, tst)

		tst = test{
			name: fmt.Sprintf("internal error %d", i),
			hd:   standardHandlerData,
			form: &Form{
				LostId:   strconv.Itoa(i),
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "internal-error",
		}
		tests = append(tests, tst)

		tst = test{
			name: fmt.Sprintf("not found %d", i),
			hd:   standardHandlerData,
			form: &Form{
				LostId:   strconv.Itoa(i),
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "not-found",
		}
		tests = append(tests, tst)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var httpStatus int
			var losts []models.Lost
			var err error
			switch tt.testType {
			case "usual":
				losts = GetStandardLost(20)
				httpStatus = http.StatusOK
				err = nil
			case "wrong-arg":
				losts = GetStandardLost(1)
				httpStatus = http.StatusBadRequest
				err = nil
			case "internal-error":
				losts = GetStandardLost(5)
				httpStatus = http.StatusInternalServerError
				err = errors.New("some internal error")
			case "not-found":
				losts = GetStandardLost(5)
				httpStatus = http.StatusNotFound
				err = errs.LostNotFound
			}

			tt.hd.ProfileController = mocks.NewMockProfileController(losts, []models.Found{}, err)
			e.PUT("/lost").WithForm(tt.form).Expect().
				Status(httpStatus).NoContent()
		})
	}
}

func TestHandlerData_ProfileFoundHandler(t *testing.T) {
	t.Parallel()
	standardHandlerData := &HandlerData{
		DebugMode: false,
	}
	server := httptest.NewServer(http.HandlerFunc(standardHandlerData.ProfileFoundHandler))
	e := httpexpect.New(t, server.URL)

	type test struct {
		name     string
		hd       *HandlerData
		authorId string
		testType string
	}
	tests := []test{
		{
			name:     "1",
			hd:       standardHandlerData,
			authorId: "13940",
			testType: "usual",
		},
		{
			name:     "wrong vk_id 1",
			hd:       standardHandlerData,
			authorId: "7857hggvhgvvv",
			testType: "wrong-id",
		},
	}
	for i := 1; i < 15; i++ {
		test := test{
			name:     strconv.Itoa(i + 1),
			hd:       standardHandlerData,
			authorId: strconv.Itoa(gofakeit.Number(1000, 100000)),
			testType: "usual",
		}
		tests = append(tests, test)
	}

	for i := 1; i < 15; i++ {
		test := test{
			name:     fmt.Sprintf("wrong vk_id %d", i+1),
			hd:       standardHandlerData,
			authorId: gofakeit.Word(),
			testType: "wrong-id",
		}
		tests = append(tests, test)
	}

	for i := 1; i < 15; i++ {
		test := test{
			name:     fmt.Sprintf("some-error %d", i+1),
			hd:       standardHandlerData,
			authorId: strconv.Itoa(gofakeit.Number(1000, 100000)),
			testType: "internal-error",
		}
		tests = append(tests, test)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.testType {
			case "usual":
				founds := GetStandardFound(10)
				ownCount := gofakeit.Number(1, 3)
				authorId, _ := strconv.Atoi(tt.authorId)
				var ownfounds []models.Found
				for i := 0; i < ownCount; i++ {
					founds[i].AuthorId = authorId
					ownfounds = append(ownfounds, founds[i])
				}
				tt.hd.ProfileController = mocks.NewMockProfileController([]models.Lost{}, founds, nil)

				resultfounds := e.GET("/profile/found").WithQuery("vk_id", tt.authorId).Expect().
					Status(http.StatusOK).JSON().Object().
					ContainsKey("payload").Value("payload").
					Array()
				resultfounds.Equal(ownfounds)
			case "wrong-id":
				founds := []models.Found{}
				tt.hd.ProfileController = mocks.NewMockProfileController([]models.Lost{}, founds, errors.New("some error"))
				e.GET("/profile/found").WithQuery("vk_id", tt.authorId).Expect().
					Status(http.StatusBadRequest).NoContent()
			case "internal-error":
				founds := GetStandardFound(10)
				ownCount := gofakeit.Number(1, 3)
				authorId, _ := strconv.Atoi(tt.authorId)
				var ownFounds []models.Found
				for i := 0; i < ownCount; i++ {
					founds[i].AuthorId = authorId
					ownFounds = append(ownFounds, founds[i])
				}
				tt.hd.ProfileController = mocks.NewMockProfileController([]models.Lost{}, founds, errors.New("some error"))
				e.GET("/profile/found").WithQuery("vk_id", tt.authorId).Expect().
					Status(http.StatusInternalServerError).NoContent()
			}

		})
	}
}

func TestHandlerData_ProfileFoundOpeningHandler(t *testing.T) {
	t.Parallel()
	standardHandlerData := &HandlerData{
		DebugMode: false,
	}
	server := httptest.NewServer(http.HandlerFunc(standardHandlerData.ProfileFoundOpeningHandler))
	e := httpexpect.New(t, server.URL)

	type Form struct {
		FoundId  string `form:"found_id"`
		StatusId string `form:"status_id"`
	}

	type test struct {
		name     string
		hd       *HandlerData
		form     *Form
		testType string
	}
	tests := []test{
		{
			name: "empty id",
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  "",
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "wrong-arg",
		},
		{
			name: "empty status",
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  strconv.Itoa(gofakeit.Number(1, 20)),
				StatusId: "",
			},
			testType: "wrong-arg",
		},
		{
			name: "wrong special id",
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  "454930hkefwkjkwefw",
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "wrong-arg",
		},
		{
			name: "wrong special status",
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  strconv.Itoa(gofakeit.Number(1, 20)),
				StatusId: "mfkrel90903333",
			},
			testType: "wrong-arg",
		},
	}
	for i := 1; i <= 20; i++ {
		tst := test{
			name: strconv.Itoa(i),
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  strconv.Itoa(i),
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "usual",
		}
		tests = append(tests, tst)

		// For id

		tst = test{
			name: fmt.Sprintf("wrong id %d", i),
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  "iuniugtyr",
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "wrong-arg",
		}
		tests = append(tests, tst)

		// For status

		tst = test{
			name: fmt.Sprintf("wrong status %d", i),
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  strconv.Itoa(i),
				StatusId: "moibutfrrty",
			},
			testType: "wrong-arg",
		}
		tests = append(tests, tst)

		tst = test{
			name: fmt.Sprintf("internal error %d", i),
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  strconv.Itoa(i),
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "internal-error",
		}
		tests = append(tests, tst)

		tst = test{
			name: fmt.Sprintf("not found %d", i),
			hd:   standardHandlerData,
			form: &Form{
				FoundId:  strconv.Itoa(i),
				StatusId: strconv.Itoa(gofakeit.Number(2, 5)),
			},
			testType: "not-found",
		}
		tests = append(tests, tst)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var httpStatus int
			var founds []models.Found
			var err error
			switch tt.testType {
			case "usual":
				founds = GetStandardFound(20)
				httpStatus = http.StatusOK
				err = nil
			case "wrong-arg":
				founds = GetStandardFound(1)
				httpStatus = http.StatusBadRequest
				err = nil
			case "internal-error":
				founds = GetStandardFound(5)
				httpStatus = http.StatusInternalServerError
				err = errors.New("some internal error")
			case "not-found":
				founds = GetStandardFound(5)
				httpStatus = http.StatusNotFound
				err = errs.TheFoundNotFound
			}

			tt.hd.ProfileController = mocks.NewMockProfileController([]models.Lost{}, founds, err)
			e.PUT("/found").WithForm(tt.form).Expect().
				Status(httpStatus).NoContent()
		})
	}

}
