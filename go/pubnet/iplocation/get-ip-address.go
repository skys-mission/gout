// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package iplocation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetIPAddressResp represents the response structure for an IP address query.
// It includes the IP address, IP version, response code, and response message.
type GetIPAddressResp struct {
	IP              string                      `json:"ip"`               // The queried IP address
	IPVersion       ResponseIPVersion           `json:"ip_version"`       // The version of the IP address
	ResponseCode    ResponseCodeForGetIPAddress `json:"response_code"`    // The response code for the query
	ResponseMessage ResponseMessage             `json:"response_message"` // The response message for the query
}

// RequestGetIPAddress sends a GET request to the server to obtain IP address information.
// It takes a context as a parameter to cancel the request if needed.
// It returns a pointer to GetIPAddressResp containing the response information, or an error if the request fails.
func RequestGetIPAddress(ctx context.Context) (*GetIPAddressResp, error) {
	// Construct the request URL
	url := fmt.Sprintf("%s%s", baseIplocationURL, getIPAddressSubURLPath)

	// Create a GET request with the specified context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		// Return an error if the request creation fails
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Send the HTTP request
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Return an error if the request execution fails
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	// Ensure the response body is closed
	defer httpResp.Body.Close()

	// Check the HTTP status code
	if httpResp.StatusCode != http.StatusOK {
		// Return an error if an unexpected status code is received
		return nil, fmt.Errorf("unexpected status code: %d, status: %s", httpResp.StatusCode, httpResp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		// Return an error if reading the response body fails
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	var result *GetIPAddressResp
	if err = json.Unmarshal(body, &result); err != nil {
		// Return an error if JSON parsing fails
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Return the parsed result
	return result, nil
}

// RequestGetIPAddressOnlyV4 请求获取IPv4地址信息。
// 该函数通过发送GET请求到指定的URL来获取IPv4地址信息，并解析响应体为GetIPAddressResp类型。
// 参数:
//
//	ctx context.Context: 上下文对象，用于取消请求和传递请求级值。
//
// 返回值:
//
//	*GetIPAddressResp: 一个指向GetIPAddressResp类型的指针，包含IPv4地址信息。
//	error: 错误对象，如果请求失败或解析响应体失败，则返回错误。
func RequestGetIPAddressOnlyV4(ctx context.Context) (*GetIPAddressResp, error) {
	// 创建GET请求
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ipv4IplocationURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// 发送HTTP请求
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}

	// 确保在函数返回前关闭响应体
	defer httpResp.Body.Close()

	// 检查HTTP状态码
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, status: %s", httpResp.StatusCode, httpResp.Status)
	}

	// 读取响应体
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// 解析JSON响应
	var result *GetIPAddressResp
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// 返回解析后的结果
	return result, nil
}
