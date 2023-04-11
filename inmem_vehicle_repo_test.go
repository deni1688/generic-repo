package main

import (
	"testing"
)

func TestInMemVehicleRepo_FindByMake(t *testing.T) {
	vr := newInMemVehicleRepo(make(map[int]Vehicle))

	_, _ = vr.Create(newVehicle("Ford"))
	id, _ := vr.Create(newVehicle("Volvo"))

	vehicles, err := vr.FindByMake("Volvo")
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if len(vehicles) != 1 {
		t.Errorf("length is not 1")
	}

	if vehicles[0].GetID() != id {
		t.Errorf("id is not %d", id)
	}
}
