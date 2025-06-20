package models

import "sync"

type Car struct {
	Id                   string
	Name                 string
	Model                string
	Year                 string
	Licence_Plate_Number string
	Avilable             bool
	PricePerDay          float64
	Mu                   sync.Mutex
}

func (c *Car) IsAvilable() bool {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Avilable

}
func (c *Car) SetAvilability(av bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Avilable = av
}
