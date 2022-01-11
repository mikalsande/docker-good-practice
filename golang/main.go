package main

import (
	"fmt"

	"github.com/enescakir/emoji"
)

func main() {
	fmt.Printf("Hello %v\n", emoji.WavingHand)
	fmt.Printf("I am %v from %v\n",
		emoji.ManTechnologist,
		emoji.FlagForTurkey,
	)
	fmt.Printf("Different skin tones.\n  default: %v light: %v dark: %v\n",
		emoji.ThumbsUp,
		emoji.OkHand.Tone(emoji.Light),
		emoji.CallMeHand.Tone(emoji.Dark),
	)
	fmt.Printf("Emojis with multiple skin tones.\n  both medium: %v light and dark: %v\n",
		emoji.PeopleHoldingHands.Tone(emoji.Medium),
		emoji.PeopleHoldingHands.Tone(emoji.Light, emoji.Dark),
	)
	fmt.Println(emoji.Parse("Emoji aliases are :sunglasses:"))
	emoji.Println("Use fmt wrappers :+1: with emoji support :tada:")
}
