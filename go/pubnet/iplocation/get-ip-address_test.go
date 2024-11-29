// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package iplocation

import (
	"context"
	"testing"
)

func TestRequestGetIPAddress(t *testing.T) {
	// 创建一个成功的测试案例
	t.Run("success", func(t *testing.T) {

		// 创建一个上下文
		ctx := context.Background()

		// 调用待测函数，并处理结果
		got, err := RequestGetIPAddress(ctx)
		if err != nil {
			t.Errorf("RequestGetIPAddress() error = %v, wantErr %v", err, nil)
			return
		}
		t.Logf("RequestGetIPAddress() = %+v", got)
	})
}

func TestRequestGetIPAddressOnlyV4(t *testing.T) {
	// 创建一个成功的测试案例
	t.Run("success", func(t *testing.T) {

		// 创建一个上下文
		ctx := context.Background()

		// 调用待测函数，并处理结果
		got, err := RequestGetIPAddressOnlyV4(ctx)
		if err != nil {
			t.Errorf("RequestGetIPAddress() error = %v, wantErr %v", err, nil)
			return
		}
		t.Logf("RequestGetIPAddress() = %+v", got)
	})
}
