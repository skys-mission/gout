# gout

[![CodeQL Advanced](https://github.com/skys-mission/gout/actions/workflows/codeql.yml/badge.svg)](https://github.com/skys-mission/gout/actions/workflows/codeql.yml)
[![Codacy Security Scan](https://github.com/skys-mission/gout/actions/workflows/codacy.yml/badge.svg)](https://github.com/skys-mission/gout/actions/workflows/codacy.yml)

我注意到杂乱无章的项目（虽然本项目有清晰的目录，但涉及的领域太多），不利于项目发展，所以该仓库已经弃用。

由于我发现我的三方博客关于该项目有一定的收藏量，所以保留该仓库以供调用。如果出现BUG请提ISSUE，我依旧会修复。

其它语言：[English](README.md), (Currently unable to translate more)

<!-- TOC -->
* [gout](#gout)
  * [如何食用](#如何食用)
  * [AI](#ai)
    * [本地AI](#本地ai)
      * [音频处理](#音频处理)
        * [Vosk(CGO)](#voskcgo)
  * [公开网络服务](#公开网络服务)
    * [获取IP信息](#获取ip信息)
      * [iplocation.net](#iplocationnet)
  * [GUI](#gui)
    * [Windows平台](#windows平台)
      * [Windows消息框（无CGO，依赖系统API）](#windows消息框无cgo依赖系统api)
  * [操作系统相关](#操作系统相关)
    * [Windows 平台](#windows-平台)
      * [查询系统信息（无CGO）](#查询系统信息无cgo)
      * [进程管理（无CGO）](#进程管理无cgo)
      * [权限相关（无CGO）](#权限相关无cgo)
  * [编译相关](#编译相关)
    * [GO类库（无需编译）和GO工具](#go类库无需编译和go工具)
    * [CGO](#cgo)
  * [版本兼容](#版本兼容)
  * [开发计划](#开发计划)
    * [当前计划](#当前计划)
<!-- TOC -->

## 如何食用

**Go类库（无CGO）正常加载既可**

```cmd
go get -u github.com/skys-mission/gout
```

推荐在代码页使用`import 子目录`，之后使用`go mod tidy`或IDE工具拉取项目

## AI

### 本地AI

#### 音频处理

##### Vosk(CGO)
- 包：github.com\skys-mission\gout\cgo\ai\audio\vosk
  - 功能：调用Vosk引擎进行音频转文本，使用本地模型，暂时不支持类库（用源码也可以），github release有二进制文件和封装好的包。
  - 使用文档：[gout-vosk文档](cgo/ai/audio/vosk/README.md)

## 公开网络服务

**你应当遵循网络服务提供者的使用条款。
本项目只是提供了Go语言调用网络服务的代码，网络服务本身与本项目无关，请根据对应服务官方条款使用。**

### 获取IP信息

#### iplocation.net

- 包：github.com/skys-mission/gout/go/pubnet/iplocation
  - 功能：查询指定IP信息，查询出口公网IP，查询当前网络环境是否在Chinese-GW内。
  - 使用文档：[iplocation库文档](go/pubnet/iplocation/README.md)

## GUI

### Windows平台

#### Windows消息框（无CGO，依赖系统API）

- 包：github.com/skys-mission/gout/go/gui/win/mbw
  - 通过windows API 弹出一个消息框。有简单消息，简单错误/警告消息，自定义消息框四个方法。仅支持Windows。
  - 使用文档：[mbw库文档](go/gui/win/mbw/README.md)
  - 依赖：user32.dll（一般正常安装的Windows系统都包含该API,无需特别设置，可直接使用本类库）

## 操作系统相关

### Windows 平台

#### 查询系统信息（无CGO）

- 包：github.com/skys-mission/gout/go/os/win/displayw
  - 通过windows API查询显示屏分辨率与帧数，颜色信息等参数。有所有显示器和主显示器两个方法。仅支持Windows。
  - 使用文档：[displayw库文档](go/os/win/displayw/README.md)
  - 依赖：user32.dll（一般正常安装的Windows系统都包含该API,无需特别设置，可直接使用本类库）

- 包：github.com/skys-mission/gout/go/os/win/systemlw
  - 通过windows API查询当前操作系统使用的语言。有返回LCID标准和IANA标准国家代码两个方法。仅支持Windows。
  - 使用文档：[systemlw库文档](go/os/win/systemlw/README.md)
  - 依赖：kernel32.dll（一般正常安装的Windows系统都包含该API,无需特别设置，可直接使用本类库）
  - 注意：LCID是Windows API的原生返回（十六进制），会直接返回代码，而IANA这里只写了几个国家的语言代码，需要的话可以提PR或使用LCID。

#### 进程管理（无CGO）

- 包：github.com/skys-mission/gout/go/os/win/processw
  - 功能：查询所有进程，设置进程优先级和相关性等方法，我在生产中使用过，但没有完善过相关文档，现在我自己也有点看不懂了。仅支持Windows。
  - 文档待完善
  - 依赖：kernel32.dll（一般正常安装的Windows系统都包含该API,无需特别设置，可直接使用本类库）

#### 权限相关（无CGO）

- 包：github.com/skys-mission/gout/go/os/win/privilegew
  - 功能：申请Windows调试权限，比管理员权限更高，首先需要工作在管理员权限之下，我在生产中使用过，但没有完善过相关文档，现在我自己也有点看不懂了。仅支持Windows。
  - 文档待完善
  - 依赖：advapi32.dll（一般正常安装的Windows系统都包含该API,无需特别设置，可直接使用本类库）

## 编译相关

### GO类库（无需编译）和GO工具

- Golang最低版本：1.18
    - 为了使用泛型，虽然现在没有使用泛型，但未来可能会使用

### CGO

**详见CGO项目文档，因为CGO工具链原因，不同项目编译方式不同**

## 版本兼容

虽然自本项目v0.2版本后，本项目将尽可能的保证向上兼容，但如果遇到安全隐患，BUG或其它风险，本项目将直接在小版本破坏向上兼容（这通常不会是全局性的）。

## 开发计划

### 当前计划

| 描述                 | 方向      | 类型     | 排期               |
|--------------------|---------|--------|------------------|
| 通过ip-api.com获取IP信息 | GO类库    | 公开网络服务 | 无排期              |
| 完善cgo-vosk库        | GO类库和工具 | 本地AI功能 | 音频转文本已完成，完善功能计划中 |
