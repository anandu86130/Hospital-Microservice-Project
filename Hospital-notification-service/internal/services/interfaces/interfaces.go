package interfaces

type NotificationServiceInter interface {
	SubscribeAndConsumePaymentEvents() error
	SubScribeAndConsumeAppointmentEvents() error
}
