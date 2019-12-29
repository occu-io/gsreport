package report

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/occu-io/gsreport/src/utils"
)

func GetAuditLogins(report *utils.Report) {
	report.Xls.File.NewSheet(report.Xls.Sheet)

	srv := getService(report.Cl)

	var itr int64 = 2

	userkey := report.Cfg.Section(report.Xls.Sheet).Key("userkey").String()
	appname := report.Cfg.Section(report.Xls.Sheet).Key("appname").String()

	r, err := srv.Activities.List(userkey, appname).MaxResults(100).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve logs for account. %v", err)
	}

	if len(r.Items) == 0 {
		fmt.Println("No logs found.")
	} else {
		report.Xls.File.SetCellValue(report.Xls.Sheet, "A1", "TIME")
		report.Xls.File.SetCellValue(report.Xls.Sheet, "B1", "E-MAIL")
		report.Xls.File.SetCellValue(report.Xls.Sheet, "C1", "IP_ADDRESS")
		report.Xls.File.SetCellValue(report.Xls.Sheet, "D1", "EVENT")
		report.Xls.File.SetCellValue(report.Xls.Sheet, "E1", "PARAM_NAME")
		report.Xls.File.SetCellValue(report.Xls.Sheet, "F1", "PARAM_VALUE")

		for _, a := range r.Items {
			t, err := time.Parse(time.RFC3339Nano, a.Id.Time)
			if err != nil {
				fmt.Println("Unable to parse login time.")
				// Set time to zero.
				t = time.Time{}
			}
			for _, e := range a.Events {
				s := strconv.FormatInt(itr, 10)

				report.Xls.File.SetCellValue(report.Xls.Sheet, "A"+s, t.Format(time.RFC822))
				report.Xls.File.SetCellValue(report.Xls.Sheet, "B"+s, a.Actor.Email)
				report.Xls.File.SetCellValue(report.Xls.Sheet, "C"+s, a.IpAddress)
				report.Xls.File.SetCellValue(report.Xls.Sheet, "D"+s, e.Name)

				for _, par := range e.Parameters {
					s := strconv.FormatInt(itr, 10)

					report.Xls.File.SetCellValue(report.Xls.Sheet, "E"+s, par.Name)
					report.Xls.File.SetCellValue(report.Xls.Sheet, "F"+s, par.Value)

					itr++
				}
			}
		}
	}
}
