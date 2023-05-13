package main

import (
	_ "github.com/disintegration/imaging"
	_ "github.com/gofiber/fiber/v2"

	_ "skyfall/services/config"
	_ "skyfall/services/database"
)

func main() {
	println("Hello, world!")
}
