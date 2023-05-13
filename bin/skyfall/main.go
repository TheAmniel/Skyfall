package main

import (
	_ "github.com/disintegration/imaging"
	_ "github.com/gofiber/fiber/v2"

	_ "skyfall/services/config"
	_ "skyfall/services/database"
	"skyfall/utils"
)

func main() {
	print("Name: "+utils.AppName+"\n",
		"Version: "+utils.Version+"\n",
		"Branch: "+utils.Branch+" ["+utils.Commit+"]\n",
		"Config file: "+utils.ConfigFile+"\n",
		"Built at: "+utils.BuiltAt+"\n",
	)
}
