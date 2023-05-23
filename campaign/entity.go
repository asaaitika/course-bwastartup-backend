package campaign

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	GoalAmout        int
	CurrentAmout     int
	Perks            string
	Slug             string
	BackerCount      int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImages
}

type CampaignImages struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
