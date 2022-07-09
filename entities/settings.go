package entities

type NotificationSettings struct {
	UserID             int  `json:"userid"`
	EmailNotifications bool `json:"email_notifications" binding:"required"`
	Promotions         bool `json:"promotions" binding:"required"`
	SmsNotifications   bool `json:"sms_notifications" binding:"required"`
}

type PaymentSettings struct {
	UserID               int  `json:"userid"`
	PrimaryPaymentMethod bool `json:"primary_payment_method" binding:"required"`
	TipPerPayment        int  `json:"tip_per_payment" binding:"required"`
}

type SecuritySettings struct {
	UserID       int  `json:"userid"`
	Contacts     bool `json:"contacts" binding:"required"`
	HideEmail    bool `json:"hide_email" binding:"required"`
	HideMobile   bool `json:"hide_mobile" binding:"required"`
	HideActivity bool `json:"hide_activity" binding:"required"`
}
