// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package iplocation

import (
	"context"
	"testing"
)

func TestRequestIPLocator(t *testing.T) {

	// Test cases
	tests := []struct {
		name    string
		ip      string
		wantErr bool
	}{
		{"ValidIP", "8.8.8.8", false},
		{"EmptyIP", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := RequestIPLocator(ctx, tt.ip)
			if (err != nil) != tt.wantErr {
				t.Errorf("RequestIPLocator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantErr {
				t.Errorf("RequestIPLocator() got = %v, want %v", got, tt.ip)
				return
			}
			t.Logf("RequestIPLocator() got = %+v", got)
		})
	}
}
