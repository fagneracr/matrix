package ginserver

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_setupRouter(t *testing.T) {
	tests := []struct {
		name string
		want *gin.Engine
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setupRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
