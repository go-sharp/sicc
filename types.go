package main

import (
	"fmt"
	"strings"

	"github.com/hypebeast/go-osc/osc"
)

type config struct {
	Server string `short:"s" long:"server" description:"Name or IP adress to connect to" required:"true"`
	Port   int    `short:"p" long:"port" description:"Port on which the server is listening" required:"true"`
	Color  string `short:"c" long:"color" description:"Color to use (ex. #a2ff13)"`
	Mode   int
}

type oscAddress string

func (o oscAddress) String() string {
	return string(o)
}

const (
	knbrAddr  oscAddress = "/knbr"
	knbgAddr  oscAddress = "/knbg"
	knbbAddr  oscAddress = "/knbb"
	delayAddr oscAddress = "/delay"
	audioAddr oscAddress = "/audio"
	modeAddr  oscAddress = "/mode"
	moddAddr  oscAddress = "/modd"
	cfgAddr   oscAddress = "/cfg"
)

type client struct {
	oscClient *osc.Client
}

func (c client) sendColor(color color) error {

	msgs := []*osc.Message{
		osc.NewMessage(knbrAddr.String(), color.red),
		osc.NewMessage(knbgAddr.String(), color.green),
		osc.NewMessage(knbbAddr.String(), color.blue),
	}

	for _, m := range msgs {
		if err := c.oscClient.Send(m); err != nil {
			return err
		}
	}

	return nil
}

type color struct {
	red   uint32
	green uint32
	blue  uint32
}

var predefinedColors = map[string]color{
	"white": color{0xff, 0xff, 0xff},
	"black": color{0x00, 0x00, 0x00},
	"red":   color{0xff, 0x00, 0x00},
	"blue":  color{0x00, 0x00, 0xff},
	"green": color{0x00, 0xff, 0x00},
}

func parseColor(str string) (color, error) {
	if strings.HasPrefix(str, "#") {
		str = str[1:]
		if len(str) == 3 {

		} else {

		}
	} else {
		if col, ok := predefinedColors[str]; ok {
			return col, nil
		}
		return color{}, fmt.Errorf("invalid predefined color value: %v", str)
	}
}
