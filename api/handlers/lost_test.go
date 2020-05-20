package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kotyarich/find-your-pet/models"
	"github.com/Kotyarich/find-your-pet/store/db/pg"
	"github.com/brianvoe/gofakeit/v4"
	"github.com/bxcodec/faker"
	"github.com/gavv/httpexpect/v2"
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
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	// lc := pg.NewLostControllerPg(4, db, queryLost)
	standardHandlerData := &HandlerData{
		// LostController: lc,
		FileMaxSize: 10,
		// LostAddingManager: managers.
	}
	server := httptest.NewServer(http.HandlerFunc(standardHandlerData.AddLostHandler))
	e := httpexpect.New(t, server.URL)
	type Form struct {
		TypeId      int     `form:"type_id"`
		AuthorId    int     `form:"vk_id"`
		Sex         string  `form:"sex"`
		Breed       string  `form:"breed"`
		Description string  `form:"description"`
		Latitude    float64 `form:"latitude"`
		Longitude   float64 `form:"longitude"`
		Address     string  `form:"address"`
	}
	tests := []struct {
		name string
		hd   *HandlerData
		form *Form
	}{
		{
			name: "1",
			hd:   standardHandlerData,
			form: &Form{
				TypeId:      gofakeit.Number(1, 3),
				AuthorId:    gofakeit.Number(1, 100),
				Sex:         "m",
				Breed:       gofakeit.AnimalType(),
				Description: "Earum ea quia id ea nulla porro sequi voluptatem. Ut nemo eius non labore eaque. Suscipit numquam.",
				Latitude:    gofakeit.Latitude(),
				Longitude:   gofakeit.Longitude(),
				Address:     gofakeit.Address().Address,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e.POST("/lost").WithMultipart().WithFile("picture", "./img.jpg").
				WithForm(tt.form).Expect().Status(http.StatusOK)
		})
	}
}

func TestHandlerData_RemoveLostHandler(t *testing.T) {
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
			tt.hd.RemoveLostHandler(tt.args.w, tt.args.r)
		})
	}
}
