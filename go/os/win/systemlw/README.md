# systemlw

Query the language used by the current operating system through the Windows API. There are two methods: returning the LCID standard and the IANA standard country code. Only supports Windows.

LCID is the native return of the Windows API, which will directly return the code (hexadecimal). The IANA here only lists the language codes of a few countries. If needed, you can submit a PR or use LCID.

Other languages: [简体中文](README_zh.md), (Currently unable to translate more)

<!-- TOC -->
* [systemlw](#systemlw)
  * [How to Use](#how-to-use)
  * [Code Example](#code-example)
    * [Query the language used by the current operating system, return IANA code](#query-the-language-used-by-the-current-operating-system-return-iana-code)
    * [Query the language used by the current operating system, return LCID code](#query-the-language-used-by-the-current-operating-system-return-lcid-code)
  * [Additional Knowledge](#additional-knowledge)
<!-- TOC -->

## How to Use

```cmd
go get -u github.com/skys-mission/gout/go/os/win/systemlw
```

## Code Example

### Query the language used by the current operating system, return IANA code

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

**Example Output**

```cmd
zh-Hans
Process finished with the exit code 0
```

### Query the language used by the current operating system, return LCID code

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

**Example Output**

```cmd
0804
Process finished with the exit code 0
```

## Additional Knowledge

IANA, International Assigned Numbers Authority

LCID: Windows Language Code Identifier

---