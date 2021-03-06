package erp_clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ErpClient struct {
	appID           string
	appSecret       string
	baseURL         string
	getCustomerURL  string
	getFileURL      string
	migrateImageURL string
	token           string
}

type Resource struct {
	Name string `json:"name"`
}

type Resources struct {
	Data []Resource `json:"data"`
}

func NewERPClient(appID, appSecret, baseURL string) *ErpClient {
	result := &ErpClient{
		appID:     appID,
		appSecret: appSecret,
		baseURL:   baseURL,
	}

	result.getCustomerURL = baseURL + "/api/resource/Customer"
	result.token = fmt.Sprintf("token %s:%s", result.appID, result.appSecret)
	result.migrateImageURL = fmt.Sprintf("%s/api/method/en_crm.controller.migrade_file_of_users_to_private", baseURL)
	result.getFileURL = fmt.Sprintf("%s/api/resource/File", baseURL)
	return result
}

func (c *ErpClient) GetCustomer(limitStart int, limitPageLength int) (*Resources, error) {
	url := fmt.Sprintf("%s?limit_start=%d&limit_page_length=%d", c.getCustomerURL, limitStart, limitPageLength)
	httpRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("[ERROR] make http request error: ", err)
		return nil, err
	}
	httpRequest.Header.Set("Authorization", c.token)

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Println("[ERROR] do http request error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	var result Resources
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] read body: ", err)
		return nil, err
	}

	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Println("[ERROR] read body: ", err)
		return nil, err
	}

	return &result, nil

}

func (c *ErpClient) UpdateCustomerImageToPrivate(customerName string) (string, error) {
	bodyStr := fmt.Sprintf(`{"customer": "%s"}`, customerName)
	httpRequest, err := http.NewRequest(http.MethodPost, c.migrateImageURL, strings.NewReader(bodyStr))
	if err != nil {
		log.Println("[ERROR] make http request error: ", err)
		return "", err
	}
	httpRequest.Header.Set("Authorization", c.token)
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Println("[ERROR] do http request error: ", err)
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] read body: ", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ERROR_CODE_%d", resp.StatusCode)
	}

	return string(b), nil
}

func (c *ErpClient) GetCustomerByName(customerName string) (*CustomerData, error) {
	url := fmt.Sprintf("%s/%s", c.getCustomerURL, customerName)

	httpRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("[ERROR] make http request error: ", err)
		return nil, err
	}
	httpRequest.Header.Set("Authorization", c.token)
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Println("[ERROR] do http request error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] read body: ", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ERROR_CODE_%d", resp.StatusCode)
	}

	var result CustomerDataResp
	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Println("[ERROR] json marshal: ", err)
		return nil, err
	}
	//
	//log.Printf("[DB] customer data = %+v\n", result)

	return &result.Data, nil
}

func (c *ErpClient) GetFileByURL(fileURL string) (*FileDetails, error) {

	query := fmt.Sprintf("fields=[\"name\",\"file_url\",\"attached_to_doctype\", \"attached_to_name\"]&filters=[[\"file_url\", \"=\", \"%s\"]]", fileURL)

	values, err := url.ParseQuery(query)
	if err != nil {
		log.Println("[ERROR] parse query error: ", err)
		return nil, err
	}

	url := fmt.Sprintf("%s?%s", c.getFileURL, values.Encode())
	httpRequest, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println("[ERROR] make http request error: ", err)
		return nil, err
	}
	httpRequest.Header.Set("Authorization", c.token)
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Println("[ERROR] do http request error: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[ERROR] read body: ", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ERROR_CODE_%d", resp.StatusCode)
	}

	var result GetFileResp
	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Println("[ERROR] json marshal: ", err)
		return nil, err
	}

	if len(result.Data) == 0 {
		return nil, fmt.Errorf("ERROR_CODE_404")
	}

	return &result.Data[0], nil

}

func (c *ErpClient) AttachFileToCustomer(customerName string, fileName string) error  {
	url := fmt.Sprintf("%s/%s", c.getFileURL, fileName)
	requestBody := AttachFileToCustomerRequest{
		AttachedToDoctype: "Customer",
		AttachedToName:    customerName,
	}
	b, _ := json.Marshal(requestBody)


	httpRequest, err := http.NewRequest(http.MethodPut, url, bytes.NewReader(b))
	if err != nil {
		log.Println("[ERROR] make http request error: ", err)
		return err
	}
	httpRequest.Header.Set("Authorization", c.token)
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		log.Println("[ERROR] do http request error: ", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ERROR_CODE_%d", resp.StatusCode)
	}

	return nil
}