package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	chars = ` !"#$%&()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_abcdefghijklmnopqrstuvwxyz{|}~`
)

func main() {
	length := flag.Int("l", 8, fmt.Sprintf("password length"))
	flag.Parse()

	runes := []rune(chars)

	rand.Seed(time.Now().UnixNano())

	var pwd strings.Builder
	for i := 0; i < *length; i++ {
		pwd.WriteRune(runes[rand.Intn(len(runes))])
	}
	print := color.New(color.FgRed, color.BgBlack).PrintlnFunc()
	print(pwd.String())
}
