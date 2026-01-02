/*
Copyright © 2026 Ronmi Ren

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

// Package komodo defines the types the komodo alert system uses.
package komodo

import (
	"encoding/json"
	"time"
)

type PayloadItem json.RawMessage

func (p *PayloadItem) UnmarshalJSON(buf []byte) error {
	*p = buf
	return nil
}
func (p PayloadItem) MarshalJSON() ([]byte, error) {
	return p, nil
}
func (p PayloadItem) IsDict() bool {
	return len(p) > 0 && p[0] == '{'
}
func (p PayloadItem) Dict() Map {
	if !p.IsDict() {
		return nil
	}
	var m map[string]PayloadItem
	err := json.Unmarshal(p, &m)
	if err != nil {
		return nil
	}
	return m
}

func (p PayloadItem) IsArray() bool {
	return len(p) > 0 && p[0] == '['
}
func (p PayloadItem) Array() []PayloadItem {
	if !p.IsArray() {
		return nil
	}
	var a []PayloadItem
	err := json.Unmarshal(p, &a)
	if err != nil {
		return nil
	}
	return a
}

func (p PayloadItem) IsStr() bool {
	if len(p) > 0 && p[0] == '"' {
		return true
	}
	if len(p) == 0 {
		return true
	}

	return false
}
func (p PayloadItem) Str() string {
	if !p.IsStr() {
		return ""
	}
	if len(p) == 0 {
		return ""
	}
	var s string
	err := json.Unmarshal(p, &s)
	if err != nil {
		return ""
	}
	return s
}

func (p PayloadItem) IsBool() bool {
	return string(p) == "true" || string(p) == "false" || len(p) == 0
}
func (p PayloadItem) Bool() bool {
	if !p.IsBool() {
		return false
	}
	var b bool
	err := json.Unmarshal(p, &b)
	if err != nil {
		return false
	}
	return b
}

func (p PayloadItem) IsNum() bool {
	return !p.IsDict() && !p.IsArray() && !p.IsStr() && !p.IsBool()
}
func (p PayloadItem) Num() float64 {
	if !p.IsNum() {
		return 0
	}
	var n float64
	err := json.Unmarshal(p, &n)
	if err != nil {
		return 0
	}
	return n
}
func (p PayloadItem) Int() int64 {
	if !p.IsNum() {
		return 0
	}
	var n int64
	err := json.Unmarshal(p, &n)
	if err != nil {
		return 0
	}
	return n
}

type Map map[string]PayloadItem

func (m Map) Get(key string) PayloadItem {
	return m[key]
}
func (m Map) Has(key string) bool {
	_, ok := m[key]
	return ok
}

type AlertTarget struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AlertData struct {
	Type    string `json:"type"`
	Payload Map    `json:"data"`
}

type AlertInfo struct {
	Timestamp        int64       `json:"ts"` // js timestamp in milliseconds
	Level            string      `json:"level"`
	Resolved         bool        `json:"resolved"`
	ResolveTimestamp int64       `json:"resolve_at"` // js timestamp in milliseconds
	Target           AlertTarget `json:"target"`
	Data             AlertData   `json:"data"`
}

var TZ = time.UTC

func (a *AlertInfo) IssuedAt() time.Time {
	return time.UnixMilli(a.Timestamp).In(TZ)
}

func (a *AlertInfo) ResolvedAt() time.Time {
	return time.UnixMilli(a.ResolveTimestamp).In(TZ)
}
