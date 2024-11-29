# displayw

Query display resolution, frame rate, color information, and other parameters through the Windows API. There are two methods: one for all monitors and one for the primary monitor. Only supports Windows.

Other languages: [简体中文](README_zh.md), (Currently unable to translate more)

<!-- TOC -->
* [displayw](#displayw)
* [How to Use](#how-to-use)
* [Code Examples](#code-examples)
  * [Query Information for All Displays](#query-information-for-all-displays)
  * [Display Information for the Primary Monitor](#display-information-for-the-primary-monitor)
<!-- TOC -->

# How to Use

```shell
go get -u github.com/skys-mission/gout/go/os/win/displayw
```

# Code Examples

## Query Information for All Displays

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

**Example Output**

```cmd
&{Name:\\.\DISPLAY1 Width:2560 Height:1440 RefreshRate:75 BitDepth:32 ColorFormat:RGBA IsPrimary:true}
&{Name:\\.\DISPLAY2 Width:1920 Height:1080 RefreshRate:60 BitDepth:32 ColorFormat:RGBA IsPrimary:false}
&{Name:\\.\DISPLAY3 Width:2560 Height:1440 RefreshRate:60 BitDepth:32 ColorFormat:RGBA IsPrimary:false}
```

## Display Information for the Primary Monitor

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

**Example Output**

```cmd
&{Name:\\.\DISPLAY1 Width:2560 Height:1440 RefreshRate:75 BitDepth:32 ColorFormat:RGBA IsPrimary:true}
```