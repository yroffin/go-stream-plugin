// MIT License
//
// Copyright (c) 2017 yroffin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	"log"
	"plugin"
	"reflect"

	"github.com/yroffin/go-stream-plugin/commons"
)

// Lookup for instance plugin
func Lookup(plug string) (commons.IPlugin, error) {
	so, err := plugin.Open(plug + ".so")
	if err != nil {
		return nil, err
	}
	sym, err := so.Lookup("Instance")
	if err != nil {
		panic(err)
	}

	log.Println("Plugin", sym)
	plugin, errPlug := sym.(func() commons.IPlugin)
	if !errPlug {
		panic("Error")
	}

	return plugin(), nil
}

func receive(p commons.IPlugin) {
	for {
		log.Println("REC")
		value := <-p.Chan()
		log.Println("REC", value, reflect.TypeOf(value))
	}
}

// Main
func main() {
	p, err := Lookup("plugin-producer")
	if err != nil {
		panic(err)
	}
	p.Start()
	receive(p)
}
