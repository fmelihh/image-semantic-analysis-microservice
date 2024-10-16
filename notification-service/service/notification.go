package service

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (s *NotificationService) Notify(consumerMessage map[string]any) error {
	return nil
}
