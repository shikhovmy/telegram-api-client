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
	err := TelegramApi{}.ParseResponse(res, &Update{})
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
	user := api.GetMe()
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
	resp := api.GetSpecifiedUpdates(1, 1, 0)
	if len(resp) == 1 && resp[0].UpdateId != 934217543 {
		t.Error("fail", resp)
	}

}

type clientMock struct {
	Response *http.Response
}

func (client *clientMock) Get(endpoint string, queryParams map[string]string) (resp *http.Response, err error) {
	return client.Response, nil
}

func (client *clientMock) Post(endpoint string, queryParams map[string]string) (resp *http.Response, err error) {
	return client.Response, nil
}

func body(json string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(json)))
}
