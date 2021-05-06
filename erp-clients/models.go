package erp_clients

type CustomerData struct {
	IdFrontImage   string            `json:"id_front_image,omitempty"`
	IdBackImage    string            `json:"id_back_image,omitempty"`
	SelfieImage    string            `json:"selfie_image,omitempty"`
	SignatureImage string            `json:"signature_image,omitempty"`
	Image          string            `json:"image,omitempty"`
	Name           string            `json:"name,omitempty"`
	Error          string            `json:"error,omitempty"`
	MigrateStatus  string            `json:"migrate_status,omitempty"`
	AttachStatus   map[string]string `json:"attach_status,omitempty"`
}

type FileDetails struct {
	Name              string `json:"name,omitempty"`
	FileURL           string `json:"file_url,omitempty"`
	AttachedToDoctype string `json:"attached_to_doctype,omitempty"`
	AttachedToName    string `json:"attached_to_name,omitempty"`
}

type GetFileResp struct {
	Data []FileDetails `json:"data"`
}

type CustomerDataResp struct {
	Data CustomerData `json:"data"`
}

type AttachFileToCustomerRequest struct {
	AttachedToDoctype string `json:"attached_to_doctype,omitempty"`
	AttachedToName    string `json:"attached_to_name,omitempty"`
}
