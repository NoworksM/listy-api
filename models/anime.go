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

import (
	"encoding/xml"
	"time"
)

// Anime Struct to represent an Anime data object
type Anime struct {
	XMLName      xml.Name   `json:"-" xml:"anime"`
	ID           uint       `json:"id" xml:"id"`
	Title        string     `json:"title" xml:"title"`
	EnglishTitle string     `json:"englishTitle" xml:"english-title"`
	Description  string     `json:"description" xml:"description"`
	Episodes     *uint      `json:"episodes" xml:"episodes"`
	Score        float32    `json:"score" xml:"score"`
	Type         string     `json:"type" xml:"type"`
	Status       string     `json:"status" xml:"status"`
	Premiered    string     `json:"premiered" xml:"premiered"`
	Rank         uint       `json:"rank" xml:"rank"`
	Popularity   uint       `json:"popularity" xml:"popularity"`
	StartDate    *time.Time `json:"startDate" xml:"start-date"`
	EndDate      *time.Time `json:"endDate" xml:"end-date"`
	Favorites    uint       `json:"favorites" xml:"favorites"`
	CreatedAt    *time.Time `json:"createdAt" xml:"created-at"`
	UpdatedAt    *time.Time `json:"updatedAt" xml:"updated-at"`
}
