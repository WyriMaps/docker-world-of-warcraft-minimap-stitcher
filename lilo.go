package main

import (
	"os"
	"runtime"

	stitcher "github.com/WyriMaps/MinimapStitcher"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "sitcher"
	app.Usage = "Sitcher World of Warcraft minimaps together"
	app.Action = func(c *cli.Context) {
		_, errTiles := os.Stat("/opt/map_tiles/")
		if os.IsNotExist(errTiles) {
			return
		}

		_, errMaps := os.Stat("/opt/maps/")
		if os.IsNotExist(errMaps) {
			return
		}

		runtime.GOMAXPROCS(runtime.NumCPU()) // Auto detects the number of cores to use
		stitcher.Stitch("/opt/map_tiles/", "/opt/maps/");
	}

	app.Run(os.Args)
}
