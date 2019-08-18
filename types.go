package main

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/hypebeast/go-osc/osc"
)

type config struct {
	Server  string `short:"s" long:"server" description:"Name or IP adress to connect to (required)"`
	Port    int    `short:"p" long:"port" description:"Port on which the server is listening (required)"`
	Color   string `short:"c" long:"color" description:"Color to use (ex. #a2ff13)"`
	Mode    int    `short:"m" long:"mode" description:"Mode to use (0-15)" default:"-1"`
	Version bool   `short:"v" long:"version" description:"Show the current version"`
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

func (c client) sendMode(mode int) error {

	msg := osc.NewMessage(moddAddr.String(), float32(mode))

	if err := c.oscClient.Send(msg); err != nil {
		return err
	}

	return nil
}

type color struct {
	red   float32
	green float32
	blue  float32
}

var predefinedColors = map[string]color{
	"white": color{0xff, 0xff, 0xff},
	"black": color{0x00, 0x00, 0x00},
	"red":   color{0xff, 0x00, 0x00},
	"blue":  color{0x00, 0x00, 0xff},
	"green": color{0x00, 0xff, 0x00},
}

func parseColor(str string) (col color, err error) {
	if strings.HasPrefix(str, "#") {
		str = strings.ToLower(str[1:])

		if len(str) == 3 {
			var s string
			for i := 0; i < 3; i++ {
				s += strings.Repeat(string(str[i]), 2)
			}
			str = s
		}

		var ok bool
		if col.red, col.green, col.blue, ok = convertColor(str); ok {
			return col, nil
		}
		return color{}, fmt.Errorf("invalid color definition: %v", str)
	}

	if col, ok := predefinedColors[str]; ok {
		return col, nil
	}
	return color{}, fmt.Errorf("invalid predefined color value: %v", str)
}

func convertColor(str string) (r, g, b float32, ok bool) {
	if len(str) != 6 {
		return r, g, b, false
	}

	if decoded, err := hex.DecodeString(str); err == nil {
		return float32(decoded[0]), float32(decoded[1]), float32(decoded[2]), true
	}
	return r, g, b, false
}
