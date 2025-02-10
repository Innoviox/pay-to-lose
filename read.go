package main

import (
	"fmt"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	common "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type Data struct {
	Deaths map[KnifeType]int
	Owners map[KnifeType]map[string]int
}

func NewData() *Data {
	return &Data{
		Deaths: make(map[KnifeType]int),
		Owners: make(map[KnifeType]map[string]int),
	}
}


func (d *Data) Read(fp string) {
	f, err := os.Open(fp)
	if err != nil {
		fmt.Printf("failed to open demo file: %s", fp)
	}
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()

	p.RegisterEventHandler(func(e events.Kill) {
		var weapon = e.Victim.ActiveWeapon()
		if weapon != nil && weapon.Type == common.EqKnife {
			var knifeType = (KnifeType)(weapon.Entity.Property("m_iItemDefinitionIndex").Value().S2UInt64())

			d.Deaths[knifeType]++
			if d.Owners[knifeType] == nil {
				d.Owners[knifeType] = make(map[string]int)
			}

			d.Owners[knifeType][e.Victim.String()] = 1
			d.Owners[knifeType][e.Killer.String()] = 1
		}
	})

	err = p.ParseToEnd()
	if err != nil {
		fmt.Printf("failed to parse demo: %s", fp)
	}
}

func (d *Data) GetDeaths() map[string]int {
	var ret = make(map[string]int)
	for kt, deaths := range d.Deaths {
		ret[ToString(kt)] = deaths
	}
	return ret
}

func (d *Data) GetOwners() map[string]int {
	var ret = make(map[string]int)
	for kt, owners := range d.Owners {
		ret[ToString(kt)] = len(owners)
	}
	return ret
}