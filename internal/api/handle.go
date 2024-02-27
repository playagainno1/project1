package api

import "context"
import . "taylor-ai-server/internal/domain"

type Profile struct {
}

func newProfile() *Profile {
	return &Profile{}
}

func (h *Profile) Handle(ctx context.Context, u User) (*ProfileResponse, error) {
	return NewProfileResponse(), nil
}

type Ranks struct {
}

func newRanks() *Ranks {
	return &Ranks{}
}

func (h *Ranks) Handle(ctx context.Context, u User) (*RanksResponse, error) {
	return NewRanksResponse(), nil
}
