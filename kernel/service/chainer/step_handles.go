package chainer

type StepHandle int

const (
	NonStep StepHandle = iota // Non step

	StartRequestEmailStep // /start
	StartSendConfirmCodeStep
	StartCheckConfirmCodeStep
	StartChangeCampusStep
	StartSetCampusStep

	StaffShowBtnBookingsStep // бронирование
	StaffCreateBookingStep
	StaffShowBookingsStep

	StaffCreateBookingsStep // 1
	StaffChangeTypeStep     // 2
	StaffChangeCategoryStep // 3
	StaffChangeObjectStep   // 4
	StaffChangeDateStep     // 5
	StaffChangeTimeStep     // 6
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
}
