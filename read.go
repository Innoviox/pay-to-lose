package main

import (
	"fmt"
	"os"
	// "log"

	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	common "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)

type AggregatedData struct {
	Deaths map[KnifeType]int
	Owners map[KnifeType]int
}

func NewAggregatedData() *AggregatedData {
	d := &AggregatedData{
		Deaths: make(map[KnifeType]int),
		Owners: make(map[KnifeType]int),
	}

	return d
}

func (ad *AggregatedData) Add (d *Data) {
	for _, v := range d.Owners {
		ad.Owners[v]++
	}

	for k, v := range d.Deaths {
		ad.Deaths[d.Owners[k]] += v
	}
}

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


func Read(fp string, c chan *Data) {
	d := NewData()
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
		
		var vfp = e.Victim.String() + fp
		if _, ok := d.Owners[vfp]; !ok {
			var vkt = GetKnife(e.Victim)
			if vkt != 0 {
				d.Owners[e.Victim.String() + "@@@" + fp] = (KnifeType)(vkt)
			}
		}

		var kfp = e.Killer.String() + fp
		if _, ok := d.Owners[kfp]; !ok {
			var kkt = GetKnife(e.Killer)
			if kkt != 0 {
				d.Owners[e.Killer.String() + "@@@" + fp] = (KnifeType)(kkt)
			}
		}
		
		// try to rule out knife rounds
		// todo: were they inspecting
		// todo: do they have a different gun out but it's not ready to fire since they just switched
		if weapon != nil && weapon.Type == common.EqKnife && len(e.Victim.Weapons()) > 1 {
			d.Deaths[vfp]++
		}
	})


	err = p.ParseToEnd()
	if err != nil {
		fmt.Printf("failed to parse demo: %s", err)
	}

	c <- d
}
