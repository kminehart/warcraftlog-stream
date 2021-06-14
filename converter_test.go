package main

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestSingleCombatLine(t *testing.T) {
	line := []byte(`6/14 09:09:29.827  DAMAGE_SHIELD_MISSED,Creature-0-4391-530-32234-16879-0000475DF3,"Starving Helboar",0x10a48,0x0,Player-4726-00A97164,"Toter-Sulfuras",0x511,0x0,33908,"Burning Spikes",0x4,RESIST,nil,0`)
	ts, err := time.Parse(WoWTimeFormat, `6/14 09:09:29.827`)
	if err != nil {
		t.Fatal(err)
	}

	expected := Entry{
		Timestamp: ts,
	}

	received, err := Convert(line)
	if err != nil {
		t.Fatal(err)
	}

	if cmp.Equal(expected, received) != true {
		t.Fatal(cmp.Diff(expected, received))
	}
}
