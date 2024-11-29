# systemlw

通过windows API查询当前操作系统使用的语言。有返回LCID标准和IANA标准国家代码两个方法。仅支持Windows。

LCID是Windows API的原生返回，会直接返回代码(十六进制)，而IANA这里只写了几个国家的语言代码，需要的话可以提PR或使用LCID。

其它语言：[English](README.md), (Currently unable to translate more)

<!-- TOC -->
* [systemlw](#systemlw)
  * [如何食用](#如何食用)
  * [代码样例](#代码样例)
    * [查询当前操作系统使用的语言，返回IANA代码](#查询当前操作系统使用的语言返回iana代码)
    * [查询当前操作系统使用的语言，返回LCID代码](#查询当前操作系统使用的语言返回lcid代码)
  * [补充知识](#补充知识)
<!-- TOC -->

## 如何食用

```cmd
go get -u github.com/skys-mission/gout/go/os/win/systemlw
```

## 代码样例

### 查询当前操作系统使用的语言，返回IANA代码

```go
package main

import (
	"fmt"

	"github.com/skys-mission/gout/go/os/win/systemlw"
)

func main() {
	displays, err := systemlw.GetSystemLanguageCodeIANA()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", displays)
}

```

**样例输出**

```cmd
zh-Hans
Process finished with the exit code 0
```

### 查询当前操作系统使用的语言，返回LCID代码

```go
package main

import (
	"fmt"

	"github.com/skys-mission/gout/go/os/win/systemlw"
)

func main() {
	displays, err := systemlw.GetSystemLanguageCodeLCID()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", displays)
}

```

**样例输出**

```cmd
0804
Process finished with the exit code 0
```

## 补充知识

IANA，International Assigned Numbers Authority

LCID：Windows Language Code Identifier