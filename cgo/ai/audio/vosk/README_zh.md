# gout-vosk

为什么不使用原生vosk-api，因为官方的样例我实在无法在Windows上编译通过，我还需要一个二进制vosk工具，我又不想写C，所以就有了这个项目。

其它语言：[English](README.md), (Currently unable to translate more)

<!-- TOC -->
* [gout-vosk](#gout-vosk)
  * [如何编译二进制工具](#如何编译二进制工具)
  * [二进制工具使用](#二进制工具使用)
  * [库使用](#库使用)
  * [支持](#支持)
  * [Vosk-API 说明](#vosk-api-说明)
<!-- TOC -->

## 如何编译二进制工具

你可以先在github release下载二进制文件，编译需要等待我后续准备的DOCKERFILE

## 二进制工具使用

```cmd
Usage: .\gout-vosk.exe
Options:
  -audio string
        Path to the audio file (required)
  -model string
        Path to the Vosk model (required)
  -out string
        Output json file path (Disabled by default, outputs to stdout)
  -rate float
        Audio sample rate (default 16000)
Example:
example 1:
.\gout-vosk.exe -audio=".\audio.wav" -model=".\vosk-model" -rate=16000
example 2:
.\gout-vosk.exe -audio=".\audio.wav" -model=".\vosk-model" -rate=16000 -out=".\putout.json"
```

## 库使用

你除非侵入本库仓库，把vosk-api的动态链接库放入lib文件夹中，否则无法成功编译本库，我目前没有找到很好的办法处理这个事，除非我把.dll文件上传到github，但这不是我想要的，我需要更友好的办法。

## 支持

暂时只支持Windows，计划支持macOS和Linux(暂无排期)。

## Vosk-API 说明

vosk-api相关文件在Apache License Version 2.0下发行

vosk-api项目地址：https://github.com/alphacep/vosk-api

vosk模型下载页：https://alphacephei.com/vosk/models