package pg

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kotyarich/find-your-pet/mocks"
	"github.com/Kotyarich/find-your-pet/models"
)

func TestNewLostControllerPg(t *testing.T) {
	type args struct {
		itemsPerPage int
		db           *sql.DB
	}
	mockDb, _, _ := sqlmock.New()
	tests := []struct {
		name string
		args args
	}{
		{
			"First with nil",
			args{
				itemsPerPage: 5,
				db:           nil,
			},
		},
		{
			"Second with nil",
			args{
				itemsPerPage: 10,
				db:           nil,
			},
		},
		{
			"Third with mock",
			args{
				itemsPerPage: 999999,
				db:           mockDb,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLostControllerPg(tt.args.itemsPerPage, tt.args.db)
			currentItems := got.GetItemsPerPageCount()
			if currentItems != tt.args.itemsPerPage {
				t.Errorf("items per page must be %v, but got %v",
					tt.args.itemsPerPage, currentItems)
			}
			currentDb := got.GetDbAdapter()
			if currentDb != tt.args.db {
				t.Errorf("db must be %v, but got %v",
					tt.args.db, currentDb)
			}
		})
	}
}

func TestLostControllerPg_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	type fields struct {
		itemsPerPage int
		db           *sql.DB
	}
	type args struct {
		id int
	}
	realLost := mocks.Generate()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Lost
		wantErr bool
	}{
		{
			"First",
			fields{
				itemsPerPage: 5,
				db:           db,
			},
			args{
				id: 1,
			},
			&realLost[0],
			false,
		},
		{
			"Third",
			fields{
				itemsPerPage: 5,
				db:           db,
			},
			args{
				id: 3,
			},
			&realLost[2],
			false,
		},
		{
			"Last",
			fields{
				itemsPerPage: 5,
				db:           db,
			},
			args{
				id: 5,
			},
			&realLost[4],
			false,
		},
		{
			"Not exists",
			fields{
				itemsPerPage: 5,
				db:           db,
			},
			args{
				id: 45434543,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LostControllerPg{
				itemsPerPage: tt.fields.itemsPerPage,
				db:           tt.fields.db,
			}
			rows := sqlmock.NewRows([]string{
				"id",
				"type_id",
				"author_id",
				"sex",
				"breed",
				"description",
				"status_id",
				"date",
				"place",
			})
			id := tt.args.id
			expectQuery := mock.ExpectQuery("SELECT (.+) FROM lost WHERE id = \\$1").
				WithArgs(id)
			if id >= 1 && id <= 5 {
				lost := realLost[id-1]
				rows.AddRow(lost.Id, lost.TypeId, lost.AuthorId, lost.Sex,
					lost.Breed, lost.Description, lost.StatusId,
					lost.Date, lost.Place)
				expectQuery.WillReturnRows(rows)
			} else {
				expectQuery.WillReturnError(errors.New("Doesn't exist"))
			}
			got, err := lc.GetById(id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostControllerPg.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostControllerPg.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_Add(t *testing.T) {
	type fields struct {
		itemsPerPage int
		db           *sql.DB
	}
	type args struct {
		params *models.Lost
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LostControllerPg{
				itemsPerPage: tt.fields.itemsPerPage,
				db:           tt.fields.db,
			}
			got, err := lc.Add(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostControllerPg.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LostControllerPg.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_Search(t *testing.T) {
	type fields struct {
		itemsPerPage int
		db           *sql.DB
	}
	type args struct {
		params *models.Lost
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Lost
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LostControllerPg{
				itemsPerPage: tt.fields.itemsPerPage,
				db:           tt.fields.db,
			}
			got, err := lc.Search(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostControllerPg.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostControllerPg.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_GetItemsPerPageCount(t *testing.T) {
	type fields struct {
		itemsPerPage int
		db           *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LostControllerPg{
				itemsPerPage: tt.fields.itemsPerPage,
				db:           tt.fields.db,
			}
			if got := lc.GetItemsPerPageCount(); got != tt.want {
				t.Errorf("LostControllerPg.GetItemsPerPageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

// for _, lost := range realLost {
// 		rows.AddRow(lost.Id, lost.TypeId, lost.AuthorId, lost.Sex,
// 			lost.Breed, lost.Description, lost.StatusId,
// 			lost.Date, lost.Place)
// 	}
