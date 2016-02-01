package mal

import "encoding/xml"

// UserAnimeInfo Information about totals for a user for different categories
type UserAnimeInfo struct {
	XMLName          xml.Name `xml:"myinfo"`
	ID               uint     `xml:"user_id"`
	Username         string   `xml:"user_name"`
	WatchingCount    uint     `xml:"user_watching"`
	CompletedCount   uint     `xml:"user_completed"`
	OnHoldCount      uint     `xml:"user_onhold"`
	DroppedCount     uint     `xml:"user_dropped"`
	PlanToWatchCount uint     `xml:"user_plantowatch"`
	DaysWatched      float32  `xml:"user_days_spent_watching"`
}
