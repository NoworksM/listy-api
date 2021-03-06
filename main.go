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
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/noworksm/listy-api/configuration"
	"github.com/noworksm/listy-api/dal"
	"github.com/noworksm/listy-api/handlers"
	"github.com/noworksm/listy-api/logging"

	_ "github.com/lib/pq"
)

func main() {
	logging.Init()
	config.InitConfig("config.json")

	conn, err := sql.Open("postgres", config.EnvironmentDetail.ConnectionString)
	if err != nil {
		panic(err)
	}
	dal.Connection = conn
	router := mux.NewRouter().StrictSlash(true)
	handlers.InitAnimeRoutes(router)

	log.Fatal(http.ListenAndServe(config.EnvironmentDetail.Domain, router))
}
