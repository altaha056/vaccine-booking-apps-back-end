package factory

import (
	"vac/driver"

	userData "vac/features/user/data"
	userPresent "vac/features/user/presentation"
	userService "vac/features/user/service"

	adminData "vac/features/admin/data"
	adminPresent "vac/features/admin/presentation"
	adminService "vac/features/admin/service"

	vacData "vac/features/vac/data"
	vacPresent "vac/features/vac/presentation"
	vacService "vac/features/vac/service"

	participantData "vac/features/participant/data"
	participantPresent "vac/features/participant/presentation"
	participantService "vac/features/participant/service"

	
)

type vacPresenter struct{
	UserPresentation 		userPresent.UserHandler
	AdminPresentation 		adminPresent.AdminHandler
	VacPresentation 		vacPresent.VacHandler
	ParticipantPresentation	participantPresent.ParticipantHandler
}

func Init() vacPresenter{
	userData:=userData.NewMysqlUserRepository(driver.DB)
	userService:=userService.NewUserService(userData)

	adminData:=adminData.NewAdminRepository(driver.DB)
	adminService:=adminService.NewAdminService(adminData)
	
	vacData:=vacData.NewMysqlVaccineRepository(driver.DB)
	vacService:=vacService.NewVacUseCase(vacData)

	parData:=participantData.NewMysqlParticipantRepository(driver.DB)
	parService:=participantService.NewParService(parData, vacService, userService)

	return vacPresenter{
		UserPresentation: *userPresent.NewUserHandler(userService),
		AdminPresentation: *adminPresent.NewAdminHandler(adminService),
		VacPresentation: *vacPresent.NewVacHandler(vacService),
		ParticipantPresentation: *participantPresent.NewParticipantHandler(parService),
	}
}