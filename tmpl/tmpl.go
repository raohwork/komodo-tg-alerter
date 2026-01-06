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

// Package tmpl defines templates used in komodo.
package tmpl

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/go-telegram/bot"
	"github.com/raohwork/komodo-tg-alerter/komodo"
	"github.com/rs/zerolog/log"
)

//go:embed *.txt
var Files embed.FS

type Renderer struct {
	fs fs.FS
	tz *time.Location
}

func prepareTemplate(tz *time.Location) *template.Template {
	return template.New("").
		Funcs(template.FuncMap{
			"timefmt": func(t time.Time) string {
				return t.In(tz).Format("2006-01-02 15:04:05")
			},
			"escape": bot.EscapeMarkdown,
			"e":      bot.EscapeMarkdown, // short alias for escape
			"f": func(f float64) string {
				return bot.EscapeMarkdown(fmt.Sprintf("%.4f", f))
			},
		})
}

func NewRenderer(fs fs.FS, tz *time.Location) *Renderer {
	if fs == nil {
		fs = Files
	}
	return &Renderer{fs: fs, tz: tz}
}

// NewRendererFromPath creates a Renderer using a custom template path.
// If path is empty, it uses the embedded templates.
// If path is not empty, it uses os.DirFS to load templates from the filesystem.
func NewRendererFromPath(path string, tz *time.Location) *Renderer {
	if path == "" {
		return NewRenderer(nil, tz)
	}
	return NewRenderer(os.DirFS(path), tz)
}

func (r Renderer) Render(data *komodo.AlertInfo) (string, error) {
	log.Info().
		Interface("data", data).
		Str("type", data.Data.Type).
		Msg("rendering template")

	typ := data.Data.Type
	t, err := prepareTemplate(r.tz).ParseFS(r.fs, typ+".txt")
	if err != nil {
		return "", fmt.Errorf("parse template %s: %w", typ, err)
	}

	var buf strings.Builder
	err = t.ExecuteTemplate(&buf, typ+".txt", data)
	if err != nil {
		return "", fmt.Errorf("execute template %s: %w", typ, err)
	}

	return buf.String(), nil
}
