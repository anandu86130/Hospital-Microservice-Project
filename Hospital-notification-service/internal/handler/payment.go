package handler

func (n *notifcationHandler) PaymentHandler() error {
	return n.services.SubscribeAndConsumePaymentEvents()
}

func (n *notifcationHandler) AppointmentResultHandler() error {
	return n.services.SubScribeAndConsumeAppointmentEvents()
}
