package features

import "testing"

func TestCalculatePageCount(t *testing.T) {
	type args struct {
		objectsCount int
		itemsPerPage int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Clear 1",
			args{
				objectsCount: 10,
				itemsPerPage: 2,
			},
			5,
		},
		{
			"Clear 2",
			args{
				objectsCount: 10,
				itemsPerPage: 10,
			},
			1,
		},
		{
			"Float 1",
			args{
				objectsCount: 11,
				itemsPerPage: 3,
			},
			4,
		},
		{
			"Float 2",
			args{
				objectsCount: 1009,
				itemsPerPage: 2,
			},
			505,
		},
		{
			"Clear 3",
			args{
				objectsCount: 1008443,
				itemsPerPage: 1,
			},
			1008443,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePageCount(tt.args.objectsCount, tt.args.itemsPerPage); got != tt.want {
				t.Errorf("CalculatePageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
