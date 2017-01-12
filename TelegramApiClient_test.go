package TelegramApiClient

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestParseResponse(t *testing.T) {

	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       body(`{"ok": true, "result": []}`),
	}
	err := ParseResponse(res, &Update{})
	if err != nil {
		t.Error("not ok", err)
	}
}

func TestGetMe(t *testing.T) {
	cli := &clientMock{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body(`{"ok":true,"result":{"id":236079868,"first_name":"bot","username":"bot"}}`),
		},
	}
	api := TelegramApi{client: cli}
	user, _ := api.GetMe()
	if user.Id != 236079868 {
		t.Error("not ok", user)
	}
}

func TestGetUpdates(t *testing.T) {
	cli := &clientMock{
		Response: &http.Response{
			StatusCode: http.StatusOK,
			Body:       body(`{"ok":true,"result":[{"update_id":934217543,"message":{}}]}`),
		},
	}
	api := TelegramApi{client: cli}
	resp, _ := api.GetSpecifiedUpdates(1, 1, 0)
	if len(resp) == 1 && resp[0].UpdateId != 934217543 {
		t.Error("fail", resp)
	}

}

type clientMock struct {
	Response *http.Response
}

func (client *clientMock) Get(endpoint Endpoint, queryParams map[string]string, objectToReturn interface{}) (err error) {
	return nil
}

func (client *clientMock) Post(endpoint Endpoint, data interface{}, objectToReturn interface{}) (err error) {
	return nil
}

func body(json string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}
