package repository

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
)

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
	_, err := r.db.Exec(constants.UpdateNotificationSettings, s.EmailNotifications, s.Promotions, s.SmsNotifications, userID)
	fmt.Println(err)

	return err
}

func (r *SettingsPostgres) UpdatePaymentSettings(userID *int, s *entities.PaymentSettings) error {
	_, err := r.db.Exec(constants.UpdatePaymentSettings, s.PrimaryPaymentMethod, s.TipPerPayment, userID)
	return err
}

func (r *SettingsPostgres) UpdateSecuritySettings(userID *int, s *entities.SecuritySettings) error {
	_, err := r.db.Exec(constants.UpdateSecuritySettings, s.Contacts, s.HideEmail, s.HideMobile, s.HideActivity, userID)
	return err
}

func (r *SettingsPostgres) UpdatePrivacySettings() error {
	return nil
}
