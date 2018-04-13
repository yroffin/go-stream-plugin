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
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/yroffin/go-stream-plugin/commons"
)

// Plugin instance
type Plugin struct {
	c chan string
}

// Instance instance
func Instance() commons.IPlugin {
	instance := Plugin{}
	return &instance
}

// GenUUID uuid
func GenUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	return hex.EncodeToString(uuid), nil
}

func (p *Plugin) send() {
	for {
		time.Sleep(1 * time.Second)
		g, _ := GenUUID()
		p.c <- g
	}
}

// Start the plugin
func (p *Plugin) Start() {
	log.Println("Start plugin")

	p.c = make(chan string)
	go p.send()
}

// Chan access
func (p *Plugin) Chan() chan string {
	return p.c
}
