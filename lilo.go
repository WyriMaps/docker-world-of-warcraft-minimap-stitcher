package main

import (
	"os"
	"runtime"
	"strconv"
	"strings"

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

		runtime.GOMAXPROCS(runtime.GOMAXPROCS(runtime.NumCPU())) // Auto detects the number of cores to use
		s := []string{"Found", strconv.Itoa(runtime.NumCPU()), "CPU cores we'll be using"};
		println(strings.Join(s, " "));
		stitcher.Stitch("/opt/map_tiles/", "/opt/maps/");
		println("Done!");
	}

	app.Run(os.Args)
}
