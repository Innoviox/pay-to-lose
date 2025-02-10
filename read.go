package main

import (
	"fmt"
	"log"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/events"
	common "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
)


func read(fp string) {
	f, err := os.Open(fp)
	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()

	p.RegisterEventHandler(func(e events.Kill) {
		var weapon = e.Victim.ActiveWeapon()
		if weapon != nil && weapon.Type == common.EqKnife {
			var knifeType = (KnifeType)(weapon.Entity.Property("m_iItemDefinitionIndex").Value().S2UInt64())
			fmt.Printf("%s died holding %s\n", e.Victim, ToString(knifeType))
		}
	})

	// Parse to end
	err = p.ParseToEnd()
	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}
}