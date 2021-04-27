package erp_clients

import (
	"fmt"
	"testing"
)

func TestErpClient_GetCustomer(t *testing.T) {
	type fields struct {
		appID          string
		appSecret      string
		baseURL        string
		getCustomerURL string
		token          string
	}
	type args struct {
		limitStart      int
		limitPageLength int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Resources
		wantErr bool
	}{
		{
			name: "normal case",
			fields: fields{
				appID:          "76d97bb400d4df0",
				appSecret:      "f3383bca501367e",
				baseURL:        "http://localhost:8000",
				getCustomerURL: "http://localhost:8000" + "/api/resource/Customer",
				token:          fmt.Sprintf("token %s:%s", "76d97bb400d4df0", "f3383bca501367e"),
			},
			args: args{
				limitStart:      0,
				limitPageLength: 10,
			},
			wantErr: false,
			want:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ErpClient{
				appID:          tt.fields.appID,
				appSecret:      tt.fields.appSecret,
				baseURL:        tt.fields.baseURL,
				getCustomerURL: tt.fields.getCustomerURL,
				token: tt.fields.token,
			}
			got, err := c.GetCustomer(tt.args.limitStart, tt.args.limitPageLength)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Log(got)
		})
	}
}
