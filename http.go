package main

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"github.com/aude/newman-webservice/newman"
)

func helpHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add help message
	http.Error(w, "not implemented: help", http.StatusNotImplemented)
}

func collectionsHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

var validId     = regexp.MustCompile("^[0-9A-Za-z-]+$")
var validAPIKey = regexp.MustCompile("^[0-9a-z]+$")

func getQueryParameter(_url *url.URL, name string) (value string, err error) {
	query, err := url.ParseQuery(_url.RawQuery)
	if err != nil {
		return "", err
	}
	if _, ok := query[name]; !ok {
		return "", fmt.Errorf("query parameter missing: %s", name)
	}
	if len(query[name]) > 1 {
		return "", fmt.Errorf("more than one query parameter with the same name: %s", name)
	}

	value = query[name][0]

	return value, nil
}

func collectionHandler(w http.ResponseWriter, r *http.Request) {
	// get collection ID
	collectionID := r.URL.Path[len("/collections/"):]
	if ! validId.MatchString(collectionID) {
		http.Error(w, "invalid collection ID", http.StatusBadRequest)
		return
	}

	// get environment ID
	environmentID, err := getQueryParameter(r.URL, "environment")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ! validId.MatchString(environmentID) {
		http.Error(w, "invalid environment ID", http.StatusBadRequest)
		return
	}

	// get API key
	apiKey, err := getQueryParameter(r.URL, "apikey")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if ! validAPIKey.MatchString(apiKey) {
		http.Error(w, "invalid API key", http.StatusBadRequest)
		return
	}

	// invoke Newman
	report, err := newman.JUnitReport(collectionID, environmentID, apiKey)
	if err != nil {
		http.Error(w, fmt.Sprintf("error invoking Newman: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, report)
}
