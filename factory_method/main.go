package main

type IMethod interface {
	Drive() string
}

type Car struct{}

func (c *Car) Drive() string {
	return "Driving a car"
}

type Bike struct{}

func (b *Bike) Drive() string {
	return "Riding a bike"
}

type VehicleFactory interface {
	CreateVehicle() IMethod
}

type CarFactory struct{}

func (f *CarFactory) CreateVehicle() IMethod {
	return &Car{}
}

type BikeFactory struct{}

func (f *BikeFactory) CreateVehicle() IMethod {
	return &Bike{}
}

func GetVehicleFactory(vehicleType string) VehicleFactory {
	if vehicleType == "car" {
		return &CarFactory{}
	} else if vehicleType == "bike" {
		return &BikeFactory{}
	}
	return nil
}

func main() {
	carFactory := GetVehicleFactory("car")
	car := carFactory.CreateVehicle()
	println(car.Drive())

	bikeFactory := GetVehicleFactory("bike")
	bike := bikeFactory.CreateVehicle()
	println(bike.Drive())
}
