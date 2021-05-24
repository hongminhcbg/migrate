package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"strings"
	erp_clients "update-customer-image/erp-clients"
)

func attachFileToCustomer(fileURL, customerName string) error {
	if !strings.Contains(fileURL, "private"){
		return fmt.Errorf("file not private")
	}

	fileDetail, err := erpClient.GetFileByURL(fileURL)
	if err != nil {
		log.Printf("[ERROR] Get file name error:%v, customer_name:%s, file_url:%s", err, customerName, fileURL)
		return err
	}

	if fileDetail.AttachedToName == customerName {
		log.Printf("[DB] Attach file %s to customer %s success", fileDetail.FileURL, customerName)
		return nil
	}

	err = erpClient.AttachFileToCustomer(customerName, fileDetail.Name)
	if err != nil {
		log.Printf("[ERROR] Attach file %s to customer %s error %v", fileDetail.FileURL, customerName, err)
		return err
	}

	log.Printf("[DB] Attach file %s to customer %s success", fileDetail.FileURL, customerName)
	return nil
}

var attachFieCmd = &cobra.Command{
	Use: "attach-file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[DB] This is attach file command")
		users := make(map[string]erp_clients.CustomerData)
		b, err := ioutil.ReadFile("./users.json.bak")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(b, &users)
		if err != nil {
			panic(err)
		}

		errorList := make([]CheckError, 0)

		for userName, userData := range users {
			if len(userData.IdFrontImage) > 0 {
				err = attachFileToCustomer(userData.IdFrontImage, userName)
				if err != nil {
					errorList = append(errorList, CheckError{
						Username: userName,
						Reason:   fmt.Sprintf("Attach file '%s' have error '%s'", userData.IdFrontImage, err.Error()),
					})
				}
			}

			if len(userData.IdBackImage) > 0 {
				attachFileToCustomer(userData.IdBackImage, userName)
				if err != nil {
					errorList = append(errorList, CheckError{
						Username: userName,
						Reason:   fmt.Sprintf("Attach file '%s' have error '%s'", userData.IdBackImage, err.Error()),
					})
				}

			}

			if len(userData.Image) > 0 {
				attachFileToCustomer(userData.Image, userName)
				if err != nil {
					errorList = append(errorList, CheckError{
						Username: userName,
						Reason:   fmt.Sprintf("Attach file '%s' have error '%s'", userData.Image, err.Error()),
					})
				}

			}

			if len(userData.SelfieImage) > 0 {
				attachFileToCustomer(userData.SelfieImage, userName)
				if err != nil {
					errorList = append(errorList, CheckError{
						Username: userName,
						Reason:   fmt.Sprintf("Attach file '%s' have error '%s'", userData.SelfieImage, err.Error()),
					})
				}

			}

			if len(userData.SignatureImage) > 0 {
				attachFileToCustomer(userData.SignatureImage, userName)
				if err != nil {
					errorList = append(errorList, CheckError{
						Username: userName,
						Reason:   fmt.Sprintf("Attach file '%s' have error '%s'", userData.SignatureImage, err.Error()),
					})
				}

			}
		}

		bytes, _ := json.MarshalIndent(errorList, "", "  ")
		err = ioutil.WriteFile("users_attach_check.json", bytes, 0664)
		if err != nil {
			panic(err)
		}

	},
}
