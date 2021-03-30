/*
Copyright © 2021 Mohammad Shoaei

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
	"github.com/spf13/cobra"
)

// newRootCommand represents the base command when called without any subcommands
func newRootCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "fxdl",
		Short: "fxdl downloads historical data of an instrument from https://forexsb.com",
		Long: `fxdl download maximum of 200,000 candles from https://forexsb.com 
and saves the data in either JSON or CSV format`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("fxdl called")
			return nil
		},
	}
	return cmd
}

func Execute() {
	cobra.CheckErr(newRootCommand().Execute())
}
