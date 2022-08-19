package chainer

type StepHandle int

const (
	NonStep StepHandle = iota // Non step

	StartRequestEmailStep   // /start
	StartSendConfirmURLStep // /start
	StartWaitingConfirmStep
	StartChangeCampusStep

	CreateBookingReceiveDataStep // /create_booking
	CreateBookingValidAndSetDataStep

	DeleteBookingGetTypeAndReasonStep // /delete_bookings
	DeleteBookingExecTypeStep

	ShowBookingsStep // /show_bookings

	Help // /help

	BookingChoiceStep // /booking
	BookingSaveStep
)

var StartSteps = []StepHandle{
	StartRequestEmailStep,
	StartSendConfirmURLStep,
	StartWaitingConfirmStep,
	StartChangeCampusStep}
