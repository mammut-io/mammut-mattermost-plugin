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

	w.Header().Set("Content-Type", "application/json")

	switch path := r.URL.Path; path {
	case "/mammuthook":
		p.httpMeetingSettings(w, r)
	case "/":
		p.serveHTTPOriginal(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (p *Plugin) serveHTTPOriginal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func (p *Plugin) httpMeetingSettings(w http.ResponseWriter, r *http.Request) {

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
		ChannelId: mammutresponse.ChanelID,
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
