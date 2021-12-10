package book

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetAll(t *testing.T) {
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetAll(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
