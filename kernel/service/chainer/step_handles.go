package chainer

type StepHandle int

const (
	NonStep StepHandle = 0 // Non step

	StartRequestEmailStep     StepHandle = 1 // /start ================
	StartSendConfirmCodeStep  StepHandle = 2
	StartCheckConfirmCodeStep StepHandle = 3
	StartChangeCampusStep     StepHandle = 4
	StartSetCampusStep        StepHandle = 5

	StaffShowBtnBookingsStep StepHandle = 6 // бронирование ===========
	StaffProxyCreateVSShow   StepHandle = 7
	StaffShowBookingsStep    StepHandle = 8
	StaffChangeTypeStep      StepHandle = 9  // 2
	StaffChangeCategoryStep  StepHandle = 10 // 3
	StaffChangeObjectStep    StepHandle = 11 // 4
	StaffChangeDateStep      StepHandle = 12 // 5
	StaffChangeTimeStep      StepHandle = 13 // 6
	StaffCreateBookingStep   StepHandle = 14 // 7

	CfgShowBtnStep    StepHandle = 15 // конфигурация ========================
	CfgProxyItemsStep StepHandle = 16

	CfgCampusStep           StepHandle = 17
	CfgGetCampusNameStep    StepHandle = 177
	CfgSetCampusNameStep    StepHandle = 178
	CfgCampusCreateStep     StepHandle = 1777
	CfgCampusEditStep       StepHandle = 171
	CfgCampusDeleteStep     StepHandle = 172
	CfgCampusUpdateStep     StepHandle = 173
	CfgCampusUpdateExecStep StepHandle = 174
	CfgCategoryStep         StepHandle = 18
	CfgPlaceStep            StepHandle = 19
	CfgInventoryStep        StepHandle = 20
	CfgStudentsStep         StepHandle = 21
)

var StartSteps = []StepHandle{
	StartRequestEmailStep,
	StartSendConfirmCodeStep,
	StartCheckConfirmCodeStep,
	StartChangeCampusStep,
	StartSetCampusStep,
}

var StaffBookingSteps = []StepHandle{
	StaffShowBtnBookingsStep,
	StaffProxyCreateVSShow,

	StaffShowBookingsStep,
	StaffChangeTypeStep,
	StaffChangeCategoryStep,
	StaffChangeObjectStep,
	StaffChangeDateStep,
	StaffChangeTimeStep,
	StaffCreateBookingStep,
	CfgCampusEditStep,
}

var CfgSteps = []StepHandle{
	CfgShowBtnStep,
	CfgProxyItemsStep,
	CfgCampusStep,
	CfgCategoryStep,
	CfgPlaceStep,
	CfgInventoryStep,
	CfgStudentsStep,
	CfgGetCampusNameStep,
	CfgSetCampusNameStep,
	CfgCampusEditStep,
	CfgCampusDeleteStep,
	CfgCampusUpdateStep,
	CfgCampusUpdateExecStep,
}
