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
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/noworksm/listy-api/configuration"
	"github.com/noworksm/listy-api/dal"
	"github.com/noworksm/listy-api/logging"
	"github.com/noworksm/listy-api/parsers"
)

// initAnimeRoutes Initialize routes for the Anime API
func initAnimeRoutes(router *mux.Router) {
	router.HandleFunc("/anime/{animeId}", GetAnimeByID)
	router.HandleFunc("/users/{username}/anime", GetAnimeByUser)
}

// GetAnimeByID Handle Requests to get Anime by ID
func GetAnimeByID(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	rawID := vars["animeId"]
	animeID, err := strconv.Atoi(rawID)
	if err != nil {
		writeError(w, r, "An unknown error occured", err.Error())
		return
	}

	anime, err := dal.QueryAnimeByID(animeID)
	if err != nil && err != sql.ErrNoRows {
		logging.Error.Printf("An unknown error occured: %s", err.Error())
		writeError(w, r, "An unknown error occured", err.Error())
		return
	}
	if anime == nil {
		if config.EnvironmentDetail.Debug {
			logging.Debug.Printf("Anime %d not found, fetching data from server\n", animeID)
		}
		resp, err := http.Get(fmt.Sprintf("http://myanimelist.net/anime/%d", animeID))
		// TODO: Handle possible errors with MAL Requests
		fetched, err := parsers.ParseAnime(resp.Body)
		if err != nil {
			logging.Error.Print(err.Error())
			writeError(w, r, "An unknown error has occured", err.Error())
			return
		}
		_, err = dal.InsertAnime(&fetched)
		if err != nil {
			logging.Error.Print(err.Error())
			writeError(w, r, "An unknown error has occured", err.Error())
			return
		}
		anime, _ = dal.QueryAnimeByID(animeID)
	} else if time.Since(*anime.UpdatedAt).Hours() > 12 {
		resp, err := http.Get(fmt.Sprintf("http://myanimelist.net/anime/%d", animeID))
		// TODO: Handle possible errors with MAL Requests
		fetched, err := parsers.ParseAnime(resp.Body)
		if err != nil {
			logging.Error.Print(err.Error())
			writeError(w, r, "An unknown error has occured", err.Error())
			return
		}
		_, err = dal.UpdateAnime(&fetched)
		if err != nil {
			logging.Error.Print(err.Error())
			writeError(w, r, "An unknown error has occured", err.Error())
			return
		}
		anime, _ = dal.QueryAnimeByID(animeID)
	}

	writeResponse(w, r, anime)
	logging.Info.Printf("Get Anime request handled in %s", time.Since(start))
}

// GetAnimeByUser Get Anime on a users list
func GetAnimeByUser(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	username := vars["username"]
	resp, err := http.Get(fmt.Sprintf("http://myanimelist.net/malappinfo.php?u=%s&status=all&type=anime", username))
	if err != nil {
		writeError(w, r, "An unknown error occured", err.Error())
		return
	}

	_, entries, err := parsers.ParseAnimeList(resp.Body)

	writeResponse(w, r, entries)

	logging.Info.Printf("GET: /users/%s/anime request handled in %s", username, time.Since(start))
}
