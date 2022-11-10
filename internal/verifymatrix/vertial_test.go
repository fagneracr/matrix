package verifymatrix

import (
	"reflect"
	"testing"
)

func TestBuildVertical(t *testing.T) {

	tests := []struct {
		name        string
		args        []string
		wantMatrixv []string
	}{
		{
			name:        "vertical Test1",
			args:        []string{"AA", "BB"},
			wantMatrixv: []string{"AB", "AB"},
		},
		{
			name:        "vertical Test2",
			args:        []string{"ABC", "DEF"},
			wantMatrixv: []string{"AD", "BE", "CF"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMatrixv := BuildVertical(tt.args); !reflect.DeepEqual(gotMatrixv, tt.wantMatrixv) {
				t.Errorf("BuildVertical() = %v, want %v", gotMatrixv, tt.wantMatrixv)
			}
		})
	}
}
