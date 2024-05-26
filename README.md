# gout

This is a tool library designed for quickly building business code, its ease of use and practicality may make it
unsuitable for infrastructure and low-level code.

Other languages: [简体中文](README_zh.md), (Currently unable to translate more)

# Index

<!-- TOC -->

* [Quick Start](#quick-start)
* [Tools utilizing public APIs](#tools-utilizing-public-apis)
* [Windows API](#windows-api)
* [Available Environment](#available-environment)
* [Development Roadmap](#development-roadmap)

<!-- TOC -->

# Quick Start

You can use the methods in api.go or load a specific utility from the util folder individually(Note that some
platform-specific methods are not included in api.go).

Example code reference: api_test.go

# Tools utilizing public APIs

| package                                      | brief                                                                                                                          |
|----------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------|
| github.com/skys-mission/gout/util/iplocation | Obtain IP-related information through the public service of iplocation.net. You should not call the public API very frequently |

# Windows API

| package                                        | cgo | brief                                                                                                                                                                | windows api          |
|------------------------------------------------|-----|----------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------|
| github.com/skys-mission/gout/util/win/mbw      | no  | Utilize the Windows API to pop up a message box. There are three methods: simple message, simple error, and custom message box. Only supports Windows                | user32.dll           |
| github.com/skys-mission/gout/util/win/displayw | no  | Utilize the Windows API to query the display resolution and refresh rate. There are two methods: for all monitors and for the primary monitor. Only supports Windows | user32.dll&gdi32.dll |
| github.com/skys-mission/gout/util/win/systemlw | no  | Utilize the Windows API to query the current system's default language. There are two methods: returning the name and the code. Only supports Windows                | kernel32.dll         |

# Available Environment

- Minimum Golang version: 1.18
    - To use generics

# Development Roadmap

There is no definitive roadmap; we will proceed step by step as we go along.