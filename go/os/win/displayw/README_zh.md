# displayw

通过windows API查询显示屏分辨率与帧数，颜色信息等参数。有所有显示器和主显示器两个方法。仅支持Windows。

其它语言：[English](README.md), (Currently unable to translate more)

<!-- TOC -->
* [displayw](#displayw)
* [如何食用](#如何食用)
* [代码样例](#代码样例)
  * [查询所有显示器信息](#查询所有显示器信息)
  * [显示主显示器信息](#显示主显示器信息)
<!-- TOC -->

# 如何食用

```shell
go get -u github.com/skys-mission/gout/go/os/win/displayw
```

# 代码样例

## 查询所有显示器信息

```go
package main

import (
	"fmt"

	"github.com/skys-mission/gout/go/os/win/displayw"
)

func main() {
	displays, err := displayw.GetAllDisplayInfo()
	if err != nil {
		panic(err)
	}
	for _, v := range displays {
		fmt.Printf("%+v\n", v)
	}
}

```

**示例输出**

```cmd
&{Name:\\.\DISPLAY1 Width:2560 Height:1440 RefreshRate:75 BitDepth:32 ColorFormat:RGBA IsPrimary:true}
&{Name:\\.\DISPLAY2 Width:1920 Height:1080 RefreshRate:60 BitDepth:32 ColorFormat:RGBA IsPrimary:false}
&{Name:\\.\DISPLAY3 Width:2560 Height:1440 RefreshRate:60 BitDepth:32 ColorFormat:RGBA IsPrimary:false}
```

## 显示主显示器信息

```go
package main

import (
	"fmt"

	"github.com/skys-mission/gout/go/os/win/displayw"
)

func main() {
	displays, err := displayw.GetPrimaryDisplayInfo()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", displays)
}

```

**示例输出**

```cmd
&{Name:\\.\DISPLAY1 Width:2560 Height:1440 RefreshRate:75 BitDepth:32 ColorFormat:RGBA IsPrimary:true}
```