package seeds

import (
	"apart_community/internals/common/utils"
	"apart_community/internals/user/domain"
	"apart_community/internals/user/domain/roles"
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func UserSeed(db *gorm.DB) {
	var users []domain.User

	for i := 1; i <= 10; i++ {
		var user domain.User
		user.PublicID = utils.GeneratePublicId()
		user.Email = "test" + strconv.Itoa(i) + "@test.com"
		hashed, _ := utils.HashPassword("test1234")
		user.Password = string(hashed)
		users = append(users, user)
	}

	userRoles := []domain.UserRole{
		{
			TargetType: "SYSTEM",
			TargetID:   0,
			RoleName:   roles.RoleSystemAdmin.String(),
			IsVerified: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			TargetType: "APARTMENT",
			TargetID:   1,
			RoleName:   roles.RoleApartmentAdmin.String(),
			IsVerified: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			TargetType: "APARTMENT",
			TargetID:   1,
			RoleName:   roles.RoleApartmentAdmin.String(),
			IsVerified: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			TargetType: "BUILDING",
			TargetID:   1,
			RoleName:   roles.RoleBuildingAdmin.String(),
			IsVerified: true,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			TargetType: "BUILDING",
			TargetID:   2,
			RoleName:   roles.RoleBuildingAdmin.String(),
			IsVerified: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
		{
			TargetType: "GUEST",
			TargetID:   0,
			RoleName:   roles.RoleGuest.String(),
			IsVerified: false,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		},
	}

	userUnitRoles := []domain.UserUnitRole{
		{
			UnitID:     1,
			RoleName:   roles.RoleOwner.String(),
			IsVerified: true,
		},
		{
			UnitID:     2,
			RoleName:   roles.RoleOwner.String(),
			IsVerified: false,
		},
		{
			UnitID:     1,
			RoleName:   roles.RoleResident.String(),
			IsVerified: true,
		},
		{
			UnitID:     3,
			RoleName:   roles.RoleResident.String(),
			IsVerified: false,
		},
	}

	for i, user := range users {
		err := db.Create(&user).Error

		if err != nil {
			fmt.Println("user seeding failed - ", err)
			return
		}

		profile := domain.Profile{
			UserID:   user.ID,
			Nickname: strings.Split(user.Email, "@")[0],
		}

		err = db.Create(&profile).Error

		if err != nil {
			fmt.Println("profile seeding failed - ", err)
			return
		}

		if i <= 5 {
			var userRole domain.UserRole
			userRole = userRoles[i]
			userRole.UserID = user.ID

			err = db.Create(&userRole).Error

			if err != nil {
				fmt.Println("user role seeding failed - ", err)
				return
			}
		} else {
			var userUnitRole domain.UserUnitRole
			userUnitRole = userUnitRoles[i-6]
			userUnitRole.UserID = user.ID

			err = db.Create(&userUnitRole).Error

			if err != nil {
				fmt.Println("user unit role seeding failed - ", err)
				return
			}
		}

	}
}
