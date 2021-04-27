package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"time"

	erp_clients "update-customer-image/erp-clients"
)

var getUserCommand = &cobra.Command{
	Use: "get-users",
	Run: func(cmd *cobra.Command, args []string) {
		customersData := make(map[string]erp_clients.CustomerData)

		tasks := make(chan string, 100)
		customers := make(chan erp_clients.CustomerData, 100)
		defer close(tasks)
		defer close(customers)

		go func() {
			startIndex := 0
			pageSize := 20
			for {
				customers, err := erpClient.GetCustomer(startIndex, pageSize)
				if err != nil {
					continue
				}

				if len(customers.Data) == 0 {
					break
				}

				startIndex += len(customers.Data)
				for _, customerName := range customers.Data {
					fmt.Println("[DB] customer name = ", customerName.Name)
					tasks <- customerName.Name
				}
			}
		}()

		go func(tasks chan string, results chan erp_clients.CustomerData) {
			for customerName := range tasks {
				customerData, err := erpClient.GetCustomerByName(customerName)
				if err != nil {
					results <- erp_clients.CustomerData{
						Error: err.Error(),
					}
					continue
				}
				results <- *customerData
			}
		}(tasks, customers)

		for {
			select {
			case customerData := <-customers:
				if len(customerData.IdBackImage) > 0 ||
					len(customerData.IdFrontImage) > 0 {
					fmt.Println("[DB] Get customer success: ", customerData)
					customersData[customerData.Name] = customerData
				} else {
					fmt.Printf("[DB] Get customer success but filter fail, customer = %+v\n", customerData)
				}

			case <-time.After(3 * time.Minute):
				b, _ := json.MarshalIndent(customersData, "", "  ")
				err := ioutil.WriteFile("users.json", b, 0664)
				if err != nil {
					panic(err)
				}

				fmt.Println("get customers success")
			}
		}
	},
}
