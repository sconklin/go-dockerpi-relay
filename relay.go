package relay

// Interface to the dockerpi I2C-controlled relay board

import (
	"errors"

	"github.com/sconklin/go-i2c"
)

type Relay struct {
	i2c *i2c.I2C
}

func NewRelay(i2c *i2c.I2C) (*Relay, error) {
	// TODO
	thisrelay := &Relay{i2c: i2c}
	return thisrelay, nil
}

func (thisrelay *Relay) On(rnum uint8) error {
	if (rnum > 4) || (rnum < 1) {
		return errors.New("Relay number must be in range 1-4")
	}
	logrelay.Debug("Turning ON relay %d\n", rnum)
	return thisrelay.i2c.WriteRegU8(rnum, 1)
}

func (thisrelay *Relay) Off(rnum uint8) error {
	if (rnum > 4) || (rnum < 1) {
		return errors.New("Relay number must be in range 1-4")
	}
	logrelay.Debug("Turning OFF relay %d\n", rnum)
	return thisrelay.i2c.WriteRegU8(rnum, 0)
}
