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

package main

import (
	"log"
	"net/http"
	"noworks/listy-mal-api/handlers"

	"github.com/gorilla/mux"
	"gopkg.in/redis.v3"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "192.168.99.100:32769",
		Password: "",
		DB:       0,
	})
	router := mux.NewRouter().StrictSlash(true)
	handlers.InitAnimeRoutes(router, redisClient)

	log.Fatal(http.ListenAndServe(":8080", router))
}
