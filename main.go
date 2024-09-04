package main

import (
	"fmt"

	"github.com/bagashyt/learn_golang_echo/model"
)

var user1 = &model.User{
	Id:       "u001",
	Name:     "Sylva Windrunner",
	Password: "f0r Th3 H0rD3",
	Gender:   model.UserGender_FEMALE,
}

var userList = &model.UserList{
	List: []*model.User{
		user1,
	},
}

var garage1 = &model.Garage{
	Id:   "g001",
	Name: "Kalimdor",
	Coordinate: &model.GarageCoordinate{
		Latitude:  23.2212847,
		Longitude: 53.22033123,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

var garageListByUser = &model.GarageListByUser{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main() {

	fmt.Printf("# === Original \n     %#v \n", user1)

	fmt.Printf("# === As String\n     %#v \n", user1.String())

}
