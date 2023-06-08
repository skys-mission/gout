package iplocation

// power by api.iplocation.net

import "context"

const (
	GetIPAddressOkRespCode         = 200
	GetIPAddressBadRequestRespCode = 400
	GetIPAddressNotFoundRespCode   = 404

	IPLocatorOkRespCode         = "200"
	IPLocatorBadRequestRespCode = "400"
	IPLocatorNotFoundRespCode   = "404"
)

// Client api interface
type Client interface {
	IPLocator(ctx context.Context, ip string) (*IPLocatorResp, error)

	GetIPAddress(ctx context.Context) (*GetIPAddressResp, error)
}

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
	IPNumber        string `json:"ip_number"` // IPVersion just 4/6
	IPVersion       int8   `json:"ip_version"`
	CountryName     string `json:"country_name"`
	CountryCode2    string `json:"country_code2"` // CountryCode2 maybe US ?
	ISP             string `json:"isp"`
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}
