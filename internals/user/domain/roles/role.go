package roles

type Role string

const (
	RoleSystemAdmin    Role = "SYSTEM_ADMIN"
	RoleApartmentAdmin Role = "APT_ADMIN"
	RoleBuildingAdmin  Role = "BUILDING_ADMIN"
	RoleOwner          Role = "OWNER"
	RoleResident       Role = "RESIDENT"
	RoleGuest          Role = "GUEST"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleSystemAdmin, RoleApartmentAdmin, RoleBuildingAdmin, RoleOwner, RoleResident, RoleGuest:
		return true
	}
	return false
}

func (r Role) String() string {
	return string(r)
}
