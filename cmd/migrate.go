package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	erp_clients "update-customer-image/erp-clients"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		users := make(map[string] erp_clients.CustomerData)
		b, err := ioutil.ReadFile("./users.json.bak")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(b, &users)
		if err != nil {
			panic(err)
		}

		for userName, userData := range users {
			_, err = erpClient.UpdateCustomerImageToPrivate(userName)
			if err != nil {
				log.Printf("[ERROR] error=%s, username= %s", err.Error(), userName)
				userData.MigrateStatus = "FAIL"
				userData.Error = err.Error()
				users[userName] = userData
				continue
			}

			log.Println("[DB] migrate success, user_name = ", userName)
			userData.MigrateStatus = "SUCCESS"
			users[userName] = userData
		}

		bytes, _ := json.MarshalIndent(users, "", "  ")
		err = ioutil.WriteFile("users_migrate.json", bytes, 0664)
		if err != nil {
			panic(err)
		}

		fmt.Println("migrate success")
	},
}
