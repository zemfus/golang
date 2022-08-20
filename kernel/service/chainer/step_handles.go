package chainer

type StepHandle int

const (
	NonStep                   StepHandle = 0 // Non step
	StartRequestEmailStep     StepHandle = 1 // /start
	StartSendConfirmCodeStep  StepHandle = 2
	StartCheckConfirmCodeStep StepHandle = 3
	StartChangeCampusStep     StepHandle = 4
	StartSetCampusStep        StepHandle = 5
	StaffShowBtnBookingsStep  StepHandle = 6 // бронирование
	StaffProxyCreateVSShow    StepHandle = 7
	StaffShowBookingsStep     StepHandle = 8
	StaffChangeTypeStep       StepHandle = 9  // 2
	StaffChangeCategoryStep   StepHandle = 10 // 3
	StaffChangeObjectStep     StepHandle = 11 // 4
	StaffChangeDateStep       StepHandle = 12 // 5
	StaffChangeTimeStep       StepHandle = 13 // 6
	StaffCreateBookingStep    StepHandle = 14 // 7

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
}
