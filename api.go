package gout

import (
	"context"

	"github.com/skys-mission/gout/go/pubnet/iplocation"
)

// IPLocatorByIP Query the ISP information for an IP address based on the IP from api.iplocation.net.
func IPLocatorByIP(ctx context.Context, ip string) (resp *iplocation.IPLocatorResp, err error) {
	return iplocation.IPLocatorByIP(ctx, ip)
}

// IPAddressLocator Obtain the IP address for accessing the internet from api.iplocation.net.
func IPAddressLocator(ctx context.Context) (resp *iplocation.GetIPAddressResp, err error) {
	return iplocation.IPAddressLocator(ctx)
}
