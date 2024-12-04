// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package iplocation

import (
	"context"
	"testing"
)

func TestIsOutsideGW(t *testing.T) {
	outsideGW, err := IsOutsideGW(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(outsideGW)
}

func TestIsOutsideGWithExtra(t *testing.T) {
	outsideGW, err := IsOutsideGWithExtra(context.Background(), LocalCodeChinaMacao, LocalCodeChinaHongKong)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(outsideGW)
}
