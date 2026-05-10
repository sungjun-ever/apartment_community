package seeds

import (
	"apart_community/internals/apartment/domain"
	"apart_community/internals/common/utils/idgen"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func ApartmentSeed(db *gorm.DB) {
	apt := domain.Apartment{
		PublicID: idgen.GeneratePublicId(), Code: "001", Name: "Apartment 1",
	}

	err := db.Create(&apt).Error

	if err != nil {
		fmt.Println("apartment seeding failed - ", err)
		return
	}

	var building domain.Building
	buildingName := "101"
	building.ApartmentID = apt.ID
	building.Name = &buildingName

	err = db.Create(&building).Error

	if err != nil {
		fmt.Println("building seeding failed - ", err)
		return
	}

	for i := 1; i <= 2; i++ {
		var unit domain.Unit
		unit.BuildingID = building.ID
		unit.No = strconv.Itoa(i) + "01"
		err = db.Create(&unit).Error

		if err != nil {
			fmt.Println("unit seeding failed - ", err)
			return
		}
	}

}
