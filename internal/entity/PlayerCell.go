package entity

import "log"

type PlayerCell struct {
	Cell
}

func (p *PlayerCell) Eat(prey Entity) {

}

func (p *PlayerCell) OnEaten(hunter Entity) {
	switch t := hunter.(type) {
	case *PlayerCell:
		log.Println("eaten by player", t)
	}
	log.Println(hunter.CanEat(p))
}
