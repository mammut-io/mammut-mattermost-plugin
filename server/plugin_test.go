package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"testing"
	"strings"
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/mattermost/mattermost-server/v5/model"
)


func TestServeHTTP(t *testing.T) {
	assert := assert.New(t)
	plugin := Plugin{}
	api := &plugintest.API{}
    api.On("LogInfo", "SERVE HTPP FUNCTION ACTIVATED").Return("SERVE HTPP FUNCTION ACTIVATED", nil)
	plugin.SetAPI(api)
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	plugin.ServeHTTP(nil, w, r)

	result := w.Result()
	assert.NotNil(result)
	bodyBytes, err := ioutil.ReadAll(result.Body)
	assert.Nil(err)
	bodyString := string(bodyBytes)

	assert.Equal("Hello, world!", bodyString)
}

func TestServeHTTP2(t *testing.T) {
	assert := assert.New(t)
	plugin := Plugin{}
	api := &plugintest.API{}
    api.On("LogInfo", "SERVE HTPP FUNCTION ACTIVATED").Return("SERVE HTPP FUNCTION ACTIVATED", nil)
    api.On("LogInfo", "MAMMUT response on hook","mammutresponseronhook", &MammutResponse{UserID:"user_id_1_placeholder", ChannelID:"chanel_id_placeholder", Message:"message_placeholder"}).Return(nil, nil)
    api.On("LogInfo", ">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>").Return(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", nil)
	api.On("CreatePost", &model.Post{
		UserId:    "user_id_1_placeholder",
		ChannelId: "chanel_id_placeholder",
		Message:   "message_placeholder",
	}).Return(&model.Post{
		UserId:    "user_id_1_placeholder",
		ChannelId: "chanel_id_placeholder",
		Message:   "message_placeholder",
	}, nil)
	plugin.SetAPI(api)
	testbody := &MammutResponse{
		UserID: "user_id_1_placeholder",
		ChannelID: "chanel_id_placeholder",
		Message: "message_placeholder",
	}
	jsonTest, err := json.Marshal(testbody)
	assert.Nil(err)


	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/mammuthook", strings.NewReader(string(jsonTest)))
	r.Header.Add("Mattermost-User-Id", "theuserid")

	plugin.ServeHTTP(nil, w, r)

	result := w.Result()
	assert.NotNil(result)
	status := result.StatusCode
	assert.Equal(200, status)
}
