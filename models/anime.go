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

package models

import "time"

// Anime Struct to represent an Anime data object
type Anime struct {
	ID           uint       `json:"id"`
	Title        string     `json:"title"`
	EnglishTitle string     `json:"englishTitle"`
	Description  string     `json:"description"`
	Episodes     *uint      `json:"episodes"`
	Score        float32    `json:"score"`
	Type         string     `json:"type"`
	Status       string     `json:"status"`
	Premiered    string     `json:"premiered"`
	Rank         uint       `json:"rank"`
	Popularity   uint       `json:"popularity"`
	StartDate    *time.Time `json:"startDate"`
	EndDate      *time.Time `json:"endDate"`
	Favorites    uint       `json:"favorites"`
	CreatedAt    *time.Time `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
}
