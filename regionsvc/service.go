package regionsvc

import "context"

type Region struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type Service interface {
	PostRegion(ctx context.Context, r Region) error
}
