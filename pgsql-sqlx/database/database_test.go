package database

import (
	"testing"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		code int
	}{
		{
			name: "test",
			wantErr: true,
			code: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
