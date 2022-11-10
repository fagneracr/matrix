package verifymatrix

import "testing"

func TestArrayisValid(t *testing.T) {

	tests := []struct {
		name      string
		args      []string
		wantValid bool
	}{
		{
			name:      "Teste1",
			args:      []string{"aaa", "bbbb"},
			wantValid: false,
		},
		{
			name: "Teste2",
			args: []string{"DUHBHB",
				"DUBUBD"},
			wantValid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValid := ArrayisValid(tt.args); gotValid != tt.wantValid {
				t.Errorf("ArrayisValid() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}
