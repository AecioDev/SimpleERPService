package seeders

import "gorm.io/gorm"

func RunAll(db *gorm.DB) {
	SeedCountries(db)
	SeedStates(db)
	SeedCities(db)

	SeedRolesPermissions(db)
	SeedUserAdm(db)
	SeedMeasurementUnit(db)
	SeedPaymentMethod(db)
	SeedProductCategory(db)
}
