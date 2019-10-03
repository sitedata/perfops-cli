package cmd

import (
	"context"
	"github.com/ProspectOne/perfops-cli/cmd/internal"
	"github.com/ProspectOne/perfops-cli/perfops"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	countriesCmd = &cobra.Command{
		Use:     "countries",
		Short:   "Get a list of countries where PerfOps nodes are present",
		Long:    `Get a list of countries where PerfOps nodes are present`,
		Example: `perfops countries`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := newPerfOpsClient()
			if err != nil {
				return err
			}
			return chkRunError(runCountriesCmd(c))
		},
	}
)

func initCountriesCmd(parentCmd *cobra.Command) {
	parentCmd.AddCommand(countriesCmd)
}

func runCountriesCmd(c *perfops.Client) error {
	var res *[]perfops.Country

	ctx := context.Background()
	u := c.BasePath + "/analytics/dns/countries"

	f := internal.NewFormatter(debug && !outputJSON)
	f.StartSpinner()

	req, _ := http.NewRequest("GET", u, nil)
	req = req.WithContext(ctx)

	err := c.DoRequest(req, &res);
	f.StopSpinner()

	if err != nil {
		return err
	}

	internal.PrintOutputJSON(res)

	return nil
}
