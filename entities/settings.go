package entities

type NotificationSettings struct {
	EmailNotifications bool `json:"email_notifications" binding:"required"`
	Promotions         bool `json:"promotions" binding:"required"`
	SmsNotifications   bool `json:"sms_notifications" binding:"required"`
}

type PaymentSettings struct {
	PrimaryPaymentMethod bool `json:"primary_payment_method" binding:"required"`
	TipPerPayment        int  `json:"tip_per_payment" binding:"required"`
}

type SecuritySettings struct {
	EmailNotifications bool `json:"email_notifications" binding:"required"`
	Promotions         bool `json:"promotions" binding:"required"`
	SmsNotifications   bool `json:"sms_notifications" binding:"required"`
}
