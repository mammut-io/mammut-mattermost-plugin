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
    api.On("LogInfo", "EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR").Return("EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR", nil)
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
    api.On("LogInfo", "EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR").Return("EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR", nil)
	api.On("CreatePost", &model.Post{
		UserId:    "elieser",
		ChannelId: "Pereira",
		Message:   "Pereira",
	}).Return(&model.Post{
		UserId:    "elieser",
		ChannelId: "Pereira",
		Message:   "Pereira",
	}, nil)
	plugin.SetAPI(api)
	testbody := &MammutResponse{
		UserID: "elieser",
		ChannelID: "Pereira",
		Message: "Pereira",
	}
	jsonTest, err := json.Marshal(testbody)
	assert.Nil(err)


	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/mammuthook", strings.NewReader(string(jsonTest)))
	r.Header.Add("Mattermost-User-Id", "theuserid")

	plugin.ServeHTTP(nil, w, r)

	result := w.Result()
	assert.NotNil(result)
	//bodyBytes, err := ioutil.ReadAll(result.Body)
	//assert.Nil(err)
	//bodyString := string(bodyBytes)

	//assert.Equal(200, bodyString)
	status := result.StatusCode
	assert.Equal(200, status)
}
