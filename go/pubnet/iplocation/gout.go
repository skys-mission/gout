// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package iplocation

import (
	"context"
	"fmt"
)

// IsOutsideGW 判断请求是否来自中国大陆境外的网关。
// 该函数通过获取请求者的IP地址，并查询IP地址所属的国家代码来判断。
// 如果IP地址所属国家代码不是中国大陆，或者查询过程中出现错误，则认为请求来自境外。
// 参数:
//
//	ctx context.Context: 上下文对象，用于传递请求相关的配置和元数据。
//
// 返回值:
//
//	isOutGW bool: 请求是否来自境外的网关。如果是，返回true；否则，返回false。
//	err error: 错误对象，如果在判断过程中出现错误，返回具体的错误信息；否则，返回nil。
func IsOutsideGW(ctx context.Context) (isOutGW bool, err error) {
	// 获取请求者的IP地址信息。
	address, err := RequestGetIPAddress(ctx)
	if err != nil {
		// 如果获取IP地址信息时出现错误，直接返回。
		return
	}
	// 检查获取IP地址的响应码是否正常。
	if address.ResponseCode != ResponseCodeOKForGetIPAddress {
		// 如果响应码不正常，返回错误信息。
		return false, fmt.Errorf("error: %s", address.ResponseMessage)
	}
	// 查询IP地址的地理位置信息。
	resp, err := RequestIPLocator(ctx, address.IP)
	if err != nil {
		// 如果查询地理位置信息时出现错误，直接返回。
		return
	}
	// 检查查询地理位置信息的响应码是否正常。
	if resp.ResponseCode != ResponseCodeOKForIPLocator {
		// 如果响应码不正常，返回错误信息。
		return false, fmt.Errorf("error: %s", resp.ResponseMessage)
	}
	// 判断IP地址所属的国家代码是否为中国大陆。
	if resp.CountryCode2 != string(LocalCodeChinaMainland) {
		// 如果不是中国大陆，认为请求来自境外。
		return true, nil
	}
	// 如果是中国大陆，认为请求来自境内。
	return
}

// IsOutsideGWithExtra 判断给定IP地址是否位于中国大陆之外的地区。
// 该函数通过发送网络请求获取IP地址信息，并根据ISO3166V1Alpha2标准的国家代码判断IP地址的位置。
// 参数:
//
//	ctx: 上下文，用于传递请求超时、取消信号等信息。
//	locals: 可变参数，包含需要特别考虑的ISO3166V1Alpha2标准国家代码。
//
// 返回值:
//
//	isOut: 布尔值，如果IP地址位于中国大陆之外且不在locals参数中指定的国家内，则为true，否则为false。
//	err: 错误值，如果在获取IP地址信息或定位IP地址位置时发生错误，则返回相应的错误。
func IsOutsideGWithExtra(ctx context.Context, locals ...ISO3166V1Alpha2) (isOut bool, err error) {
	// 获取请求IP地址信息。
	address, err := RequestGetIPAddress(ctx)
	if err != nil {
		// 如果发生错误，直接返回。
		return
	}
	// 检查获取IP地址信息的响应码。
	if address.ResponseCode != ResponseCodeOKForGetIPAddress {
		// 如果响应码不表示成功，返回错误信息。
		return false, fmt.Errorf("error: %s", address.ResponseMessage)
	}
	// 使用获取的IP地址请求IP定位信息。
	resp, err := RequestIPLocator(ctx, address.IP)
	if err != nil {
		// 如果发生错误，直接返回。
		return
	}
	// 检查IP定位信息的响应码。
	if resp.ResponseCode != ResponseCodeOKForIPLocator {
		// 如果响应码不表示成功，返回错误信息。
		return false, fmt.Errorf("error: %s", resp.ResponseMessage)
	}

	// 判断IP定位的国家是否为中国大陆。
	if resp.CountryCode2 == string(LocalCodeChinaMainland) {
		// 如果是中国大陆，返回false，表示不在中国大陆之外。
		return false, nil
	}
	// 遍历locals参数中的国家代码。
	for _, v := range locals {
		// 判断IP定位的国家是否在locals参数中。
		if resp.CountryCode2 == string(v) {
			// 如果在locals参数中，返回false，表示不在中国大陆之外。
			return false, nil
		}
	}
	// 如果IP定位的国家既不是中国大陆，也不在locals参数中，返回true，表示在中国大陆之外。
	return true, nil
}
