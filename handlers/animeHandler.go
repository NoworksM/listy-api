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
	"io/ioutil"
	"net/http"
	"noworks/listy-mal-api/models"

	"github.com/gorilla/mux"
	"gopkg.in/redis.v3"
)

var redisClient *redis.Client

// InitAnimeRoutes Initialize routes for the Anime API
func InitAnimeRoutes(router *mux.Router, client *redis.Client) {
	redisClient = client
	router.HandleFunc("/anime/{animeId}", GetAnimeByID)
	router.HandleFunc("/users/{userId}/anime", GetAnimeByUser)
}

// GetAnimeByID Handle Requests to get Anime by ID
func GetAnimeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	animeID := vars["animeId"]

	anime, err := redisClient.Get("anime:" + animeID).Result()
	if err == redis.Nil {
		// TODO: Fetch new Item from server
		response, err := http.Get("http://myanimelist.net/anime/" + animeID)
		if err != nil {
			json.NewEncoder(w).Encode(models.Error{
				Message: "An unknown error occured",
				Reason:  err.Error(),
			})
			return
		}
		body, _ := ioutil.ReadAll(response.Body)
	} else if err != nil {
		json.NewEncoder(w).Encode(models.Error{
			Message: "An unknown error occured",
			Reason:  "",
		})
		return
	}
	json.NewEncoder(w).Encode(anime)
}

// GetAnimeByUser Get Anime on a users list
func GetAnimeByUser(w http.ResponseWriter, r *http.Request) {

}
