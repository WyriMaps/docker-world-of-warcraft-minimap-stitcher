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
		runtime.GOMAXPROCS(runtime.GOMAXPROCS(runtime.NumCPU())) // Auto detects the number of cores to use
		s := []string{"Found", runtime.NumCPU(), "CPU cores we'll be using"};
		println(strings.Join(s, " "));
		var report = func(message map[string]string) {
			println(message["minimap"] + ": " + message["type"]);
		}
		stitcher.Stitch(report, "/opt/map_tiles/", "/opt/maps/");
		println("Done!");
	}

	app.Run(os.Args)
}
