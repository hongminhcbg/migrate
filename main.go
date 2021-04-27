package main

import "update-customer-image/cmd"

func main() {
	cmd.Execute()
	//fileName := fmt.Sprintf("result_%d.csv", time.Now().Unix())
	//f, err := os.OpenFile(fileName,
	//	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//ioutil.WriteFile(fileName, []byte("Name, \tError, \tMessage\n"), 0666)
	//
	//fmt.Println("Hello World")
	//appID := os.Getenv("APP_ID")
	//appSecret := os.Getenv("APP_SECRET")
	//baseURL := os.Getenv("BASE_URL")
	//fmt.Printf(`[DB] Run with config, {"app_id":"%s", "app_secret":"%s", "base_url":"%s"`, appID, appSecret, baseURL)
	//fmt.Println()
	//
	//for i:= 3; i >= 0; i-- {
	//	fmt.Println("[DB] Start in ", i)
	//	time.Sleep(1 * time.Second)
	//}
	//
	//erpClient := erp_clients.NewERPClient(appID, appSecret, baseURL)
	//workers.NewUpdateCustomerImage(erpClient)
	//
	//tasks := make(chan string, 100)
	//results := make(chan workers.ResultTask, 100)
	//defer close(tasks)
	//defer close(results)
	//
	//go func() {
	//	startIndex := 0
	//	pageSize := 20
	//	for {
	//		customers, err := erpClient.GetCustomer(startIndex, pageSize)
	//		if err != nil {
	//			continue
	//		}
	//
	//		if len(customers.Data) == 0{
	//			break
	//		}
	//
	//		startIndex += len(customers.Data)
	//		for _, customerName := range customers.Data {
	//			fmt.Println("[DB] customer name = ", customerName.Name)
	//			tasks <- customerName.Name
	//		}
	//	}
	//}()
	//
	//worker := workers.NewUpdateCustomerImage(erpClient)
	//go worker.Start(tasks, results)
	//
	//for {
	//	select {
	//	case resultTask := <- results:
	//		fmt.Println("[DB] Result task = ", resultTask)
	//		f.WriteString(fmt.Sprintf("%s, \t%t, \t%s\n", resultTask.CustomerName, resultTask.Error, resultTask.Message))
	//	case <- time.After(3 * time.Minute):
	//		fmt.Println("[DB] Task is finish, bye")
	//		time.Sleep(2 * time.Second)
	//		os.Exit(0)
	//	}
	//}
}

