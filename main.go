package main

import (
	_ "github.com/BurntSushi/toml"
	_ "github.com/disintegration/imaging"
	_ "github.com/gofiber/fiber/v2"
	_ "github.com/klauspost/compress"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gorm"
)

func main() {
	println("Hello, world!")
}
