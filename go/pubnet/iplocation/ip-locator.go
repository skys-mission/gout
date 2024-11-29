// Copyright (c) 2024, https://github.com/skys-mission and SoyMilkWhisky

package iplocation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// IPLocatorResp represents the response structure for IP location query.
// It provides detailed information about the IP address, including its location and network details.
type IPLocatorResp struct {
	IP              string                   `json:"ip"`               // IPv4 or IPv6 address used to lookup geolocation.
	IPNumber        string                   `json:"ip_number"`        // IP number in long integer.
	IPVersion       ResponseIPVersion        `json:"ip_version"`       // IP version either 4 or 6.
	CountryName     string                   `json:"country_name"`     // Full name of the IP country.
	CountryCode2    string                   `json:"country_code2"`    // ISO ALPHA-2 Country Code.
	ISP             string                   `json:"isp"`              // Internet Service Provider (ISP) who owns the IP address.
	ResponseCode    ResponseCodeForIPLocator `json:"response_code"`    // Response status code to indicate success or failed completion of the API call.
	ResponseMessage ResponseMessage          `json:"response_message"` // Response message to indicate success or failed completion of the API call.
}

// RequestIPLocator performs an HTTP request to an IP location service to get the geographical location information of the specified IP.
// This function takes a context and an IP address string as input, and returns the parsed response data or an error.
// The ctx parameter is used to cancel the request, manage deadlines, or handle request cancellation.
// The ip parameter is the IP address whose geographical location information is to be queried.
// The function returns a pointer to IPLocatorResp type which contains the response data, or an error if the operation fails.
func RequestIPLocator(ctx context.Context, ip string) (resp *IPLocatorResp, err error) {
	// Check if the ip parameter is empty
	if ip == "" {
		return nil, fmt.Errorf("ip cannot be empty")
	}

	// Prepare URL parameters
	params := url.Values{}
	params.Add("ip", ip)

	// Construct the complete request URL
	reqURL := fmt.Sprintf("%s/?%s", baseIplocationURL, params.Encode())

	// Create an HTTP GET request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Send the HTTP request
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	// Ensure the response body is closed
	defer httpResp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Check the HTTP response status code
	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d: %s", httpResp.StatusCode, string(body))
	}

	// Parse JSON response
	var result *IPLocatorResp
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	// Return the parsed response data
	return result, nil
}
