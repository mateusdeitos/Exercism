package house

import (
	"fmt"
	"strings"
)

type verse struct {
	phrase              string
	previousVersePrefix string
}

func Verse(v int) string {
	return buildSongRecursive(0, v, "")
}

func buildSongRecursive(index, target int, song string) string {
	if index >= target {
		song = strings.TrimRight(song, "\n")
		song = fmt.Sprintf("This is %s", song)
		return song
	}

	vs := verses()[index]
	if index > 0 {
		song = fmt.Sprintf("%s %s", vs.previousVersePrefix, song)
	}
	song = fmt.Sprintf("%s\n%s", vs.phrase, song)

	return buildSongRecursive(index+1, target, song)
}

func Song() string {
	song := ""
	vs := verses()
	for i := 0; i < len(vs); i++ {
		song = fmt.Sprintf("%s\n\n%s", song, Verse(i+1))
	}

	song = strings.Trim(song, "\n")
	song = strings.Trim(song, "\n")

	return song
}

func verses() []verse {
	return []verse{
		{"the house that Jack built.", ""},
		{"the malt", "that lay in"},
		{"the rat", "that ate"},
		{"the cat", "that killed"},
		{"the dog", "that worried"},
		{"the cow with the crumpled horn", "that tossed"},
		{"the maiden all forlorn", "that milked"},
		{"the man all tattered and torn", "that kissed"},
		{"the priest all shaven and shorn", "that married"},
		{"the rooster that crowed in the morn", "that woke"},
		{"the farmer sowing his corn", "that kept"},
		{"the horse and the hound and the horn", "that belonged to"},
	}
}
