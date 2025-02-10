package main

import (
	"fmt"
	"os"
	// "log"

	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	common "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type Data struct {
	Deaths map[string]int
	Owners map[string]KnifeType
}

func NewData() *Data {
	d := &Data{
		Deaths: make(map[string]int),
		Owners: make(map[string]KnifeType),
	}

	return d
}


func (d *Data) Read(fp string) {
	f, err := os.Open(fp)
	if err != nil {
		fmt.Printf("failed to open demo file: %s", fp)
	}
	defer f.Close()
	defer func() {
        if recover() != nil {
            fmt.Printf("failed to parse demo: %s", fp)
        }
    }()


	p := dem.NewParser(f)
	defer p.Close()

	p.RegisterEventHandler(func(e events.Kill) {
		var weapon = e.Victim.ActiveWeapon()

		var vkt = GetKnife(e.Victim)
		var vfp = e.Victim.String() + fp
		if _, ok := d.Owners[vfp]; !ok && vkt != 0 {
			d.Owners[e.Victim.String() + "@@@" + fp] = (KnifeType)(vkt)
		}

		var kkt = GetKnife(e.Killer)
		var kfp = e.Killer.String() + fp
		if _, ok := d.Owners[kfp]; !ok && kkt != 0 {
			d.Owners[e.Killer.String() + "@@@" + fp] = (KnifeType)(kkt)
		}
		
		// try to rule out knife rounds
		if weapon != nil && weapon.Type == common.EqKnife && len(e.Victim.Weapons()) > 1 {
			d.Deaths[vfp]++
		}
	})


	err = p.ParseToEnd()
	if err != nil {
		fmt.Printf("failed to parse demo: %s", err)
	}
}
