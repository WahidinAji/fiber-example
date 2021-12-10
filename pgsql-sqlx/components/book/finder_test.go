package book

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestFinder(t *testing.T) {
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
			if err := Finder(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Finder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
