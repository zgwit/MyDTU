package models

import (
	"github.com/robertkrimen/otto"
	"time"
)

type Project struct {
	ProjectTemplate `storm:"inline"`

	//Disabled   bool  `json:"disabled"`
	TemplateId int   `json:"template_id"`
	LinkBinds  []int `json:"link_binds"`
}

type ProjectTemplate struct {
	ID          int    `json:"id"`
	UUID        string `json:"uuid" storm:"unique"` //唯一码，自动生成
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`

	Links      []ProjectLink      `json:"links"`
	Jobs       []ProjectJob       `json:"jobs"`
	Strategies []ProjectStrategy  `json:"strategies"`
	Validators []ProjectValidator `json:"validators"`

	Created time.Time `json:"created" storm:"created"`
	Updated time.Time `json:"updated" storm:"updated"`
}

type ProjectLink struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`

	//轮询
	LoopEnable   bool `json:"loop_enable"`
	LoopPeriod   int  `json:"loop_period"`   //ms
	LoopInterval int  `json:"loop_interval"` //ms

	Elements []ProjectElement `json:"elements"`
}

type ProjectElement struct {
	Element string `json:"element"` //uuid

	Name  string `json:"name"`
	Alias string `json:"alias"`
	Slave uint8  `json:"slave"` //从站号

	//轮询
	LoopTimes int `json:"loop_times"` //轮询多少次，才会检查一次
}

type ProjectValidator struct {
	Message string `json:"message"`
	Script  string `json:"script"`
}

type ProjectJob struct {
	Name   string `json:"name"`
	Cron   string `json:"cron"`
	Script string `json:"script"` //javascript
}

type ProjectStrategy struct {
	Name   string `json:"name"`
	Script string `json:"script"` //javascript

}

type Script struct {
	source    string
	variables []string
	script    *otto.Script
}
