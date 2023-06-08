package iplocation

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// client ...
type client struct {
	*resty.Client
}

// New a Client
func New() Client {
	c := &client{
		Client: resty.New(),
	}
	c.BaseURL = baseURL
	return c
}

// GetIPAddress ...
func (c *client) GetIPAddress(ctx context.Context) (*GetIPAddressResp, error) {
	resp, err := c.R().SetContext(ctx).Get(fmt.Sprintf("%s%s", c.BaseURL, getIPAddrPath))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	result := new(GetIPAddressResp)
	if err = json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, errors.WithStack(err)
	}
	if result.ResponseCode != GetIPAddressOkRespCode {
		return nil, errors.New(result.ResponseMessage)
	}
	return result, err
}

// IPLocator ...
func (c *client) IPLocator(ctx context.Context, ip string) (*IPLocatorResp, error) {
	resp, err := c.R().SetContext(ctx).SetQueryParam("ip", ip).Get(c.BaseURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	result := new(IPLocatorResp)
	if err = json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, errors.WithStack(err)
	}
	if result.ResponseCode != IPLocatorOkRespCode {
		return nil, errors.New(result.ResponseMessage)
	}
	return result, err
}
