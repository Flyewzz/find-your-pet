package pg

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/Kotyarich/find-your-pet/models"
)

func TestNewFileControllerPg(t *testing.T) {
	t.Skip()
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *FileControllerPg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileControllerPg(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileControllerPg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileControllerPg_GetById(t *testing.T) {
	t.Skip()
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fc      *FileControllerPg
		args    args
		want    *models.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.GetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileControllerPg.GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileControllerPg.GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileControllerPg_AddToLost(t *testing.T) {
	t.Skip()
	type args struct {
		ctx    context.Context
		file   *models.File
		lostId int
	}
	tests := []struct {
		name    string
		fc      *FileControllerPg
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.AddToLost(tt.args.ctx, tt.args.file, tt.args.lostId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileControllerPg.AddToLost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileControllerPg.AddToLost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileControllerPg_AddToFound(t *testing.T) {
	t.Skip()
	type args struct {
		ctx     context.Context
		file    *models.File
		foundId int
	}
	tests := []struct {
		name    string
		fc      *FileControllerPg
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fc.AddToFound(tt.args.ctx, tt.args.file, tt.args.foundId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileControllerPg.AddToFound() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileControllerPg.AddToFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
