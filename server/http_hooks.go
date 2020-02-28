package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println("can see this on test runnig")
	p.API.LogInfo(
		"EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR",
	)

	w.Header().Set("Content-Type", "application/json")
	//updateURL := "app:mammut-1/graph/user:[userId]?mattermost-user-id=[\"(mattermostdeployment.com,bot-identifier-2)\"]"
	updateURL := "/app:mammut-1/graph/user:1234"

	switch path := r.URL.Path; path {
	case "/mammuthooktemporal":
		p.httpMeetingSettingsTemporal(w, r)
	case updateURL:
		p.httpMeetingSettingsTemporal(w, r)
	case "/mammuthook":
		p.httpMeetingSettings(w, r)
	case "/":
		p.serveHTTPOriginal(w, r)
	default:
		p.API.LogInfo(
			">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>",
		)
		p.API.LogInfo(
			"entra en el case",
			"updateURL", updateURL,
			"r.URL.Path", r.URL.Path,
			"r.URL.query", r.URL.Query(),
		)
		http.NotFound(w, r)
	}
}

func (p *Plugin) serveHTTPOriginal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func (p *Plugin) httpMeetingSettings(w http.ResponseWriter, r *http.Request) {
	p.API.LogInfo(
		"EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR",
	)

	//not available if not authenticated, we remove this simple validation
	mattermostUserID := r.Header.Get("Mattermost-User-Id")
	//if mattermostUserID == "" {
	//	http.Error(w, "Not Authorized", http.StatusUnauthorized)
	//}

	switch r.Method {
	case http.MethodPost:
		p.httpMeetingSaveSettings(w, r, mattermostUserID)
	default:
		http.Error(w, "Request: "+r.Method+" is not allowed.", http.StatusMethodNotAllowed)
	}
	//fmt.Fprint(w, "Hello, world!")
}

func (p *Plugin) httpMeetingSaveSettings(w http.ResponseWriter, r *http.Request, mmUserID string) {

	//userID := r.Header.Get("Mattermost-User-ID")
	//if userID == "" {
	//	http.Error(w, "Not authorized", http.StatusUnauthorized)
	//	return
	//}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var mammutresponse *MammutResponse
	if err = json.Unmarshal(body, &mammutresponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, posterr := p.API.CreatePost(&model.Post{
		UserId:    mammutresponse.UserID,
		ChannelId: mammutresponse.ChannelID,
		Message:   mammutresponse.Message,
	})
	if posterr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	//fmt.Fprint(w, "Hello, world!")
}

func (p *Plugin) httpMeetingSettingsTemporal(w http.ResponseWriter, r *http.Request) {
	mattermostUserID := r.Header.Get("Mattermost-User-Id")

	switch r.Method {
	case http.MethodPost:
		p.httpMeetingSaveSettingsTemporal(w, r, mattermostUserID)
		p.API.LogInfo(
			"EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR",
		)
	default:
		p.API.LogInfo(
			"EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR",
		)
		http.Error(w, "Request: "+r.Method+" is not allowed.", http.StatusMethodNotAllowed)
	}
}

func (p *Plugin) httpMeetingSaveSettingsTemporal(w http.ResponseWriter, r *http.Request, mmUserID string) {
	p.API.LogInfo(
		"EEEEEEEEEEEELLLLLLLLIIIIIIIIIIIEESEEEEEEEEEEERRRRRRRR",
	)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.API.LogInfo(
		">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>",
	)
	p.API.LogInfo(
		"query on serving http",
		"r.URL.query", r.URL.Query(),
		"r.URL.queryuserid", r.URL.Query()["mattermost-user-id"],
		"r.URL.queryuseridindex0", r.URL.Query()["mattermost-user-id"][0],
	)
	p.API.LogInfo(
		"bodyOnTemporal",
		"body", body,
	)
	p.API.LogInfo(
		"<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<",
	)
	trb := &TaskResultBasic{
		AffectedElementID:   1234,
		AffectedElementName: "mammutname",
		AffectedElementType: "machine",
		TaskIDList:          make([]interface{}, 0)}
	au := &MammutUserCreationResponse{
		Status:     "success",
		Taskresult: []TaskResultBasic{*trb}}
	//fmt.Println(trb)
	//fmt.Println(au)
	//fmt.Println(au.Taskresult[0].AffectedElementID)ed
	resp, err := json.Marshal(au)
	//w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
