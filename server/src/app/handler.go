package app

import (
	"encoding/json"
	"flash/conf"
	"flash/ds"
	"flash/logger"
	"flash/util"
	"fmt"
	"html"
	"net/http"

	"github.com/sirupsen/logrus"
)

var config conf.Config = conf.Get()

func HomeReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func PingReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Boss is always fine")
}

func FindReq(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	searchBucket := ds.GetDs(ctx)
	setupResponse(&w)
	prefix := r.FormValue("prefix")

	// sanitize the incoming "prefix" query param
	sanitizedPrefix := util.SanitizeString(prefix)

	matches := searchBucket.Find(sanitizedPrefix)

	if config.LoggingEnabled == true {
		logger.Log.WithFields(logrus.Fields{
			"request_id": util.GetRequestID(ctx),
			// "time_elapsed": util.SetElapsedTimeStamp(ctx),
		}).Info("find_matches_received")

	}
	response := util.Response{
		Prefix:  sanitizedPrefix,
		Matches: matches,
	}
	json.NewEncoder(w).Encode(response)
}

func InsertReq(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	searchBucket := ds.GetDs(ctx)
	setupResponse(&w)
	prefix := r.FormValue("prefix")

	// sanitize the incoming "prefix" query param.
	sanitizedPrefix := util.SanitizeString(prefix)

	searchBucket.Insert(sanitizedPrefix)

	response := util.Response{
		Prefix: sanitizedPrefix,
	}
	json.NewEncoder(w).Encode(response)
}
