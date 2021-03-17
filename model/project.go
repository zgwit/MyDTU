package model

import (
	"time"
)

type Project struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`

	Disabled bool `json:"disabled"`

	Manifest TemplateManifest `json:"manifest" xorm:"json"`

	//模板项目
	IsTemplate bool  `json:"is_template"`
	TemplateId int64 `json:"template_id"`

	Created time.Time `json:"created" xorm:"created"`
}
