# iplocation

其它语言：[English](README.md), (Currently unable to translate more)

功能：查询指定IP信息，查询出口公网IP，查询当前网络环境是否在Chinese-GW内。

**你应当遵循网络服务提供者的使用条款。
本项目只是提供了Go语言调用网络服务的代码，网络服务本身与本项目无关，请根据对应服务官方条款使用。**

[iplocation官方地址](https://api.iplocation.net/)

<!-- TOC -->
* [iplocation](#iplocation)
* [如何食用](#如何食用)
* [代码样例](#代码样例)
  * [原生接口](#原生接口)
    * [获取你的公网出口IP](#获取你的公网出口ip)
    * [查询IP详细信息（地域，ISP等）](#查询ip详细信息地域isp等)
  * [gout方法](#gout方法)
    * [查询当前公网IP出口是否在Chinese GW影响之内](#查询当前公网ip出口是否在chinese-gw影响之内)
    * [查询当前公网IP出口是否在Chinese GW和附加地域影响之内](#查询当前公网ip出口是否在chinese-gw和附加地域影响之内)
* [注意事项](#注意事项)
* [地区代码查询](#地区代码查询)
<!-- TOC -->

# 如何食用

```shell
go get -u github.com/skys-mission/gout/go/pubnet/iplocation
```

# 代码样例

## 原生接口

注意你应当判断返回Resp中的ResponseCode是否为200（注意这个200不是http 200，http 200本库内已经处理了）

因不同接口有的是整数，有的是字符串，你应该使用本库中的常量，或者仔细做类型转换，参考示例代码

### 获取你的公网出口IP

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

**输出样例**

```cmd
&{IP:your_ip_address IPVersion:4 ResponseCode:200 ResponseMessage:OK}
type: *iplocation.GetIPAddressResp
Process finished with the exit code 0
```

### 查询IP详细信息（地域，ISP等）

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

**输出样例**

```cmd
&{IP:8.8.8.8 IPNumber:134744072 IPVersion:4 CountryName:United States of America CountryCode2:US ISP:Google LLC ResponseCode:200 ResponseMessage:OK}
type: *iplocation.IPLocatorResp
Process finished with the exit code 0
```

## gout方法

### 查询当前公网IP出口是否在Chinese GW影响之内

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

**输出样例**

```cmd
true
Process finished with the exit code 0
```

### 查询当前公网IP出口是否在Chinese GW和附加地域影响之内

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

**输出样例**

```cmd
false
Process finished with the exit code 0
```

# 注意事项

1. iplocation.net是全球互联网普遍使用的一个公开接口，但官方并没有给出限流文档，你不应该频繁的调用一个非商业接口
2. 建议：QPS限制在1QPS，DPS限制在1000QPS，你应该使用缓存技术，减少访问次数
3. 如果需要本项目实现内存缓存或限流可以提出ISSUE或PR，我会优先实现

# 地区代码查询

根据我的调查和分析，CountryCode2字段采用的是：[ISO_3166-1_alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
