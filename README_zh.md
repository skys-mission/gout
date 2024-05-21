# gout

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

你可以使用api.go中的方法，也可以单独加载util下的某个工具。

实例代码参考：api_test.go

# 使用互联网公开接口的工具

| package                                      | brief                                        |
|----------------------------------------------|----------------------------------------------|
| github.com/skys-mission/gout/util/iplocation | 通过iplocation.net的公开服务获取IP相关信息，你不应该很频繁的调用公开接口 |

# Windows API

| package                                   | brief                 |
|-------------------------------------------|-----------------------|
| github.com/skys-mission/gout/util/win/mbw | 通过windows API 弹出一个消息框 |

# 可用环境

- Golang最低版本：1.18
    - 为了使用泛型，虽然现在没有使用泛型，但未来可能会使用

# 开发计划

没有什么明确的计划，走一步看一步吧