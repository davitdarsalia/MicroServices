package repository

import "github.com/davitdarsalia/LendAppBackend/entities"

func (r *SettingsPostgres) GetProfileSettings() {
	//TODO implement me
}

func (r *SettingsPostgres) GetNotificationSettings() {
	//TODO implement me
}

func (r *SettingsPostgres) GetPaymentOptions() {
	//TODO implement me
}

func (r *SettingsPostgres) GetPrivacySettings() {
	//TODO implement me
}

func (r *SettingsPostgres) GetSecuritySettings() {
	//TODO implement me
}

func (r *SettingsPostgres) UpdateNotificationSettings(userID *int, s *entities.NotificationSettings) error {
	return nil
}

func (r *SettingsPostgres) UpdatePaymentSettings(userID *int, s *entities.PaymentSettings) error {
	return nil
}

func (r *SettingsPostgres) UpdateSecuritySettings(userID *int, s *entities.SecuritySettings) error {
	return nil
}

func (r *SettingsPostgres) UpdatePrivacySettings() error {
	return nil
}
