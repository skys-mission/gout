package iplocation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// power by api.iplocation.net

const (
	GetIPAddressOkRespCode         = 200
	GetIPAddressBadRequestRespCode = 400
	GetIPAddressNotFoundRespCode   = 404

	IPLocatorOkRespCode         = "200"
	IPLocatorBadRequestRespCode = "400"
	IPLocatorNotFoundRespCode   = "404"
)

// GetIPAddressResp ...
type GetIPAddressResp struct {
	IP              string `json:"ip"`
	IPVersion       int8   `json:"ip_version"` // IPVersion just 4/6
	ResponseCode    int64  `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

// IPLocatorResp ...
type IPLocatorResp struct {
	IP              string `json:"ip"`
	IPNumber        string `json:"ip_number"`
	IPVersion       int8   `json:"ip_version"`       // IPVersion just 4/6
	CountryName     string `json:"country_name"`     // CountryName e.g. United States of America
	CountryCode2    string `json:"country_code2"`    // CountryCode2 e.g. US
	ISP             string `json:"isp"`              // ISP e.g. Google LLC
	ResponseCode    string `json:"response_code"`    // ResponseCode ok is 200
	ResponseMessage string `json:"response_message"` // ResponseMessage e.g. ok
}

// IPLocatorByIP Query the ISP information for an IP address based on the IP from api.iplocation.net.
func IPLocatorByIP(ctx context.Context, ip string) (resp *IPLocatorResp, err error) {
	params := url.Values{}
	params.Add("ip", ip)

	reqURL := baseURL + "/?"
	reqURL += params.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	result := new(IPLocatorResp)
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ResponseCode != IPLocatorOkRespCode {
		return nil, errors.New(result.ResponseMessage)
	}
	return result, err
}

// IPAddressLocator Obtain the IP address for accessing the internet from api.iplocation.net.
func IPAddressLocator(ctx context.Context) (*GetIPAddressResp, error) {
	httpResp, err := http.Get(fmt.Sprintf("%s%s", baseURL, getIPAddrPath))
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	result := new(GetIPAddressResp)
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.ResponseCode != GetIPAddressOkRespCode {
		return nil, errors.New(result.ResponseMessage)
	}
	return result, err
}
