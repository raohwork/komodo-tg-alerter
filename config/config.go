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

// Package config defines all possible configuration options.
package config

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken   string
	TelegramChatID  int64
	WebBind         string
	CustemplatePath string
	LogLevel        string
	LogFile         string
}

func (c *Config) Validate() error {
	if c.TelegramToken == "" {
		return errors.New("telegram.token is not set")
	}
	if c.TelegramChatID == 0 {
		return errors.New("telegram.chat is not set")
	}

	_, err := zerolog.ParseLevel(c.LogLevel)
	if err != nil {
		return errors.New("log.level is invalid")
	}

	return nil
}

func (c *Config) GetLogger() (logger zerolog.Logger, close func(), err error) {
	level, _ := zerolog.ParseLevel(c.LogLevel)
	var w io.Writer
	w = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	close = func() {}

	if c.LogFile != "" {
		f, err := os.OpenFile(c.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return logger, close, err
		}
		w = io.MultiWriter(w, zerolog.SyncWriter(f))
		close = func() {
			f.Close()
		}
	}

	logger = zerolog.New(w).With().Timestamp().Logger().Level(level)
	return
}

func NewConfig() *Config {
	viper.SetDefault("web.bind", ":8964")
	viper.SetDefault("log.level", "info")
	return &Config{
		TelegramToken:   viper.GetString("telegram.token"),
		TelegramChatID:  viper.GetInt64("telegram.chat"),
		WebBind:         viper.GetString("web.bind"),
		CustemplatePath: viper.GetString("template.path"),
		LogLevel:        viper.GetString("log.level"),
		LogFile:         viper.GetString("log.file"),
	}
}
