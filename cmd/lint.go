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
	"fmt"
	"os"
	"time"

	"github.com/raohwork/komodo-tg-alerter/tmpl"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// lintCmd represents the lint command
var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint templates to check for errors",
	Run: func(cmd *cobra.Command, args []string) {
		w := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		l := zerolog.New(w).With().Timestamp().Logger().Level(zerolog.TraceLevel)
		log.Logger = l

		fmt.Println(tmpl.Lint(nil))
	},
}

func init() {
	rootCmd.AddCommand(lintCmd)
}
