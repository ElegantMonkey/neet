# neet - yet another anime library manager
[![Build Status](https://travis-ci.org/ElegantMonkey/neet.svg?branch=master)](https://travis-ci.org/ElegantMonket/neet)

## Requirements
- golang 1.8+

## Installation
- Configure $GOPATH
- Run:
```
go get github.com/ElegantMonkey/neet
```

The program will be available as `neet`.

## Configuration
The default configuration is borderline useless, so we'll need to create our own. A sample config file can be found under `config/default.yml`.

Place it in:
- *Windows:* `%APPDATA%\neet\config.yml`
- *OS X:* `~/Library/Application Support/neet/config.yml`
- *Linux/BSDs:* `~/.config/neet/config.yml`
