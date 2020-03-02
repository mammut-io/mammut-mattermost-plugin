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
		"SERVE HTPP FUNCTION ACTIVATED",
	)

	w.Header().Set("Content-Type", "application/json")
	//updateURL := "app:mammut-1/graph/user:[userId]?mattermost-user-id=[\"(mattermostdeployment.com,bot-identifier-2)\"]"
	updateURL := "/app:mammut-1/graph/user:1234"

	switch path := r.URL.Path; path {
	case "/mammuthook":
		p.httpMammutHooks(w, r)
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

//httpMammutHooks valida si el metodo es post e invoca funcion que crea el post en canal correspondiente
func (p *Plugin) httpMammutHooks(w http.ResponseWriter, r *http.Request) {
	p.API.LogInfo(
		"HTTPMAMMUTHOOKS FUNCTION ACTIVATED",
	)

	//not available if not authenticated, we remove this simple validation
	mattermostUserID := r.Header.Get("Mattermost-User-Id")
	//if mattermostUserID == "" {
	//	http.Error(w, "Not Authorized", http.StatusUnauthorized)
	//}

	switch r.Method {
	case http.MethodPost:
		p.httpMammutHooksPostAction(w, r, mattermostUserID)
	default:
		http.Error(w, "Request: "+r.Method+" is not allowed.", http.StatusMethodNotAllowed)
	}
}

//httpMammutHooksPostAction crea el post al canal cuando mammutAPI envia respuesta de la conversacion
func (p *Plugin) httpMammutHooksPostAction(w http.ResponseWriter, r *http.Request, mmUserID string) {

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
	p.API.LogInfo(
		">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>",
	)
	p.API.LogInfo(
		"MAMMUT response on hook",
		"mammutresponseronhook", mammutresponse,
	)

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
