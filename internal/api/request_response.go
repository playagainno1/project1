package api

import (
	"context"
	"fmt"
	. "taylor-ai-server/internal/domain"

	"github.com/pkg/errors"
)

type EmptyRequest struct{}

type EmptyResponse struct{}

type ProfileRequest struct {
	Auto bool `form:"auto"`
}

func (r *ProfileRequest) Parse(ctx context.Context) error {
	if !r.Auto {
		return errors.Wrap(ErrBadRequest, "auto is invalid")
	}
	return nil
}

type ProfileResponse struct {
	ID                string           `json:"id"`
	Nickname          string           `json:"nickname"`
	DeviceID          string           `json:"deviceId"`
	RankAbbr          string           `json:"rankAbbr"`
	ProInfo           ProInfoView      `json:"proInfo"`
	ProProducts       []ProProductView `json:"proProducts"`
	UserTermsURL      string           `json:"userTermsUrl"`
	PrivacyPolicyURL  string           `json:"privacyPolicyUrl"`
	Tests             []TestView       `json:"tests"`
	ShareWatermarkURL string           `json:"shareWatermarkUrl"`
	HomeBgURL         string           `json:"homeBgUrl"`
}

func NewProfileResponse() *ProfileResponse {
	res := &ProfileResponse{
		// ID:       utils.NewULID(),
		ID:       "01hps756fs2b859vbf35amqmxy",
		Nickname: "Test01",
		// DeviceID: utils.NewDeviceID(),
		DeviceID: "876a7a44-a0a1-44a4-92d3-88dde2698b66",
		RankAbbr: "999+",
		ProInfo: ProInfoView{
			IsPro:    false,
			ProStart: 0,
			ProEnd:   0,
		},
		ProProducts: []ProProductView{
			{
				ProductID: "sub_monthly_001",
				Name:      "Monthly Subscription",
				Discount:  "50%",
			},
			{
				ProductID: "sub_yearly_001",
				Name:      "Yearly Subscription",
				Discount:  "50%",
			},
			{
				ProductID: "sub_weekly_001",
				Name:      "Weekly Subscription",
				Discount:  "50%",
			},
		},
		UserTermsURL:     "https://dev-t.float4ai.com/terms",
		PrivacyPolicyURL: "https://dev-t.float4ai.com/privacy",
		Tests: []TestView{
			{
				ID:      "test_001",
				Subject: "Test Subject 001",
				Hint:    "Test",
				Answer:  "Test 001",
			},
			{
				ID:      "test_002",
				Subject: "Test Subject 002",
				Hint:    "Test",
				Answer:  "Test 002",
			},
			{
				ID:      "test_003",
				Subject: "Test Subject 003",
				Hint:    "Test",
				Answer:  "Test 003",
			},
			{
				ID:      "test_004",
				Subject: "Test Subject 004",
				Hint:    "Test",
				Answer:  "Test 004",
			},
			{
				ID:      "test_005",
				Subject: "Test Subject 005",
				Hint:    "Test",
				Answer:  "Test 005",
			},
			{
				ID:      "test_006",
				Subject: "Test Subject 006",
				Hint:    "Test",
				Answer:  "Test 006",
			},
		},
		ShareWatermarkURL: "https://dev-t.float4ai.com/share_watermark.png",
		HomeBgURL:         "https://dev-t.float4ai.com/home_bg.png",
	}
	return res
}

type ProInfoView struct {
	IsPro    bool  `json:"isPro"`
	ProStart int64 `json:"proStart"`
	ProEnd   int64 `json:"proEnd"`
}

type ProProductView struct {
	ProductID string `json:"productId"`
	Name      string `json:"name"`
	Discount  string `json:"discount"`
}

type TestView struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	Hint    string `json:"hint"`
	Answer  string `json:"answer"`
}

type RanksResponse struct {
	Ranks []RankView `json:"ranks"`
}

func NewRanksResponse() *RanksResponse {
	res := &RanksResponse{Ranks: []RankView{}}
	for i := 0; i < 1000; i++ {
		rank := RankView{
			Rank:     i + 1,
			Nickname: fmt.Sprintf("Test%03d", i+1),
			Hot:      1000 - i,
		}
		res.Ranks = append(res.Ranks, rank)
	}
	return res
}

type RankView struct {
	Rank     int    `json:"rank"`
	Nickname string `json:"nickname"`
	Hot      int    `json:"hot"`
}
