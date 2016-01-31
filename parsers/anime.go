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

package parsers

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/noworksm/listy-api/logging"
	"github.com/noworksm/listy-api/models"

	"gopkg.in/xmlpath.v2"
)

const dateForm = "Jan 2, 2006"

var idPath = xmlpath.MustCompile("//input[@name='aid']")
var titlePath = xmlpath.MustCompile("//h1[@class='h1']/span[@itemprop='name']")
var descriptionPath = xmlpath.MustCompile("//span[@itemprop='description']")
var englishTitlePath = xmlpath.MustCompile("//div[span='English:']")
var englishTitleRegex = regexp.MustCompile("\\s*English:\\s*(.*)\\s*")
var japaneseTitlePath = xmlpath.MustCompile("//div[span='Japanese:']")
var japaneseTitleRegex = regexp.MustCompile("\\s*Japanese:\\s*(.*)\\s*")
var synonymsPath = xmlpath.MustCompile("//div[span='Synonyms:']")
var synonymsRegex = regexp.MustCompile("\\s*Synonyms:\\s*(.*)\\s*")
var typePath = xmlpath.MustCompile("//div[span='Type:']")
var typeRegex = regexp.MustCompile("\\s*Type:\\s*(.*)\\s*")
var episodesPath = xmlpath.MustCompile("//div[span='Episodes:']")
var episodesRegex = regexp.MustCompile("\\s*Episodes:\\s*(\\d{1,})\\s*")
var statusPath = xmlpath.MustCompile("//div[span='Status:']")
var statusRegex = regexp.MustCompile("\\s*Status:\\s*(.*)\\s*")
var premieredPath = xmlpath.MustCompile("//div[span='Premiered:']")
var premieredRegex = regexp.MustCompile("\\s*Premiered:\\s*(.*)\\s*")
var airedPath = xmlpath.MustCompile("//div[span='Aired:']")
var airedRegex = regexp.MustCompile("\\s*Aired:\\s*(.*)\\s*")

var scorePath = xmlpath.MustCompile("//div[span='Score:']")
var scoreRegex = regexp.MustCompile("\\s*Score:\\s*(\\d{1,2}\\.\\d{1,4}).*")
var rankedPath = xmlpath.MustCompile("//div[span='Ranked:']")
var rankedRegex = regexp.MustCompile("\\s*Ranked:\\s*#(\\d{1,})\\s*")
var popularityPath = xmlpath.MustCompile("//div[span='Popularity:']")
var popularityRegex = regexp.MustCompile("\\s*Popularity:\\s*#?(\\d{1,})\\s*")
var membersPath = xmlpath.MustCompile("//div[span='Members:']")
var membersRegex = regexp.MustCompile("\\s*Members:\\s*(.*)\\s*")
var favoritesPath = xmlpath.MustCompile("//div[span='Favorites:']")
var favoritesRegex = regexp.MustCompile("\\s*Favorites:\\s*(.*)\\s*")

// ParseAnime Parse an Anime's data from a given html string
func ParseAnime(htmlReader io.Reader) (models.Anime, error) {
	raw, err := ioutil.ReadAll(htmlReader)
	if err != nil {
		logging.Error.Print(err.Error())
		panic(err)
	}

	buf := bytes.NewBuffer(raw)

	doc, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		logging.Error.Printf("goQuery Error: %s", err.Error())
		panic(err)
	}

	idNode := doc.Find("input[name='aid']")
	idStr, exists := idNode.Attr("value")
	if !exists {
		return models.Anime{}, errors.New("Invalid Html")
	}
	id, _ := strconv.ParseUint(idStr, 10, 32)
	airedNode := doc.Find("span:contains(Aired)").Parent()
	aired := strings.TrimSpace(airedNode.Text())
	aired = strings.TrimPrefix(aired, "Aired:")

	buf = bytes.NewBuffer(raw)
	root, err := xmlpath.ParseHTML(buf)
	if err != nil {
		logging.Error.Printf("xmlpath.v2 error: %s", err.Error())
		panic(err)
	}

	title, _ := titlePath.String(root)
	englishTitle, _ := englishTitlePath.String(root)
	description, _ := descriptionPath.String(root)
	episodesStr, _ := episodesPath.String(root)
	scoreStr, _ := scorePath.String(root)
	typeStr, _ := typePath.String(root)
	status, _ := statusPath.String(root)
	premiered, _ := premieredPath.String(root)
	rankStr, _ := rankedPath.String(root)
	popularityStr, _ := popularityPath.String(root)
	favoritesStr, _ := favoritesPath.String(root)
	dates := strings.Split(aired, " to ")
	startDate, _ := time.Parse(dateForm, strings.TrimSpace(dates[0]))
	endDate, _ := time.Parse(dateForm, strings.TrimSpace(dates[1]))

	res := episodesRegex.FindStringSubmatch(episodesStr)
	episodes, _ := strconv.ParseUint(res[1], 10, 32)
	res = englishTitleRegex.FindStringSubmatch(englishTitle)
	if len(res) > 1 {
		englishTitle = res[1]
	}
	res = typeRegex.FindStringSubmatch(typeStr)
	typeStr = res[1]
	res = scoreRegex.FindStringSubmatch(scoreStr)
	score, _ := strconv.ParseFloat(res[1], 32)
	res = rankedRegex.FindStringSubmatch(rankStr)
	rank, _ := strconv.ParseUint(res[1], 10, 32)
	res = popularityRegex.FindStringSubmatch(popularityStr)
	popularity, _ := strconv.ParseUint(res[1], 10, 32)
	res = favoritesRegex.FindStringSubmatch(favoritesStr)
	favorites, _ := strconv.ParseUint(res[1], 10, 32)
	res = statusRegex.FindStringSubmatch(status)
	status = res[1]
	res = premieredRegex.FindStringSubmatch(premiered)
	premiered = res[1]

	now := time.Now()

	return models.Anime{
		ID:           uint(id),
		Title:        title,
		Description:  description,
		EnglishTitle: englishTitle,
		Episodes:     uint(episodes),
		Score:        float32(score),
		Type:         typeStr,
		Status:       status,
		Premiered:    premiered,
		Rank:         uint(rank),
		Popularity:   uint(popularity),
		StartDate:    startDate,
		EndDate:      endDate,
		Favorites:    uint(favorites),
		UpdatedAt:    now,
	}, nil
}
