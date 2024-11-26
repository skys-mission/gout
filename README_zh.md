# gout

[![CodeQL Advanced](https://github.com/skys-mission/gout/actions/workflows/codeql.yml/badge.svg)](https://github.com/skys-mission/gout/actions/workflows/codeql.yml)
[![Codacy Security Scan](https://github.com/skys-mission/gout/actions/workflows/codacy.yml/badge.svg)](https://github.com/skys-mission/gout/actions/workflows/codacy.yml)

这是一个旨在快速构建业务代码的工具库，易用性和实用性导致它可能不适用与基础架构与底层代码。

其它语言：[English](README.md), (Currently unable to translate more)

# 目录

<!-- TOC -->

* [快速开始](#快速开始)
* [使用互联网公开接口的工具](#使用互联网公开接口的工具)
* [Windows API](#windows-api)
* [可用环境](#可用环境)
* [开发计划](#开发计划)

<!-- TOC -->

# 快速开始

你可以使用api.go中的方法，也可以单独加载util下的某个工具（注意一些无法跨平台的方法，不在api.go中）。

实例代码参考：api_test.go

# 使用互联网公开接口的工具

| package                                      | brief                                        |
|----------------------------------------------|----------------------------------------------|
| github.com/skys-mission/gout/util/iplocation | 通过iplocation.net的公开服务获取IP相关信息，你不应该很频繁的调用公开接口 |

# Windows API

| package                                        | cgo | brief                                                       | windows api                 |
|------------------------------------------------|-----|-------------------------------------------------------------|-----------------------------|
| github.com/skys-mission/gout/util/win/mbw      | no  | 通过windows API 弹出一个消息框。有简单消息，简单错误/警告消息，自定义消息框四个方法。仅支持Windows | user32.dll                  |
| github.com/skys-mission/gout/util/win/displayw | no  | 通过windows API查询显示屏分辨率与帧数。有所有显示器和主显示器两个方法。仅支持Windows         | user32.dll/gdi32.dll        |
| github.com/skys-mission/gout/util/win/systemlw | no  | 通过windows API查询当前系统默认语言。有返回名字和代码两个方法。仅支持Windows             | kernel32.dll                |
| github.com/skys-mission/gout/util/win/processw | no  | 通过windows API提供一系列的进程相关方法，用于查询PID，设置进程优先级和相关性。  仅支持Windows  | kernel32.dll/(advapi32.dll) |

# 可用环境

- Golang最低版本：1.18
    - 为了使用泛型，虽然现在没有使用泛型，但未来可能会使用

# 开发计划

不如你来提ISSUE我来实现，或者你可以自己实现然后PR。