package seeds

import (
	"apart_community/internals/apartment/domain"
	"apart_community/internals/common/utils"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

func ApartmentSeed(db *gorm.DB) {
	apartments := []domain.Apartment{
		{PublicID: utils.GeneratePublicId(), Code: "001", Name: "Apartment 1"},
		{PublicID: utils.GeneratePublicId(), Code: "002", Name: "Apartment 2"},
		{PublicID: utils.GeneratePublicId(), Code: "003", Name: "Apartment 3"},
		{PublicID: utils.GeneratePublicId(), Code: "004", Name: "Apartment 4"},
	}

	for _, apt := range apartments {
		err := db.Create(&apt).Error

		if err != nil {
			fmt.Println("apartment seeding failed - ", err)
			return
		}

		var buildings []domain.Building

		for i := 1; i <= 4; i++ {
			var building domain.Building
			buildingName := strconv.Itoa(i) + "01"
			building.ApartmentID = apt.ID
			building.Name = &buildingName

			buildings = append(buildings, building)
		}

		for _, b := range buildings {
			err = db.Create(&b).Error

			if err != nil {
				fmt.Println("building seeding failed - ", err)
				return
			}

			for i := 10; i <= 14; i++ {
				var unit domain.Unit
				unit.BuildingID = b.ID
				unit.No = strconv.Itoa(i) + "01"
				err = db.Create(&unit).Error

				if err != nil {
					fmt.Println("unit seeding failed - ", err)
					return
				}
			}
		}

	}
}
