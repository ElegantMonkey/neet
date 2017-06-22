# neet - yet another cli anime library manager
[![Build Status](https://travis-ci.org/ElegantMonkey/neet.svg?branch=master)](https://travis-ci.org/ElegantMonket/neet)

## Requirements
- golang 1.8+

## Installation
- [Ensure your Go install is properly configured](https://golang.org/doc/install#install)
- Run:
```
go get github.com/ElegantMonkey/neet
```

The program will be available as `neet`.

## Configuration
The default configuration is borderline useless, so we'll need to create our own. A sample config file can be found under `config/default.yml`.

Place it in:
- **Windows:** `%APPDATA%\neet\config.yml`
- **OS X:** `~/Library/Application Support/neet/config.yml`
- **Linux/BSDs:** `~/.config/neet/config.yml`

## Usage
### Commands
- **add:** Add an anime. Current episode is optional. `neet add "anime name" /home/user/anime_folder/ [current_episode]`
- **list:** List animes, and progress in the series. `neet list`
- **play:** Play an anime. `play current` plays the current episode; `play next` plays the next episode, and keeps playing after the episode ends. Defaults to `play next`. `neet play "anime name" next` 
- **set:** Changes anime configs. `neet set "anime name" [field] [value]`, where field can be either `name` or `episode`.
