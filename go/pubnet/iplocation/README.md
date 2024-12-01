# iplocation

Other languages: [简体中文](README_zh.md), (Currently unable to translate more)

Function: Query specified IP information, query the public IP of the exit, and check if the current network environment is within the Chinese-GW.

**You should comply with the terms of service of the internet service provider. This project only provides code for
calling network services in Go language, and the network services themselves are not related to this project. Please use
them in accordance with the official terms of the respective services.**

[iplocation官方地址](https://api.iplocation.net/)

<!-- TOC -->
* [iplocation](#iplocation)
* [How to Use](#how-to-use)
* [Code Examples](#code-examples)
  * [Native Interface](#native-interface)
    * [Get Your Public Network Exit IP](#get-your-public-network-exit-ip)
    * [Query Detailed IP Information (Region, ISP, etc.)](#query-detailed-ip-information-region-isp-etc)
  * [gout Methods](#gout-methods)
    * [Check if the Current Public Network Exit IP is Affected by Chinese-GW](#check-if-the-current-public-network-exit-ip-is-affected-by-chinese-gw)
    * [Check if the Current Public Network Exit IP is Affected by Chinese-GW and Additional Regions](#check-if-the-current-public-network-exit-ip-is-affected-by-chinese-gw-and-additional-regions)
* [Notes](#notes)
* [Region Code Lookup](#region-code-lookup)
<!-- TOC -->

# How to Use

```shell
go get -u github.com/skys-mission/gout/go/pubnet/iplocation
```

# Code Examples

## Native Interface

Note that you should check if the ResponseCode in the returned Resp is 200 (note that this 200 is not http 200, http 200
has already been handled within this library).

Since different interfaces may return integers or strings, you should use constants from this library or carefully
perform type conversions, refer to the example code.

### Get Your Public Network Exit IP

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/skys-mission/gout/go/pubnet/iplocation"
)

func main() {
	ctx, cf := context.WithTimeout(context.Background(), time.Second)
	defer cf()
	result, err := iplocation.RequestGetIPAddress(ctx)
	if err != nil {
		panic(err)
	}
	if result.ResponseCode != iplocation.ResponseCodeOKForGetIPAddress {
		panic(fmt.Errorf("code: %v msg: %v", result.ResponseCode, result.ResponseMessage))
	}
	fmt.Printf("%+v\ntype: %T", result, result)
}

```

**Example Output**

```cmd
&{IP:your_ip_address IPVersion:4 ResponseCode:200 ResponseMessage:OK}
type: *iplocation.GetIPAddressResp
Process finished with the exit code 0
```

### Query Detailed IP Information (Region, ISP, etc.)

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/skys-mission/gout/go/pubnet/iplocation"
)

func main() {
	ctx, cf := context.WithTimeout(context.Background(), time.Second)
	defer cf()
	result, err := iplocation.RequestIPLocator(ctx, "8.8.8.8")
	if err != nil {
		panic(err)
	}
	if result.ResponseCode != iplocation.ResponseCodeOKForIPLocator {
		panic(fmt.Errorf("code: %v msg: %v", result.ResponseCode, result.ResponseMessage))
	}
	fmt.Printf("%+v\ntype: %T", result, result)
}

```

**Example Output**

```cmd
&{IP:8.8.8.8 IPNumber:134744072 IPVersion:4 CountryName:United States of America CountryCode2:US ISP:Google LLC ResponseCode:200 ResponseMessage:OK}
type: *iplocation.IPLocatorResp
Process finished with the exit code 0
```

## gout Methods

### Check if the Current Public Network Exit IP is Affected by Chinese-GW

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/skys-mission/gout/go/pubnet/iplocation"
)

func main() {
	ctx, cf := context.WithTimeout(context.Background(), time.Second)
	defer cf()
	isOutGW, err := iplocation.IsOutsideGW(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(isOutGW)
}

```

**Example Output**

```cmd
true
Process finished with the exit code 0
```

### Check if the Current Public Network Exit IP is Affected by Chinese-GW and Additional Regions

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/skys-mission/gout/go/pubnet/iplocation"
)

func main() {
	ctx, cf := context.WithTimeout(context.Background(), time.Second)
	defer cf()
	isOut, err := iplocation.IsOutsideGWithExtra(
		ctx,
		iplocation.LocalCodeChinaMacao,
		iplocation.LocalCodeChinaHongKong)
	if err != nil {
		panic(err)
	}
	fmt.Println(isOut)
}

```

**Example Output**

```cmd
false
Process finished with the exit code 0
```

# Notes

1. iplocation.net is a publicly used interface globally, but the official documentation does not provide rate-limiting
   information. You should not frequently call a non-commercial interface.
2. Recommendation: Limit QPS to 1QPS, and DPS to 1000QPS. You should use caching technology to reduce the number of
   accesses.
3. If you need this project to implement in-memory caching or rate limiting, you can raise an ISSUE or PR, and I will
   prioritize implementing it.

# Region Code Lookup

Based on my investigation and analysis, the CountryCode2 field
uses: [ISO_3166-1_alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)