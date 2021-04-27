package workers

import erp_clients "update-customer-image/erp-clients"

type UpdateCustomerImage struct {
	erpClient *erp_clients.ErpClient
}

func NewUpdateCustomerImage(erpClient *erp_clients.ErpClient) *UpdateCustomerImage {
	return &UpdateCustomerImage{erpClient: erpClient}
}

type ResultTask struct {
	Error bool
	Message string
	CustomerName string
}

func (w *UpdateCustomerImage) Start(tasks chan string, results chan ResultTask) {
	for customerName := range tasks{
		resp, err := w.erpClient.UpdateCustomerImageToPrivate(customerName)
		if err  != nil {
			results <- ResultTask{
				Error:   true,
				Message: err.Error(),
				CustomerName: customerName,
			}

			continue
		}

		results <- ResultTask{
			Error:   false,
			Message: resp,
			CustomerName: customerName,
		}
	}
}