package entity

import "time"

type Resource struct {
	ID               string    `json:"id"`
	CreationTime     time.Time `json:"CreationTime,omitempty"`
	LastChange       time.Time `json:"LastChange,omitempty"`
	RelatedResources []string  `json:"RelatedResources,omitempty"`
	Risk             []string  `json:"Riks"`
	Tags             []string  `json:"Tags"`
}
