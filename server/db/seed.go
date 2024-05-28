package db

func SeedDatabase() error {
	//if !(DB.Migrator().HasTable(&models.Role{}) && DB.Migrator().HasTable(&models.Team{})) {
	//	slog.Error("Database tables are not created, execute database reinit first")
	//	return fmt.Errorf("database is not initialized")
	//}
	//
	//err := DB.Transaction(func(tx *gorm.DB) error {
	//	// Add predefined roles
	//	result := tx.Create([]models.Role{
	//		{
	//			Name:        "ATTENDING",
	//			DisplayName: "Attending Physician",
	//		},
	//		{
	//			Name:        "RESIDENT",
	//			DisplayName: "Resident Doctor",
	//		},
	//		{
	//			Name:        "NURSE",
	//			DisplayName: "Nurse",
	//		},
	//		{
	//			Name:        "TECHNICIAN",
	//			DisplayName: "Technician",
	//		},
	//	})
	//	if result.Error != nil {
	//		slog.Error(fmt.Sprintf("Failed to create predefined roles: %s", result.Error.Error()))
	//		return result.Error
	//	}
	//
	//	// Add predefined teams
	//	result = tx.Create([]models.Team{
	//		{
	//			Name:        "GENERAL",
	//			DisplayName: "Internal Medicine",
	//		},
	//		{
	//			Name:        "CARDIO",
	//			DisplayName: "Cardiology",
	//		},
	//		{
	//			Name:        "NEURO",
	//			DisplayName: "Neurology",
	//		},
	//		{
	//			Name:        "ORTHO",
	//			DisplayName: "Orthopedics",
	//		},
	//		{
	//			Name:        "OB-GYN",
	//			DisplayName: "Obstetrics and Gynecology",
	//		},
	//		{
	//			Name:        "PEDS",
	//			DisplayName: "Pediatrics",
	//		},
	//		{
	//			Name:        "ONCOLOGY",
	//			DisplayName: "Oncology",
	//		},
	//		{
	//			Name:        "PSYCH",
	//			DisplayName: "Psychiatry",
	//		},
	//		{
	//			Name:        "UROLOGY",
	//			DisplayName: "Urology",
	//		},
	//	})
	//
	//	return nil
	//})
	//if err != nil {
	//	slog.Error("Failed to seed the database")
	//	return err
	//}

	return nil
}
