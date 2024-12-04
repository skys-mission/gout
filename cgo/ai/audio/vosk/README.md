# gout-vosk

Why not use the native vosk-api? Because I couldn't compile the official example on Windows, and I also need a binary vosk tool. I don't want to write C, so this project was born.

Other languages: [简体中文](README_zh.md), (Currently unable to translate more)

<!-- TOC -->
* [gout-vosk](#gout-vosk)
  * [How to Compile the Binary Tool](#how-to-compile-the-binary-tool)
  * [Binary Tool Usage](#binary-tool-usage)
  * [Library Usage](#library-usage)
  * [Support](#support)
  * [Vosk-API Description](#vosk-api-description)
<!-- TOC -->

## How to Compile the Binary Tool

You can download the binary file from the github release first. Compilation requires waiting for the DOCKERFILE I will prepare later.

## Binary Tool Usage

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

## Library Usage

Unless you intrude into this library repository and place the vosk-api dynamic link library in the lib folder, you cannot successfully compile this library. I currently haven't found a good way to handle this, unless I upload the .dll file to github, but that's not what I want. I need a more friendly way.

## Support

Currently, only Windows is supported. macOS and Linux support are planned (no schedule yet).

## Vosk-API Description

The vosk-api related files are distributed under the Apache License Version 2.0.

vosk-api project address: https://github.com/alphacep/vosk-api

vosk model download page: https://alphacephei.com/vosk/models

---