package verifymatrix

import (
	"reflect"
	"testing"
)

func TestFindDiagonais(t *testing.T) {

	tests := []struct {
		name           string
		args           []string
		wantMtdiagonal []string
	}{
		{
			name: "Teste1",
			args: []string{"AAAAA",
				"BBBBB",
				"CCCCC",
				"AAAAA"},
			wantMtdiagonal: []string{"ABCA"},
		},
		{
			name: "Teste1",
			args: []string{"AAAAAA",
				"BBBBBB",
				"CCCCCC",
				"DDDDDD",
				"EEEEEE",
			},
			wantMtdiagonal: []string{"ABCDE", "BCDE"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMtdiagonal := FindDiagonais(tt.args); !reflect.DeepEqual(gotMtdiagonal, tt.wantMtdiagonal) {
				t.Errorf("FindDiagonais() = %v, want %v", gotMtdiagonal, tt.wantMtdiagonal)
			}
		})
	}
}
