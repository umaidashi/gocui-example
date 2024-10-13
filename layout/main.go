// Copyright 2014 The gocui Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	menu := []string{"one", "two", "three"}
	if v, err := g.SetView("side", 0, 0, int(0.2*float32(maxX)), maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Side"
		for _, item := range menu {
			fmt.Fprintln(v, item)
		}
	}
	if v, err := g.SetView("main", int(0.2*float32(maxX))+1, 0, maxX, maxY-6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Main"
	}
	if v, err := g.SetView("cmdline", 0, maxY-5, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Cmdline"
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
