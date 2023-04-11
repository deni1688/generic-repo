package main

type Driver struct {
	*Base
	Name string `json:"name"`
}

func newDriver(name string) Driver {
	return Driver{Name: name, Base: &Base{}}
}

type Vehicle struct {
	*Base
	Make string `json:"make"`
}

func newVehicle(make string) Vehicle {
	return Vehicle{Make: make, Base: &Base{}}
}
