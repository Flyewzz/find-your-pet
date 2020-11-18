package pg

import (
	"context"
	"database/sql"
	"reflect"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Kotyarich/find-your-pet/models"
	"github.com/bxcodec/faker"
)

var (
	queryLost = "SELECT id, type_id, " +
		"vk_id, sex, " +
		"breed, description, status_id, " +
		"date, st_x(location) as latitude, " +
		"st_y(location) as longitude, picture_id, address FROM lost "
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

func TestNewLostControllerPg(t *testing.T) {
	t.Skip()
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	type args struct {
		pageCapacity int
		db           *sql.DB
		id           string
	}
	tests := []struct {
		name string
		args args
		want *LostControllerPg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := NewLostControllerPg(tt.args.pageCapacity, tt.args.db, tt.args.query); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("NewLostControllerPg() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestLostControllerPg_GetById(t *testing.T) {
	t.Skip()
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	type args struct {
		ctx context.Context
		id  int
	}
	type test struct {
		name string
		lc   *LostControllerPg
		args args
		// want    *models.Lost
		err      error
		testType string
	}
	tests := []test{}

	for i := 1; i <= 15; i++ {
		tst := test{
			name: strconv.Itoa(i),
			lc:   NewLostControllerPg(4, db, queryLost),
			args: args{
				ctx: context.Background(),
				id:  1,
			},
			err: nil,
		}
		tests = append(tests, tst)
	}

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		i := tt.args.id - 1
	// 		rows := sqlmock.NewRows([]string{"id", "type_id",
	// 			"vk_id", "sex", "breed", "description", "status_id",
	// 			"date", "latitude", "longitude", "picture_id", "address"}).AddRow(losts[i].Id, losts[i].TypeId, losts[i].AuthorId,
	// 			losts[i].Sex, losts[i].Breed, losts[i].Description,
	// 			losts[i].StatusId, losts[i].Date,
	// 			losts[i].Latitude, losts[i].Longitude, losts[i].PictureId, losts[i].Address)
	// 		mock.ExpectQuery(`.*`).WithArgs(tt.args.id).WillReturnRows(rows)
	// 		got, err := tt.lc.GetById(tt.args.ctx, tt.args.id)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("LostControllerPg.GetById() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if err := mock.ExpectationsWereMet(); err != nil {
	// 			t.Errorf("there were unfulfilled expectations: %s", err)
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("LostControllerPg.GetById() = %v, \nwant %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestLostControllerPg_Add(t *testing.T) {
	t.Skip()
	type args struct {
		ctx    context.Context
		params *models.Lost
	}
	tests := []struct {
		name    string
		lc      *LostControllerPg
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.lc.Add(tt.args.ctx, tt.args.params)
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
	t.Skip()
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()

	// GetLostRandom := func(count int) []models.Lost {
	// 	var losts []models.Lost
	// 	for i := 0; i < count; i++ {
	// 		lost := models.Lost{}
	// 		faker.FakeData(&lost)
	// 		lost.Id = (i + 1)
	// 		lost.Location = ""
	// 		losts = append(losts, lost)
	// 	}
	// 	return losts
	// }
	type args struct {
		ctx    context.Context
		params *models.Lost
		query  string
		page   int
	}
	tests := []struct {
		name    string
		lc      *LostControllerPg
		args    args
		want    []models.Lost
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.lc.Search(tt.args.ctx, tt.args.params, tt.args.query, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostControllerPg.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostControllerPg.Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LostControllerPg.Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLostControllerPg_SearchByType(t *testing.T) {
	t.Skip()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	type args struct {
		ctx    context.Context
		typeId int
	}
	tests := []struct {
		name    string
		lc      *LostControllerPg
		args    args
		want    []models.Lost
		wantErr bool
	}{
		{
			name: "1",
			lc:   NewLostControllerPg(4, db, queryLost),
			args: args{
				ctx:    context.Background(),
				typeId: 1,
			},
			want:    []models.Lost{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx, _ := db.Begin()
			tt.args.ctx = context.WithValue(tt.args.ctx, "tx", tx)
			mock.ExpectQuery("(.+)").WillReturnRows()
			got, err := tt.lc.SearchByType(tt.args.ctx, tt.args.typeId)
			if (err != nil) != tt.wantErr {
				tx.Rollback()
				t.Errorf("LostControllerPg.SearchByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tx.Commit()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostControllerPg.SearchByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_SearchBySex(t *testing.T) {
	t.Skip()
	type args struct {
		ctx context.Context
		sex string
	}
	tests := []struct {
		name    string
		lc      *LostControllerPg
		args    args
		want    []models.Lost
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.lc.SearchBySex(tt.args.ctx, tt.args.sex)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostControllerPg.SearchBySex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostControllerPg.SearchBySex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_SearchByBreed(t *testing.T) {
	t.Skip()
	type args struct {
		ctx   context.Context
		breed string
	}
	tests := []struct {
		name    string
		lc      *LostControllerPg
		args    args
		want    []models.Lost
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.lc.SearchByBreed(tt.args.ctx, tt.args.breed)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostControllerPg.SearchByBreed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostControllerPg.SearchByBreed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_SearchByTextQuery(t *testing.T) {
	t.Skip()
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		lc      *LostControllerPg
		args    args
		want    []models.Lost
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.lc.SearchByTextQuery(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostControllerPg.SearchByTextQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostControllerPg.SearchByTextQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_GetPageCapacity(t *testing.T) {
	t.Skip()
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	tests := []struct {
		name string
		lc   *LostControllerPg
		want int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			lc:   NewLostControllerPg(1, nil, ""),
			want: 1,
		},
		{
			name: "4",
			lc:   NewLostControllerPg(4, nil, ""),
			want: 4,
		},
		{
			name: "Mock db",
			lc:   NewLostControllerPg(12, db, ""),
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lc.GetPageCapacity(); got != tt.want {
				t.Errorf("LostControllerPg.GetPageCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_GetDbAdapter(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		lc   *LostControllerPg
		want *sql.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lc.GetDbAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LostControllerPg.GetDbAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLostControllerPg_RemoveById(t *testing.T) {
	t.Skip()
	type fields struct {
		pageCapacity        int
		db                  *sql.DB
		searchRequiredQuery string
	}
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LostControllerPg{
				pageCapacity:        tt.fields.pageCapacity,
				db:                  tt.fields.db,
				searchRequiredQuery: tt.fields.searchRequiredQuery,
			}
			got, err := lc.RemoveById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LostControllerPg.RemoveById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LostControllerPg.RemoveById() = %v, want %v", got, tt.want)
			}
		})
	}

	// dirPath := fmt.Sprintf("./lost/%d", dbId)
	// 			_, err := fs.Stat(dirPath)
	// 			if err != nil {
	// 				t.Fatalf("Directory %s does not exist!", dirPath)
	// 			}
	// 			fileInfo, err := afs.ReadDir(dirPath)
	// 			if err != nil {
	// 				t.Errorf("Dir error: %s", dirPath)
	// 			}
	// 			for _, file := range fileInfo {
	// 				fmt.Println(file.Name())
	// 			}
	// 			if len(fileInfo) == 0 {
	// 				t.Errorf("Picture was not added")
	// 			}
}
