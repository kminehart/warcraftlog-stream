package main

import (
	"fmt"
	"strconv"
)

type suffixParser func([]string) (Fielder, error)

type Fielder interface {
	String() string
}

var suffixMap = map[string]suffixParser{
	"_DAMAGE": DamageParser,
}

//https://github.com/magey/classic-warrior/wiki/Combat-log-format
type Damage struct {
	SourceGUID      string  `json:"source_guid,omitempty"`
	SourceName      string  `json:"source_name,omitempty"`
	SourceFlags     int     `json:"source_flags,omitempty"`
	SourceRaidFlags int     `json:"source_raid_flags,omitempty"`
	DestGUID        string  `json:"dest_guid,omitempty"`
	DestName        string  `json:"dest_name,omitempty"`
	DestFlags       int     `json:"dest_flags,omitempty"`
	DestRaidFlags   int     `json:"dest_raid_flags,omitempty"`
	InfoGUID        string  `json:"info_guid,omitempty"`
	OwnerGUID       string  `json:"owner_guid,omitempty"`
	CurrentHP       int     `json:"current_hp,omitempty"`
	MaxHP           int     `json:"max_hp,omitempty"`
	AttackPower     int     `json:"attack_power,omitempty"`
	SpellPower      int     `json:"spell_power,omitempty"`
	Armor           int     `json:"armor,omitempty"`
	PowerType       int     `json:"power_type,omitempty"`
	CurrentPower    int     `json:"current_power,omitempty"`
	MaxPower        int     `json:"max_power,omitempty"`
	PowerCost       int     `json:"power_cost,omitempty"`
	PositionX       float64 `json:"position_x,omitempty"`
	PositionY       float64 `json:"position_y,omitempty"`
	UIMapID         int     `json:"ui_map_id,omitempty"`
	Facing          float64 `json:"facing,omitempty"`
	Level           int     `json:"level,omitempty"`
	Amount          int     `json:"amount,omitempty"`
	RawAmount       int     `json:"raw_amount,omitempty"`
	Overkill        bool    `json:"overkill,omitempty"`
	School          string  `json:"school,omitempty"`
	Resisted        bool    `json:"resisted,omitempty"`
	Blocked         bool    `json:"blocked,omitempty"`
	Absorbed        bool    `json:"absorbed,omitempty"`
	Critical        bool    `json:"critical,omitempty"`
	Glancing        bool    `json:"glancing,omitempty"`
	Crushing        bool    `json:"crushing,omitempty"`
}

func (d *Damage) String() string {
	return fmt.Sprintf(`{ "source_guid": "%s", "source_name": "%s", "dest_guid": "%s", "dest_name": "%s", "amount": %d  }`, d.SourceGUID, d.SourceName, d.DestGUID, d.DestName, d.Amount)
}

func DamageParser(fields []string) (Fielder, error) {
	amount, err := strconv.Atoi(fields[24])
	if err != nil {
		return nil, err
	}

	return &Damage{
		SourceGUID: fields[0],
		SourceName: fields[1],
		DestGUID:   fields[4],
		DestName:   fields[5],
		Amount:     amount,
	}, nil
}
