package model

type CheckRegistrationAvailabilityReq struct {
	Name  *string
	Email *string
	Phone *string
}

type CheckRegistrationAvailabilityResp struct {
	NameAvailable    *bool
	EmailCanRegister *bool
	PhoneCanRegister *bool
}
