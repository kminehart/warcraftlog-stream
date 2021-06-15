package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const WoWTimeFormat = "1/2 15:04:05.000"

// 6/14 09:09:29.827  DAMAGE_SHIELD_MISSED,Creature-0-4391-530-32234-16879-0000475DF3,"Starving Helboar",0x10a48,0x0,Player-4726-00A97164,"Toter-Sulfuras",0x511,0x0,33908,"Burning Spikes",0x4,RESIST,nil,0
// 6/14 14:56:40.406  SPELL_DAMAGE,Player-4726-00A97164,"Toter-Sulfuras",0x511,0x0,Creature-0-4391-530-32234-16867-000047AB7A,"Shattered Hand Grunt",0x10a48,0x0,30022,"Devastate",0x1,Creature-0-4391-530-32234-16867-000047AB7A,0000000000000000,27,100,0,0,0,-1,0,0,0,-317.82,2949.87,1944,0.9431,63,568,354,-1,1,0,0,0,1,nil,nil
type Entry struct {
	Text      string    `json:"-"`
	Timestamp time.Time `json:"timestamp"`
	Event     Fielder
}

func Convert(data []byte) (Entry, error) {
	var (
		str     = string(data)
		tokens  = strings.Split(str, "  ")
		tsStr   = tokens[0]
		entries = strings.Split(tokens[1], ",")
	)

	if len(tokens) != 2 {
		return Entry{}, errors.New("invalid number of tokens")
	}

	ts, err := time.Parse(WoWTimeFormat, tsStr)
	if err != nil {
		return Entry{}, err
	}

	var val suffixParser
	for k, v := range suffixMap {
		if strings.HasSuffix(entries[0], k) {
			val = v
			break
		}
	}

	if val == nil {
		return Entry{}, fmt.Errorf("%s was not found", entries[0])
	}
	suf, err := val(entries[1:])
	if err != nil {
		return Entry{}, err
	}

	//event, err := json.Marshal(suf)
	//if err != nil {
	//	return Entry{}, err
	//}

	return Entry{
		Timestamp: ts,
		Text:      str,
		Event:     suf,
	}, nil
}

func main() {
	// lokiURL, ok := os.LookupEnv("LOKI_URL")
	// if !ok {
	// 	lokiURL = "http://localhost:3100"
	// }

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// convert
		entry, err := Convert(scanner.Bytes())
		if err != nil {
			//fmt.Printf("error: %s\n", err.Error())
			continue
		}

		if err := json.NewEncoder(os.Stdout).Encode(entry); err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
