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

You can use the methods in api.go or load a specific utility from the util folder individually.

Example code reference: api_test.go

# Tools utilizing public APIs

| package                                      | brief                                                                                                                          |
|----------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------|
| github.com/skys-mission/gout/util/iplocation | Obtain IP-related information through the public service of iplocation.net. You should not call the public API very frequently |

# Windows API

| package                                   | brief                                 |
|-------------------------------------------|---------------------------------------|
| github.com/skys-mission/gout/util/win/mbw | Pop up a message box via Windows API. |

# Available Environment

- Minimum Golang version: 1.18
    - To use generics

# Development Roadmap

There is no definitive roadmap; we will proceed step by step as we go along.