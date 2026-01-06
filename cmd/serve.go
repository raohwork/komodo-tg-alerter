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
package cmd

import (
	"encoding/json"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/raohwork/komodo-tg-alerter/config"
	"github.com/raohwork/komodo-tg-alerter/komodo"
	"github.com/raohwork/komodo-tg-alerter/tmpl"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Komodo Telegram Alerter server",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt)
		defer stop()

		cfg := config.NewConfig()
		if err := cfg.Validate(); err != nil {
			log.Fatal().Err(err).Msg("invalid configuration")
		}

		renderer := tmpl.NewRendererFromPath(cfg.CustemplatePath, cfg.Timezone())
		l, closeLogFile, err := cfg.GetLogger()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to initialize logger")
		}
		defer closeLogFile()
		log.Logger = l

		tgapi, err := bot.New(cfg.TelegramToken)
		if err != nil {
			l.Fatal().Err(err).Msg("failed to create telegram bot")
		}

		l.Info().Msg("Starting Komodo Telegram Alerter")
		srv := &http.Server{
			Addr: cfg.WebBind,
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var data komodo.AlertInfo
				err := json.NewDecoder(r.Body).Decode(&data)
				if err != nil {
					l.Error().Err(err).Msg("failed to decode request body")
					return
				}

				msg, err := renderer.Render(&data)
				if err != nil {
					l.Error().Err(err).Msg("failed to render message")
					return
				}

				l.Info().Msgf("Rendered message:\n%s", msg)

				_, err = tgapi.SendMessage(ctx, &bot.SendMessageParams{
					ChatID:    cfg.TelegramChatID,
					Text:      msg,
					ParseMode: models.ParseModeMarkdown,
				})
				if err != nil {
					l.Error().Err(err).Msg("failed to send telegram message")
					return
				}
			}),
		}
		go func() {
			srv.ListenAndServe()
			os.Exit(0)
		}()

		<-ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
