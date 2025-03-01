# gout

[![CodeQL Advanced](https://github.com/skys-mission/gout/actions/workflows/codeql.yml/badge.svg)](https://github.com/skys-mission/gout/actions/workflows/codeql.yml)
[![Codacy Security Scan](https://github.com/skys-mission/gout/actions/workflows/codacy.yml/badge.svg)](https://github.com/skys-mission/gout/actions/workflows/codacy.yml)

I noticed that disorganized projects (although this project has a clear directory structure, it covers too many domains) are not conducive to project development, so this repository has been deprecated.

Since I found that my third-party blog has a certain number of bookmarks for this project, I'm keeping this repository available for reference. If you encounter any bugs, please raise an ISSUE, and I will still fix them.

Other languages: [简体中文](README_zh.md), (Currently unable to translate more)

<!-- TOC -->
* [gout](#gout)
  * [How to Use](#how-to-use)
  * [AI](#ai)
    * [Local AI](#local-ai)
      * [Audio Processing](#audio-processing)
        * [Vosk (CGO)](#vosk-cgo)
  * [Public Network Services](#public-network-services)
    * [IP Information Retrieval](#ip-information-retrieval)
      * [iplocation.net](#iplocationnet)
  * [GUI](#gui)
    * [Windows Platform](#windows-platform)
      * [Windows Message Box (Without CGO, dependent on system API)](#windows-message-box-without-cgo-dependent-on-system-api)
  * [Operating System Related](#operating-system-related)
    * [Windows Platform](#windows-platform-1)
      * [Query System Information (Without CGO)](#query-system-information-without-cgo)
      * [Process Management (Without CGO)](#process-management-without-cgo)
      * [Privilege Related (Without CGO)](#privilege-related-without-cgo)
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

## AI

### Local AI

#### Audio Processing

##### Vosk (CGO)

- Package: github.com\skys-mission\gout\cgo\ai\audio\vosk
    - Function: Call the Vosk engine for audio-to-text conversion using a local model. Library support is temporarily
      unavailable (source code can also be used). Binary files and packaged libraries are available in the github
      release.
    - Usage documentation: [gout-vosk documentation](cgo/ai/audio/vosk/README.md)

## Public Network Services

**You should follow the terms of service of the network service providers.
This project only provides Go language code to call network services. The network services themselves are unrelated to
this project. Please use them according to the official terms of the corresponding services.**

### IP Information Retrieval

#### iplocation.net

- Package: github.com/skys-mission/gout/go/pubnet/iplocation
    - Function: Query specified IP information, query the public network IP of the exit, and query whether the current
      network environment is within the Chinese-GW.
    - Usage documentation: [iplocation library documentation](go/pubnet/iplocation/README.md)

## GUI

### Windows Platform

#### Windows Message Box (Without CGO, dependent on system API)

- Package: github.com/skys-mission/gout/go/gui/win/mbw
    - Pop up a message box through the Windows API. There are four methods: simple message, simple error/warning
      message, and custom message box. Only supports Windows.
    - Usage documentation: [mbw library documentation](go/gui/win/mbw/README.md)
    - Dependency: user32.dll (Generally, a properly installed Windows system includes this API, no special settings are
      needed, and this library can be used directly)

## Operating System Related

### Windows Platform

#### Query System Information (Without CGO)

- Package: github.com/skys-mission/gout/go/os/win/displayw
    - Query display resolution, frame rate, color information, and other parameters through the Windows API. There are
      two methods: all displays and the main display. Only supports Windows.
    - Usage documentation: [displayw library documentation](go/os/win/displayw/README.md)
    - Dependency: user32.dll (Generally, a properly installed Windows system includes this API, no special settings are
      needed, and this library can be used directly)

- Package: github.com/skys-mission/gout/go/os/win/systemlw
    - Query the language used by the current operating system through the Windows API. There are two methods: returning
      the LCID standard and the IANA standard country code. Only supports Windows.
    - Usage documentation: [systemlw library documentation](go/os/win/systemlw/README.md)
    - Dependency: kernel32.dll (Generally, a properly installed Windows system includes this API, no special settings
      are needed, and this library can be used directly)
    - Note: LCID is the native return of the Windows API (hexadecimal) and will directly return the code. The IANA
      standard here only includes language codes for a few countries. If needed, you can submit a PR or use LCID.

#### Process Management (Without CGO)

- Package: github.com/skys-mission/gout/go/os/win/processw
    - Function: Query all processes, set process priority and affinity, etc. I have used it in production, but the
      documentation has not been perfected. Now I am a bit confused about it myself. Only supports Windows.
    - Documentation to be improved
    - Dependency: kernel32.dll (Generally, a properly installed Windows system includes this API, no special settings
      are needed, and this library can be used directly)

#### Privilege Related (Without CGO)

- Package: github.com/skys-mission/gout/go/os/win/privilegew
    - Function: Apply for Windows debugging privileges, which are higher than administrator privileges. First, it needs
      to work under administrator privileges. I have used it in production, but the documentation has not been
      perfected. Now I am a bit confused about it myself. Only supports Windows.
    - Documentation to be improved
    - Dependency: advapi32.dll (Generally, a properly installed Windows system includes this API, no special settings
      are needed, and this library can be used directly)

## Compilation Related

### GO Library (No Compilation Required) and GO Tools

- Minimum Golang version: 1.18
    - To use generics, although they are not currently used, they may be used in the future.

### CGO

**See the CGO project documentation for details, as the compilation methods vary due to the CGO toolchain.**

## Version Compatibility

Although since version v0.2 of this project, this project will strive to ensure upward compatibility as much as
possible. However, if there are security vulnerabilities, bugs, or other risks, this project will directly break upward
compatibility in a minor version (this is usually not global).

## Development Plan

### Current Plan

| Description                            | Direction           | Type                   | Schedule                                                 |
|----------------------------------------|---------------------|------------------------|----------------------------------------------------------|
| Retrieve IP information via ip-api.com | GO Library          | Public Network Service | No schedule                                              |
| Improve cgo-vosk library               | GO Library and Tool | Local AI Function      | Audio-to-text completed, improving functionality planned |