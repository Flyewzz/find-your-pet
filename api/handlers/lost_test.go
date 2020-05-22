package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kotyarich/find-your-pet/managers"
	"github.com/Kotyarich/find-your-pet/mocks"
	"github.com/Kotyarich/find-your-pet/models"
	"github.com/Kotyarich/find-your-pet/store/db/pg"
	"github.com/brianvoe/gofakeit/v4"
	"github.com/bxcodec/faker"
	"github.com/gavv/httpexpect/v2"
	"github.com/spf13/afero"
)

var (
	queryLost = "SELECT id, type_id, " +
		"vk_id, sex, " +
		"breed, description, status_id, " +
		"date, st_x(location) as latitude, " +
		"st_y(location) as longitude, picture_id, address FROM lost "
)

func TestHandlerData_LostHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		hd   *HandlerData
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hd.LostHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestHandlerData_LostByIdGetHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	var losts []*models.Lost
	for i := 0; i < 10; i++ {
		lost := &models.Lost{}
		faker.FakeData(lost)
		lost.Id = (i + 1)
		lost.Location = ""
		losts = append(losts, lost)
	}
	type args struct {
		id int
	}
	type test struct {
		name string
		hd   *HandlerData
		args args
		want *models.Lost
	}

	standardHandlerData := &HandlerData{
		LostController: pg.NewLostControllerPg(4, db, queryLost),
	}

	tests := []test{
		{
			name: "id 1",
			hd:   standardHandlerData,
			args: args{
				id: 1,
			},
		},
		{
			name: "id 4",
			hd:   standardHandlerData,
			args: args{
				id: 4,
			},
		},
		{
			name: "id 10",
			hd:   standardHandlerData,
			args: args{
				id: 10,
			},
		},
	}
	fields := []string{"id", "type_id",
		"vk_id", "sex", "breed", "description", "status_id",
		"date", "latitude", "longitude", "picture_id", "address"}

	server := httptest.NewServer(http.HandlerFunc(standardHandlerData.LostByIdGetHandler))
	e := httpexpect.New(t, server.URL)
	defer server.Close()
	for _, tt := range tests {
		i := tt.args.id - 1
		rows := sqlmock.NewRows(fields).AddRow(losts[i].Id, losts[i].TypeId, losts[i].AuthorId,
			losts[i].Sex, losts[i].Breed, losts[i].Description,
			losts[i].StatusId, losts[i].Date,
			losts[i].Latitude, losts[i].Longitude, losts[i].PictureId, losts[i].Address)
		mock.ExpectQuery(`.*`).WithArgs(tt.args.id).WillReturnRows(rows)

		e.GET("/lost").WithQuery("id", tt.args.id).
			Expect().
			Status(http.StatusOK).JSON().Object().
			ContainsKey("id").ValueEqual("id", tt.args.id)
	}

	// Wrong id
	wrongIds := []string{
		"wveboifwde53",
		"fdf",
		"5f1",
		"",
	}
	for _, id := range wrongIds {
		e.GET("/lost").WithQuery("id", id).
			Expect().Status(http.StatusBadRequest)
	}

	//
}

func TestHandlerData_AddLostHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	lc := pg.NewLostControllerPg(4, db, queryLost)
	fc := mocks.NewFileController()
	standardHandlerData := &HandlerData{
		LostController:    lc,
		FileMaxSize:       10,
		LostAddingManager: managers.NewLostAddingManager(db, lc, fc, ".img/lost"),
	}
	server := httptest.NewServer(http.HandlerFunc(standardHandlerData.AddLostHandler))
	e := httpexpect.New(t, server.URL)
	type Form struct {
		TypeId      string `form:"type_id"`
		AuthorId    string `form:"vk_id"`
		Sex         string `form:"sex"`
		Breed       string `form:"breed"`
		Description string `form:"description"`
		Latitude    string `form:"latitude"`
		Longitude   string `form:"longitude"`
		Address     string `form:"address"`
	}
	GetStandardForm := func() *Form {
		return &Form{
			TypeId:      strconv.Itoa(gofakeit.Number(1, 3)),
			AuthorId:    strconv.Itoa(gofakeit.Number(1, 100)),
			Sex:         "m",
			Breed:       gofakeit.AnimalType(),
			Description: "Earum ea quia id ea nulla porro sequi voluptatem. Ut nemo eius non labore eaque. Suscipit numquam.",
			Latitude:    fmt.Sprintf("%f", gofakeit.Latitude()),
			Longitude:   fmt.Sprintf("%f", gofakeit.Longitude()),
			Address:     gofakeit.Address().Address,
		}
	}
	tests := []struct {
		name     string
		hd       *HandlerData
		form     *Form
		testType string
	}{
		{
			name:     "1",
			hd:       standardHandlerData,
			form:     GetStandardForm(),
			testType: "usual",
		},
		{
			name:     "2",
			hd:       standardHandlerData,
			form:     GetStandardForm(),
			testType: "usual",
		},
		{
			name:     "3",
			hd:       standardHandlerData,
			form:     GetStandardForm(),
			testType: "usual",
		},
		// wrong attrs
		{
			name: "wrong type_id",
			hd:   standardHandlerData,
			form: &Form{
				TypeId:      "wdewrgtnrt",
				AuthorId:    strconv.Itoa(gofakeit.Number(1, 100)),
				Sex:         "m",
				Breed:       gofakeit.AnimalType(),
				Description: "Earum ea quia id ea nulla porro sequi voluptatem. Ut nemo eius non labore eaque. Suscipit numquam.",
				Latitude:    fmt.Sprintf("%f", gofakeit.Latitude()),
				Longitude:   fmt.Sprintf("%f", gofakeit.Longitude()),
				Address:     gofakeit.Address().Address,
			},
			testType: "wrong",
		},
		{
			name: "wrong vk_id",
			hd:   standardHandlerData,
			form: &Form{
				TypeId:      strconv.Itoa(gofakeit.Number(1, 3)),
				AuthorId:    "idwrong",
				Sex:         "m",
				Breed:       gofakeit.AnimalType(),
				Description: "",
				Latitude:    fmt.Sprintf("%f", gofakeit.Latitude()),
				Longitude:   fmt.Sprintf("%f", gofakeit.Longitude()),
				Address:     gofakeit.Address().Address,
			},
			testType: "wrong",
		},
		{
			name: "wrong sex",
			hd:   standardHandlerData,
			form: &Form{
				TypeId:      strconv.Itoa(gofakeit.Number(1, 3)),
				AuthorId:    strconv.Itoa(gofakeit.Number(1, 100)),
				Sex:         "MkfldkKk",
				Breed:       gofakeit.AnimalType(),
				Description: "Earum ea quia id ea nulla porro sequi voluptatem. Ut nemo eius non labore eaque. Suscipit numquam.",
				Latitude:    fmt.Sprintf("%f", gofakeit.Latitude()),
				Longitude:   fmt.Sprintf("%f", gofakeit.Longitude()),
				Address:     gofakeit.Address().Address,
			},
			testType: "wrong",
		},
		{
			name: "wrong latitude",
			hd:   standardHandlerData,
			form: &Form{
				TypeId:      strconv.Itoa(gofakeit.Number(1, 3)),
				AuthorId:    strconv.Itoa(gofakeit.Number(1, 100)),
				Sex:         "m",
				Breed:       gofakeit.AnimalType(),
				Description: "Earum ea quia id ea nulla porro sequi voluptatem. Ut nemo eius non labore eaque. Suscipit numquam.",
				Latitude:    "fghew",
				Longitude:   fmt.Sprintf("%f", gofakeit.Longitude()),
				Address:     gofakeit.Address().Address,
			},
			testType: "wrong",
		},
		{
			name: "wrong longitude",
			hd:   standardHandlerData,
			form: &Form{
				TypeId:      strconv.Itoa(gofakeit.Number(1, 3)),
				AuthorId:    strconv.Itoa(gofakeit.Number(1, 100)),
				Sex:         "m",
				Breed:       gofakeit.AnimalType(),
				Description: "Earum ea quia id ea nulla porro sequi voluptatem. Ut nemo eius non labore eaque. Suscipit numquam.",
				Latitude:    fmt.Sprintf("%f", gofakeit.Latitude()),
				Longitude:   "-0aefgrt3w",
				Address:     gofakeit.Address().Address,
			},
			testType: "wrong",
		},
		{
			name:     "file error",
			hd:       standardHandlerData,
			form:     GetStandardForm(),
			testType: "errfile",
		},
		{
			name:     "no multipart",
			hd:       standardHandlerData,
			form:     GetStandardForm(),
			testType: "no-multipart",
		},
		{
			name:     "large file",
			hd:       standardHandlerData,
			form:     GetStandardForm(),
			testType: "large-file",
		},
		// {
		// 	name:     "fs error",
		// 	hd:       standardHandlerData,
		// 	form:     GetStandardForm(),
		// 	testType: "fs-error",
		// },
	}
	dbId := 1
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			afs := &afero.Afero{Fs: fs}

			tt.hd.FileMaxSize = 10

			standardHandlerData.FileStoreController = mocks.NewFileStoreController(
				"./lost",
				"./found",
				fs,
			)
			switch tt.testType {
			case "usual":
				mock.ExpectBegin()
				mock.ExpectQuery("(.+)").WillReturnRows(
					sqlmock.NewRows([]string{"id"}).AddRow(dbId),
				)
				mock.ExpectCommit()
				e.POST("/lost").WithMultipart().WithFile("picture", "./test_img.jpg").
					WithForm(tt.form).Expect().Status(http.StatusOK)
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Fatalf("there were unfulfilled expectations: %s", err)
				}
				dirPath := fmt.Sprintf("./lost/%d", dbId)
				_, err := fs.Stat(dirPath)
				if err != nil {
					t.Fatalf("Directory %s does not exist!", dirPath)
				}
				fileInfo, err := afs.ReadDir(dirPath)
				if err != nil {
					t.Errorf("Dir error: %s", dirPath)
				}
				for _, file := range fileInfo {
					fmt.Println(file.Name())
				}
				if len(fileInfo) == 0 {
					t.Errorf("Picture was not added")
				}
			case "wrong":
				e.POST("/lost").WithMultipart().
					WithForm(tt.form).Expect().Status(http.StatusBadRequest)
			case "errfile":
				e.POST("/lost").WithMultipart().
					WithForm(tt.form).WithFormField("picture", "ijkhgf").Expect().Status(http.StatusInternalServerError)
			case "no-multipart":
				e.POST("/lost").WithForm(tt.form).
					Expect().Status(http.StatusBadRequest)
			case "large-file":
				tt.hd.FileMaxSize = 0 // Max size is 0 KB (the picture has 10 KB)
				e.POST("/lost").WithMultipart().WithFile("picture", "./test_img.jpg").
					WithForm(tt.form).Expect().Status(http.StatusBadRequest)
				// case "fs-error":
				// 	afs.Mkdir(fmt.Sprintf("./lost/%d", dbId), 0000)
				// 	mock.ExpectBegin()
				// 	mock.ExpectQuery("(.+)").WillReturnRows(
				// 		sqlmock.NewRows([]string{"id"}).AddRow(dbId),
				// 	)
				// 	mock.ExpectRollback()
				// 	e.POST("/lost").WithMultipart().WithFile("picture", "./test_img.jpg").
				// 		WithForm(tt.form).Expect().Status(http.StatusInternalServerError)
				// 	if err := mock.ExpectationsWereMet(); err != nil {
				// 		t.Fatalf("there were unfulfilled expectations: %s", err)
				// 	}
				// }
			}
			dbId++
		})
	}

}

func TestHandlerData_RemoveLostHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	lc := pg.NewLostControllerPg(4, db, queryLost)
	fc := mocks.NewFileController()
	standardHandlerData := &HandlerData{
		LostController:    lc,
		FileMaxSize:       10,
		LostAddingManager: managers.NewLostAddingManager(db, lc, fc, ".img/lost"),
	}
	server := httptest.NewServer(http.HandlerFunc(standardHandlerData.RemoveLostHandler))
	e := httpexpect.New(t, server.URL)

	tests := []struct {
		name     string
		hd       *HandlerData
		testType string
	}{
		{
			name:     "1",
			hd:       standardHandlerData,
			testType: "usual",
		},
	}

	dbId := 1
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			afs := &afero.Afero{Fs: fs}

			tt.hd.FileMaxSize = 10

			standardHandlerData.FileStoreController = mocks.NewFileStoreController(
				"./lost",
				"./found",
				fs,
			)
			switch tt.testType {
			case "usual":
				mock.ExpectBegin()
				mock.ExpectQuery("(.+)").WithArgs(dbId).WillReturnRows(
					sqlmock.NewRows([]string{"picture_id"}).AddRow(42),
				)
				mock.ExpectExec("(.+)").WithArgs(dbId)
				mock.ExpectExec("(.+)").WithArgs(42)
				mock.ExpectCommit()
				e.DELETE("/lost").WithQuery("id", dbId).Expect().Status(http.StatusOK)
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Fatalf("there were unfulfilled expectations: %s", err)
				}
				dirPath := fmt.Sprintf("./lost/%d", dbId)
				_, err := fs.Stat(dirPath)
				if err != nil {
					t.Fatalf("Directory %s does not exist!", dirPath)
				}
				fileInfo, err := afs.ReadDir(dirPath)
				if err != nil {
					t.Errorf("Dir error: %s", dirPath)
				}
				for _, file := range fileInfo {
					fmt.Println(file.Name())
				}
				if len(fileInfo) == 0 {
					t.Errorf("Picture was not added")
				}
			}
			dbId++
		})
	}

}
