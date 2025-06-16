package repository

import (
	"sync"

	"golang/Golang-Design/CarRental/models"
)

type CarRepo struct {
	CarRepo map[string]*models.Car
	mu      sync.Mutex
}

func NewCarRepo() *CarRepo {
	return &CarRepo{
		CarRepo: make(map[string]*models.Car),
	}
}

func (c *CarRepo) AddCar(car *models.Car) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CarRepo[car.Licence_Plate_Number] = car

}

func (c *CarRepo) GetAllCar() []*models.Car {

	carList := make([]*models.Car, 0)
	for _, c := range c.CarRepo {
		carList = append(carList, c)
	}
	return carList
}

func (c *CarRepo) FindAvilableCar(name, model string) []*models.Car {
	c.mu.Lock()
	defer c.mu.Unlock()
	var results []*models.Car
	for _, car := range c.CarRepo {
		if car.Name == name && car.Model == model {
			results = append(results, car)
		}
	}
	return results

}
