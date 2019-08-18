package main

import (
	"fmt"
	"os"

	"github.com/hypebeast/go-osc/osc"
	"github.com/jessevdk/go-flags"
)

const version = "1.0.0"

var cfg config
var parser = flags.NewParser(&cfg, flags.HelpFlag|flags.PassDoubleDash)

func main() {
	_, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfg.Version {
		fmt.Println("version:", version)
		os.Exit(0)
	}

	if cfg.Server == "" || cfg.Port == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	client := newClient(cfg.Server, cfg.Port)

	if cfg.Color != "" {
		if c, err := parseColor(cfg.Color); err != nil {
			fmt.Println("failed to parse color:", cfg.Color, " err:", err)
			os.Exit(1)
		} else {
			fmt.Println("sending color", cfg.Color, "to device...")
			if err := client.sendColor(c); err != nil {
				handleError("color "+cfg.Color, err)
			}
			fmt.Println("done")
		}
	}

	if cfg.Mode >= 0 && cfg.Mode < 15 {
		fmt.Println("sending mode ", cfg.Mode, " to device...")
		if err := client.sendMode(cfg.Mode); err != nil {
			handleError("mode", err)
		}
		fmt.Println("done")
	}
}

func newClient(addr string, port int) client {
	return client{
		oscClient: osc.NewClient(addr, port),
	}
}

func handleError(action string, err error) {
	fmt.Printf("failed to send %v to device %v:%v err: %v", action, cfg.Server, cfg.Port, err)
	os.Exit(1)
}
