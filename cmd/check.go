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

func compare(oldURL, newURL string) bool {

	if oldURL == "" && newURL == "" {
		return true
	}

	if oldURL == newURL && strings.Contains(oldURL, "private") {
		return true
	}

	if newURL == "/private" + oldURL {
		return true
	}

	return false
}

type CheckError struct {
	Username string `json:"username"`
	Reason   string `json:"reason"`
}

var checkCmd = &cobra.Command{
	Use: "check",
	Run: func(cmd *cobra.Command, args []string) {
		users := make(map[string]erp_clients.CustomerData)
		b, err := ioutil.ReadFile("./users.json.bak")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(b, &users)
		if err != nil {
			panic(err)
		}

		sum := 0
		success := 0
		errorList := make([]CheckError, 0)

		for userName, userData := range users {
			sum++
			userDataNew, err := erpClient.GetCustomerByName(userName)
			if err != nil {
				log.Println("[ERROR] get user data error: ", err)
				errorList = append(errorList, CheckError{
					Username: userName,
					Reason:   err.Error(),
				})
				continue
			}

			if ok := compare(userData.Image, userDataNew.Image); !ok {
				log.Printf("[ERROR][%s] old Image url is %s but new Image url is %s\n",userName,userData.Image, userDataNew.Image)
				errorList = append(errorList, CheckError{
					Username: userName,
					Reason:   fmt.Sprintf("old image url is %s but new image url is %s",userData.Image, userDataNew.Image),
				})
				continue
			}

			if ok := compare(userData.IdFrontImage, userDataNew.IdFrontImage); !ok {
				log.Printf("[ERROR][%s] old IdFrontImage url is %s but new IdFrontImage url is %s\n",userName,userData.IdFrontImage, userDataNew.IdFrontImage)
				errorList = append(errorList, CheckError{
					Username: userName,
					Reason:   fmt.Sprintf("old id_front url is %s but new id_front url is %s",userData.IdFrontImage, userDataNew.IdFrontImage),
				})
				continue
			}

			if ok := compare(userData.IdBackImage, userDataNew.IdBackImage); !ok {
				log.Printf("[ERROR][%s] old IdBackImage url is %s but new IdBackImage url is %s\n",userName,userData.IdBackImage, userDataNew.IdBackImage)

				errorList = append(errorList, CheckError{
					Username: userName,
					Reason:   fmt.Sprintf("old id_back url is %s but new id_back url is %s",userData.IdBackImage, userDataNew.IdBackImage),
				})
				continue
			}

			if ok := compare(userData.SelfieImage, userDataNew.SelfieImage); !ok {
				log.Printf("[ERROR][%s] old SelfieImage url is %s but new SelfieImage url is %s\n",userName,userData.SelfieImage, userDataNew.SelfieImage)
				errorList = append(errorList, CheckError{
					Username: userName,
					Reason:   fmt.Sprintf("old SelfieImage url is %s but new SelfieImage url is %s",userData.SelfieImage, userDataNew.SelfieImage),
				})
				continue
			}

			if ok := compare(userData.SignatureImage, userDataNew.SignatureImage); !ok {
				log.Printf("[ERROR][%s] old SignatureImage url is %s but new SignatureImage url is %s\n",userName,userData.SignatureImage, userDataNew.SignatureImage)
				errorList = append(errorList, CheckError{
					Username: userName,
					Reason:   fmt.Sprintf("old SignatureImage url is %s but new SignatureImage url is %s",userData.SignatureImage, userDataNew.SignatureImage),
				})
				continue
			}

			log.Printf("[DB] data match %s\n", userName)
			success++
		}

		bytes, _ := json.MarshalIndent(errorList, "", "  ")
		err = ioutil.WriteFile("users_migrate_check.json", bytes, 0664)
		if err != nil {
			panic(err)
		}

		fmt.Println("[DB] Check success, sum record = ", sum, ", record match = ", success)
	},
}
