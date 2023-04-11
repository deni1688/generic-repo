package main

type VehicleRepo interface {
	Repo[Vehicle]
	FindByMake(model string) ([]*Vehicle, error)
}

type inMemVehicleRepo struct {
	*inMemRepo[Vehicle]
}

func newInMemVehicleRepo(data map[int]Vehicle) VehicleRepo {
	return &inMemVehicleRepo{inMemRepo: newInMemRepo[Vehicle](data)}
}

func (i *inMemVehicleRepo) FindByMake(make string) ([]*Vehicle, error) {
	var result []*Vehicle

	for _, v := range i.data {
		if v.Make == make {
			result = append(result, &v)
		}
	}

	return result, nil
}
