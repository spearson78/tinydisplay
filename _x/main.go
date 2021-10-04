package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c := 0x0020
	buf := ""
	// b * d = 11 * 13
	switch c {
	case 0x103:
		buf = `
			00001111111
			00010000001
			00100000001
			01001000101
			10000101001
			10000010001
			10000101001
			01001000101
			00100000001
			00010000001
			00001111111
			00000000000
			00000000000
		`
	case 0x102:
		buf = `
			00000000000
			00000000001
			00000010001
			00000011001
			00000011101
			11111111111
			00000011101
			00000011001
			00000010001
			00000000001
			00000000000
			00000000000
			00000000000
		`
	case 0x101:
		buf = `
			00000000000
			00000000000
			00000000100
			00000000100
			00000000100
			00100000100
			01000000100
			11111111100
			01000000000
			00100000000
			00000000000
			00000000000
			00000000000
		`
	case 0x109:
		buf = `
			00000000000
			00000100000
			00001110000
			00010101000
			00100100100
			00000100000
			00000100000
			00000100000
			00000100000
			00000100000
			00000100000
			00000100000
			00000100000
		`
	case 0x108:
		buf = `
			00000000000
			00000100000
			00000100000
			00000100000
			00000100000
			00000100000
			00000100000
			00000100000
			00100100100
			00010101000
			00001110000
			00000100000
			00000000000
		`
	case 0x106:
		buf = `
			00000000000
			00000000000
			00000000000
			00000001000
			00000000100
			00000000010
			11111111111
			00000000010
			00000000100
			00000001000
			00000000000
			00000000000
			00000000000
		`
	case 0x107:
		buf = `
			00000000000
			00000000000
			00000000000
			00010000000
			00100000000
			01000000000
			11111111111
			01000000000
			00100000000
			00010000000
			00000000000
			00000000000
			00000000000
		`

	case 0x01FF:
		buf = `
			00000000000
			00000000000
			11111111111
			01111111110
			01111111110
			00111111100
			00111111100
			00011111000
			00011111000
			00001110000
			00000100000
			00000100000
			00000000000
		`

	case 0x0020:
		buf = `
			00000000000
			00000000000
			00000000000
			00000000000
			00000000000
			00000000000
			00000000000
			00000000000
			00000000000
			00000000000
			10000000001
			10000000001
			11111111111
		`

	}

	buf = strings.ReplaceAll(buf, "\n", "")
	buf = strings.ReplaceAll(buf, "\t", "")
	//fmt.Printf("%s\n", buf)

	fontBuf := []uint8{}
	tmp := uint8(0)
	for i, b := range buf {
		if b == '1' {
			tmp |= 1 << (8 - i%8 - 1)
		}
		if i%8 == 7 || i == len(buf)-1 {
			fontBuf = append(fontBuf, tmp)
			tmp = 0
		}
	}

	strBuf := ""
	for i, b := range fontBuf {
		if i == 0 {
			strBuf += fmt.Sprintf("0x%02X", b)
		} else {
			strBuf += fmt.Sprintf(", 0x%02X", b)
		}
	}
	fmt.Printf("	/* 0x%03X */ tinyfont.Glyph{Rune: 0x%03X, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{%s}},\n", c, c, strBuf)

	return nil
}
