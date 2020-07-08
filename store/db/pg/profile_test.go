package pg

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/Kotyarich/find-your-pet/models"
)

func TestNewProfileControllerPg(t *testing.T) {
	t.Skip()
	type args struct {
		pages      int
		db         *sql.DB
		queryLost  string
		queryFound string
	}
	tests := []struct {
		name string
		args args
		want *ProfileControllerPg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProfileControllerPg(tt.args.pages, tt.args.db, tt.args.queryLost, tt.args.queryFound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProfileControllerPg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileControllerPg_GetLost(t *testing.T) {
	t.Skip()
	type args struct {
		ctx    context.Context
		userId int
	}
	tests := []struct {
		name    string
		pc      *ProfileControllerPg
		args    args
		want    []models.Lost
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pc.GetLost(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileControllerPg.GetLost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileControllerPg.GetLost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileControllerPg_SetLostOpening(t *testing.T) {
	t.Skip()
	type args struct {
		ctx      context.Context
		lostId   int
		statusId int
	}
	tests := []struct {
		name    string
		pc      *ProfileControllerPg
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pc.SetLostOpening(tt.args.ctx, tt.args.lostId, tt.args.statusId); (err != nil) != tt.wantErr {
				t.Errorf("ProfileControllerPg.SetLostOpening() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfileControllerPg_GetFound(t *testing.T) {
	t.Skip()
	type args struct {
		ctx    context.Context
		userId int
	}
	tests := []struct {
		name    string
		pc      *ProfileControllerPg
		args    args
		want    []models.Found
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.pc.GetFound(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileControllerPg.GetFound() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileControllerPg.GetFound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProfileControllerPg_SetFoundOpening(t *testing.T) {
	t.Skip()
	type args struct {
		ctx      context.Context
		foundId  int
		statusId int
	}
	tests := []struct {
		name    string
		pc      *ProfileControllerPg
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pc.SetFoundOpening(tt.args.ctx, tt.args.foundId, tt.args.statusId); (err != nil) != tt.wantErr {
				t.Errorf("ProfileControllerPg.SetFoundOpening() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProfileControllerPg_GetItemsPerPageCount(t *testing.T) {
	t.Skip()
	tests := []struct {
		name string
		pc   *ProfileControllerPg
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pc.GetItemsPerPageCount(); got != tt.want {
				t.Errorf("ProfileControllerPg.GetItemsPerPageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
