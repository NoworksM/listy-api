/**
 * Copyright 2016 Michael Nowak
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handlers

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/noworksm/listy-api/models"
)

// InitRoutes for the API
func InitRoutes(router *mux.Router) {
	initAnimeRoutes(router)
	initUserRoutes(router)
}

func writeResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	acceptHeaders := r.Header.Get("Accept")
	acceptTypes := strings.Split(acceptHeaders, ",")
	for _, header := range acceptTypes {
		if header == "application/xml" || header == "text/xml" {
			xml.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, r *http.Request, message string, reason string) {
	apiError := models.Error{Message: message, Reason: reason}
	acceptHeaders := r.Header.Get("Accept")
	acceptTypes := strings.Split(acceptHeaders, ",")
	for _, header := range acceptTypes {
		if header == "application/xml" || header == "text/xml" {
			xml.NewEncoder(w).Encode(apiError)
			return
		}
	}
	json.NewEncoder(w).Encode(apiError)
}
