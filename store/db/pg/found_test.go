package pg

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/Kotyarich/find-your-pet/models"
)

func TestNewFoundControllerPg(t *testing.T) {
	t.Skip()
	type args struct {
		pageCapacity int
		db           *sql.DB
		query        string
	}
	tests := []struct {
		name string
		args args
		want *FoundControllerPg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFoundControllerPg(tt.args.pageCapacity, tt.args.db, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFoundControllerPg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundControllerPg_GetById(t *testing.T) {
	t.Skip()
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		fc      *FoundControllerPg
		args    args
		want    *models.Found
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.GetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoundControllerPg.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoundControllerPg.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundControllerPg_Add(t *testing.T) {
	t.Skip()
	type args struct {
		ctx    context.Context
		params *models.Found
	}
	tests := []struct {
		name    string
		fc      *FoundControllerPg
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.Add(tt.args.ctx, tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoundControllerPg.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FoundControllerPg.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundControllerPg_Search(t *testing.T) {
	t.Skip()
	type args struct {
		ctx    context.Context
		params *models.Found
		query  string
		page   int
	}
	tests := []struct {
		name    string
		fc      *FoundControllerPg
		args    args
		want    []models.Found
		want1   bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.fc.Search(tt.args.ctx, tt.args.params, tt.args.query, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoundControllerPg.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoundControllerPg.Search() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FoundControllerPg.Search() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFoundControllerPg_SearchByType(t *testing.T) {
	t.Skip()
	type args struct {
		ctx    context.Context
		typeId int
	}
	tests := []struct {
		name    string
		fc      *FoundControllerPg
		args    args
		want    []models.Found
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.SearchByType(tt.args.ctx, tt.args.typeId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoundControllerPg.SearchByType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoundControllerPg.SearchByType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundControllerPg_SearchBySex(t *testing.T) {
	t.Skip()
	type args struct {
		ctx context.Context
		sex string
	}
	tests := []struct {
		name    string
		fc      *FoundControllerPg
		args    args
		want    []models.Found
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.SearchBySex(tt.args.ctx, tt.args.sex)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoundControllerPg.SearchBySex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoundControllerPg.SearchBySex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundControllerPg_SearchByBreed(t *testing.T) {
	t.Skip()
	type args struct {
		ctx   context.Context
		breed string
	}
	tests := []struct {
		name    string
		fc      *FoundControllerPg
		args    args
		want    []models.Found
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.SearchByBreed(tt.args.ctx, tt.args.breed)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoundControllerPg.SearchByBreed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoundControllerPg.SearchByBreed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundControllerPg_SearchByTextQuery(t *testing.T) {
	t.Skip()
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		fc      *FoundControllerPg
		args    args
		want    []models.Found
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.SearchByTextQuery(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("FoundControllerPg.SearchByTextQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoundControllerPg.SearchByTextQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundControllerPg_GetPageCapacity(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		fc   *FoundControllerPg
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fc.GetPageCapacity(); got != tt.want {
				t.Errorf("FoundControllerPg.GetPageCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoundControllerPg_GetDbAdapter(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		fc   *FoundControllerPg
		want *sql.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fc.GetDbAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FoundControllerPg.GetDbAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}
