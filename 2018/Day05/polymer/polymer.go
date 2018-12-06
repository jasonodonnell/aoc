package polymer

import (
	"fmt"
	"math"
	"strings"
)

type Polymer struct {
	Unit string
}

func (p *Polymer) Strip(unit int) {
	remove := fmt.Sprintf("%s", string(byte(unit)))
	p.Unit = strings.Replace(p.Unit, remove, "", -1)
	remove = fmt.Sprintf("%s", string(byte(unit+32)))
	p.Unit = strings.Replace(p.Unit, remove, "", -1)
}

func (p *Polymer) React() {
Loop:
	for {
		for i := 0; i < len(p.Unit)-1; i++ {
			if math.Abs(float64(p.Unit[i])-float64(p.Unit[i+1])) == 32 {
				remove := fmt.Sprintf("%s%s", string(p.Unit[i]), string(p.Unit[i+1]))
				p.Unit = strings.Replace(p.Unit, remove, "", 1)
				goto Loop
			}
		}
		break
	}
}

func New(unit string) *Polymer {
	return &Polymer{
		Unit: unit,
	}
}
