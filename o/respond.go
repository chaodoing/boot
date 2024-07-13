package o

import (
	`encoding/xml`
)

type (
	Message[T any] struct {
		XMLName xml.Name `json:"-" xml:"root" yaml:"-"` // XML 名称，但在 JSON 和 YAML 中不使用。
		Code    int      `json:"code" xml:"code" yaml:"code"`
		Message string   `json:"message" xml:"message" yaml:"message"`
		Data    T        `json:"data" xml:"data" yaml:"data"`
	}
	Pagination[T any] struct {
		XMLName xml.Name `json:"-" xml:"root" yaml:"-"` // XML 名称，但在 JSON 和 YAML 中不使用。
		Code    int      `json:"code" xml:"code" yaml:"code"`
		Message string   `json:"message" xml:"message" yaml:"message"`
		Data    T        `json:"data" xml:"data" yaml:"data"`
		Current int      `json:"current" xml:"current" yaml:"current"`
		Total   int      `json:"total" xml:"total" yaml:"total"`
		Limit   int      `json:"limit" xml:"limit" yaml:"limit"`
	}
	Respond struct{}
)
