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

package dal

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/noworksm/listy-api/models"
)

const tableName = "anime"
const insertParams = "id, title, english_title, description, episodes, score, type, status, premiered, rank, popularity, start_date, end_date, favorites"
const selectParams = "id, title, english_title, description, episodes, score, type, status, premiered, rank, popularity, start_date, end_date, favorites, created_at, updated_at"

// QueryAnimeByID Query an Anime object by ID
func QueryAnimeByID(id int) (*models.Anime, error) {
	row := Connection.QueryRow("SELECT "+selectParams+" FROM anime WHERE id = $1", id)
	return ReadAnimeRow(row)
}

// InsertAnime Insert new anime object into the database
func InsertAnime(anime *models.Anime) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		tableName,
		insertParams,
	)
	return Connection.Exec(
		query,
		anime.ID,
		anime.Title,
		anime.EnglishTitle,
		anime.Description,
		anime.Episodes,
		anime.Score,
		anime.Type,
		anime.Status,
		anime.Premiered,
		anime.Rank,
		anime.Popularity,
		anime.StartDate,
		anime.EndDate,
		anime.Favorites,
	)
}

// UpdateAnime Update a single Anime with the data from this new entry
func UpdateAnime(anime *models.Anime) (sql.Result, error) {
	query := fmt.Sprintf("UPDATE %s SET title = $2, english_title = $3,"+
		"description = $4, episodes =  $5, score = $6, type = $7, status = $8,"+
		"premiered = $9, rank = $10, popularity = $11, start_date = $12,"+
		"end_date = $13, favorites = $14, updated_at = $15 WHERE id = $1",
		tableName,
	)
	return Connection.Exec(
		query,
		anime.ID,
		anime.Title,
		anime.EnglishTitle,
		anime.Description,
		anime.Episodes,
		anime.Score,
		anime.Type,
		anime.Status,
		anime.Premiered,
		anime.Rank,
		anime.Popularity,
		anime.StartDate,
		anime.EndDate,
		anime.Favorites,
		time.Now(),
	)
}

// ReadAnimeRow Read an Anime object in from a Database Row
func ReadAnimeRow(row *sql.Row) (*models.Anime, error) {
	anime := new(models.Anime)
	err := row.Scan(&anime.ID, &anime.Title, &anime.EnglishTitle, &anime.Description,
		&anime.Episodes, &anime.Score, &anime.Type, &anime.Status,
		&anime.Premiered, &anime.Rank, &anime.Popularity, &anime.StartDate,
		&anime.EndDate, &anime.Favorites, &anime.CreatedAt, &anime.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return anime, err
}
