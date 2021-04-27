package erp_clients

type CustomerData struct {
	IdFrontImage   string `json:"id_front_image,omitempty"`
	IdBackImage    string `json:"id_back_image,omitempty"`
	SelfieImage    string `json:"selfie_image,omitempty"`
	SignatureImage string `json:"signature_image,omitempty"`
	Image          string `json:"image,omitempty"`
	Name           string `json:"name,omitempty"`
	Error          string `json:"error,omitempty"`
	MigrateStatus  string `json:"migrate_status,omitempty"`
}

type CustomerDataResp struct {
	Data CustomerData `json:"data"`
}
