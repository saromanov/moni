package main

import
(
	"fmt"
	"os"
	"github.com/codegangsta/cli"
)

func start(path string) {
	moni := main.New(path)

}

func main() {
  app := cli.NewApp()
  app.Name = "start"
  app.Usage = "Start monitoring of the systems"
  app.Flags = []cli.Flag {
  cli.StringFlag {
    Name: "start",
    Usage: "Start of monitoring",
  },

  cli.StringFlag {
  	Name: "config",
  	Value: "",
  	Usage: "Path to the config file",
  },

  cli.StringFlag {
  	Name: "stats",
  	Value: "",
  	Usage: "Return status of target machine",
  },

}
  app.Action = func(c *cli.Context) {
    start := c.String("start")
    fmt.Println(start)
    conf := c.String("config")
    fmt.Println(conf)
    status := c.String("status")
    fmt.Println(status)
  }

  app.Run(os.Args)
}