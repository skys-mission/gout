# gout

[![CodeQL Advanced](https://github.com/skys-mission/gout/actions/workflows/codeql.yml/badge.svg)](https://github.com/skys-mission/gout/actions/workflows/codeql.yml)
[![Codacy Security Scan](https://github.com/skys-mission/gout/actions/workflows/codacy.yml/badge.svg)](https://github.com/skys-mission/gout/actions/workflows/codacy.yml)

This is a tool library designed for quickly building business code. Due to its ease of use and practicality, it may not
be suitable for infrastructure and low-level code.

Other languages: [简体中文](README_zh.md), (Currently unable to translate more)

<!-- TOC -->
* [gout](#gout)
  * [How to Use](#how-to-use)
  * [Public Network Services](#public-network-services)
    * [Get IP Information](#get-ip-information)
      * [iplocation.net](#iplocationnet)
  * [GUI](#gui)
    * [Windows Platform](#windows-platform)
      * [Windows Message Box (No CGO, dependent on system API)](#windows-message-box-no-cgo-dependent-on-system-api)
  * [Operating System Related](#operating-system-related)
    * [Windows Platform](#windows-platform-1)
      * [Query System Information (No CGO)](#query-system-information-no-cgo)
      * [Process Management (No CGO)](#process-management-no-cgo)
      * [Privilege Related (No CGO)](#privilege-related-no-cgo)
  * [Compilation Related](#compilation-related)
    * [GO Library (No Compilation Required) and GO Tools](#go-library-no-compilation-required-and-go-tools)
    * [CGO](#cgo)
  * [Version Compatibility](#version-compatibility)
  * [Development Plan](#development-plan)
    * [Current Plan](#current-plan)
<!-- TOC -->

## How to Use

**Go library (without CGO) can be loaded normally**

```cmd
go get -u github.com/skys-mission/gout
```

It is recommended to use `import subdirectory` in the code page, and then use `go mod tidy` or IDE tools to pull the
project.

## Public Network Services

**You should follow the terms of service of the network service provider.
This project only provides Go language code to call network services. The network services themselves are unrelated to
this project. Please use them according to the official terms of the corresponding service.**

### Get IP Information

#### iplocation.net

- Package: github.com/skys-mission/gout/go/pubnet/iplocation
- Documentation: [iplocation library documentation](go/pubnet/iplocation/README.md)
- Features: Query specified IP information, query public network IP, query whether the current network environment is
  within Chinese-GW.

## GUI

### Windows Platform

#### Windows Message Box (No CGO, dependent on system API)

- Package: github.com/skys-mission/gout/go/gui/win/mbw
- Documentation: [mbw library documentation](go/gui/win/mbw/README.md)
- Pops up a message box through the Windows API. There are four methods: simple message, simple error/warning message,
  and custom message box. Only supports Windows.
- Dependency: user32.dll (This API is usually included in a normally installed Windows system, no special settings are
  required, and this library can be used directly)

## Operating System Related

### Windows Platform

#### Query System Information (No CGO)

- Package: github.com/skys-mission/gout/go/os/win/displayw
- Documentation: [displayw library documentation](go/os/win/displayw/README.md)
- Queries display resolution, frame rate, color information, and other parameters through the Windows API. There are two
  methods: all displays and primary display. Only supports Windows.
- Dependency: user32.dll (This API is usually included in a normally installed Windows system, no special settings are
  required, and this library can be used directly)

- Package: github.com/skys-mission/gout/go/os/win/systemlw
- Documentation: [systemlw library documentation](go/os/win/systemlw/README.md)
- Queries the language used by the current operating system through the Windows API. There are two methods: returning
  LCID standard and IANA standard country codes. Only supports Windows.
- Dependency: kernel32.dll (This API is usually included in a normally installed Windows system, no special settings are
  required, and this library can be used directly)
- Note: LCID is the native return of the Windows API (hexadecimal) and will directly return the code. IANA here only
  lists language codes for a few countries. If needed, you can submit a PR or use LCID.

#### Process Management (No CGO)

- Package: github.com/skys-mission/gout/go/os/win/processw
- Features: Query all processes, set process priority and affinity, etc. I have used it in production, but I haven't
  documented it well, and now I can't understand it myself. Only supports Windows.
- Documentation to be improved
- Dependency: kernel32.dll (This API is usually included in a normally installed Windows system, no special settings are
  required, and this library can be used directly)

#### Privilege Related (No CGO)

- Package: github.com/skys-mission/gout/go/os/win/privilegew
- Features: Apply for Windows debug privileges, which are higher than administrator privileges. It first needs to work
  under administrator privileges. I have used it in production, but I haven't documented it well, and now I can't
  understand it myself. Only supports Windows.
- Documentation to be improved
- Dependency: advapi32.dll (This API is usually included in a normally installed Windows system, no special settings are
  required, and this library can be used directly)

## Compilation Related

### GO Library (No Compilation Required) and GO Tools

- Minimum Golang Version: 1.18
    - To use generics, although they are not currently used, they may be used in the future.

### CGO

Planning to develop an audio recognition tool.

## Version Compatibility

Although since version v0.2 of this project, this project will try to ensure upward compatibility as much as possible.
However, if there are security vulnerabilities, bugs, or other risks, this project will directly break upward
compatibility in a minor version (this is usually not global).

## Development Plan

### Current Plan

| Description                       | Direction           | Type                   | Schedule       |
|-----------------------------------|---------------------|------------------------|----------------|
| Get IP information via ip-api.com | GO Library          | Public Network Service | No Schedule    |
| Audio Recognition Tool            | GO Library and Tool | Local AI Function      | In Development |
