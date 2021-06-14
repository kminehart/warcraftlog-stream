package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const WoWTimeFormat = "1/2 15:04:05.000"

// 6/14 09:09:29.827  DAMAGE_SHIELD_MISSED,Creature-0-4391-530-32234-16879-0000475DF3,"Starving Helboar",0x10a48,0x0,Player-4726-00A97164,"Toter-Sulfuras",0x511,0x0,33908,"Burning Spikes",0x4,RESIST,nil,0
type Entry struct {
	Text      string
	Timestamp time.Time
	Event     string
}

func Convert(data []byte) (Entry, error) {
	var (
		str    = string(data)
		tokens = strings.Split(str, "  ")
		tsStr  = tokens[0]
	)

	if len(tokens) != 2 {
		return Entry{}, errors.New("invalid number of tokens")
	}

	ts, err := time.Parse(WoWTimeFormat, tsStr)
	if err != nil {
		return Entry{}, err
	}

	return Entry{
		Timestamp: ts,
		Text:      str,
	}, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// convert
		entry, err := Convert(scanner.Bytes())
		if err != nil {
			fmt.Printf("error: %s\n", err.Error())
			continue
		}

		fmt.Printf("ts: %s len: %d\n", entry.Timestamp, len(entry.Text))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
