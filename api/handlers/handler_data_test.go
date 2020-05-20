package handlers

import (
	"reflect"
	"testing"

	"github.com/Kotyarich/find-your-pet/interfaces"
	"github.com/Kotyarich/find-your-pet/managers"
)

func TestNewHandlerData(t *testing.T) {
	type args struct {
		lc      interfaces.LostController
		fc      interfaces.FileController
		lam     *managers.LostAddingManager
		fnd     interfaces.FoundController
		fam     *managers.FoundAddingManager
		pc      interfaces.ProfileController
		bc      interfaces.BreedClassifier
		isDebug bool
	}
	tests := []struct {
		name string
		args args
		want *HandlerData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandlerData(tt.args.lc, tt.args.fc, tt.args.lam, tt.args.fnd, tt.args.fam, tt.args.pc, tt.args.bc, tt.args.isDebug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandlerData() = %v, want %v", got, tt.want)
			}
		})
	}
}
