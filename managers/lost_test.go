package managers

import (
	"database/sql"
	"testing"

	"github.com/Kotyarich/find-your-pet/interfaces"
)

func TestLostAddingManager_Remove(t *testing.T) {
	type fields struct {
		db                    *sql.DB
		lostController        interfaces.LostController
		FileController        interfaces.FileController
		baseLostDirectoryPath string
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lam := &LostAddingManager{
				db:                    tt.fields.db,
				lostController:        tt.fields.lostController,
				FileController:        tt.fields.FileController,
				baseLostDirectoryPath: tt.fields.baseLostDirectoryPath,
			}
			if err := lam.Remove(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("LostAddingManager.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
