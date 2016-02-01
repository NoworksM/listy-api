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

package mal

import "encoding/xml"

// AnimeEntry to hold data from a users anime list
type AnimeEntry struct {
	XMLName               xml.Name `xml:"anime" json:"-"`
	ID                    uint     `xml:"series_animedb_id"`
	Title                 string   `xml:"series_title"`
	Synonyms              string   `xml:"series_synonyms"`
	Type                  string   `xml:"series_type"`
	Episodes              *uint    `xml:"series_episodes"`
	SeriesStatus          string   `xml:"series_status"`
	SeriesStartDate       string   `xml:"series_start"`
	SeriesEndDate         string   `xml:"series_end"`
	PosterURL             string   `xml:"series_image"`
	UserAnimeID           uint     `xml:"my_id"`
	UserEpisodesWatched   uint     `xml:"my_watched_episodes"`
	UserStartDate         string   `xml:"my_start_date"`
	UserFinishDate        string   `xml:"my_finish_date"`
	UserScore             float32  `xml:"my_score"`
	UserStatus            string   `xml:"my_status"`
	UserIsRewatching      *bool    `xml:"my_rewatching"`
	UserEpisodesRewatched uint     `xml:"my_rewatching_ep"`
	UserUpdatedAt         uint64   `xml:"my_last_updated"`
	UserTags              string   `xml:"my_tags"`
}
