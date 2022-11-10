package verifymatrix

import "testing"

func TestFindSequence(t *testing.T) {

	tests := []struct {
		name           string
		args           []string
		wantCountFound int
	}{

		{
			name: "Valid",
			args: []string{"DDDDDD",
				"BBBBBB",
				"HHHHHH",
				"UUUUUU",
				"DUHBHB",
				"DUHBHB"},
			wantCountFound: 4,
		},
		{
			name: "Valid",
			args: []string{"DUHBHB",
				"DUBUBD",
				"UBUUHU",
				"BHBDHH",
				"BDDDDUB",
				"UDBDUH"},
			wantCountFound: 1,
		},
		{
			name: "InValid",
			args: []string{"DUHBHB",
				"DUBUBD",
				"UBUUHU",
				"BHBDHH",
				"BDDUDUB",
				"UDBDUH"},
			wantCountFound: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCountFound := FindSequence(tt.args); gotCountFound != tt.wantCountFound {
				t.Errorf("FindSequence() = %v, want %v", gotCountFound, tt.wantCountFound)
			}
		})
	}
}
