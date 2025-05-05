# go-setwallpaper

A cross-platform Go library for setting desktop wallpapers with support for Linux, macOS, and Windows.

## Features

- Simple API for setting wallpapers
- Supports multiple desktop environments on Linux
- Cross-platform support (Linux, macOS, Windows)
- File type validation (JPEG, PNG, GIF, BMP)
- Error handling for common failure cases

## Installation

```bash
go get github.com/davenicholson-xyz/go-setwallpaper@latest
```
## Usage

```go
package main

import (
	"fmt"
	"github.com/davenicholson-xyz/go-setwallpaper/wallpaper"
)

func main() {
	err := wallpaper.Set("/path/to/your/wallpaper.jpg")
	if err != nil {
		fmt.Println("Error setting wallpaper:", err)
		return
	}
}
```
