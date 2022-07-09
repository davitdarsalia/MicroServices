package service

import "github.com/davitdarsalia/LendAppBackend/entities"

func (s *SettingsService) GetProfileSettings() {
	//TODO implement me
}

func (s *SettingsService) GetNotificationSettings() {
	//TODO implement me
}

func (s *SettingsService) GetPaymentOptions() {
	//TODO implement me
}

func (s *SettingsService) GetPrivacySettings() {
	//TODO implement me
}

func (s *SettingsService) GetSecuritySettings() {
	//TODO implement me
}

func (s *SettingsService) UpdateNotificationSettings(n *entities.NotificationSettings) error {
	id := s.getRedisID()

	return s.repo.UpdateNotificationSettings(&id, n)
}

func (s *SettingsService) UpdatePaymentSettings(p *entities.PaymentSettings) error {
	id := s.getRedisID()

	return s.repo.UpdatePaymentSettings(&id, p)
}

func (s *SettingsService) UpdateSecuritySettings(p *entities.SecuritySettings) error {
	id := s.getRedisID()

	return s.repo.UpdateSecuritySettings(&id, p)
}

func (s *SettingsService) UpdatePrivacySettings() {
	//TODO implement me
}
