package directory

import (
	"log"
	"strconv"

	"github.com/occu-io/gsreport/src/utils"
)

func GetTokensThirdPartyApps(report *utils.Report) {
	report.Xls.File.NewSheet(report.Xls.Sheet)
	var itr int64 = 2
	srv := getService(report.Cl)
	userkey := report.Cfg.Section(report.Xls.Sheet).Key("userKey").String()

	r, err := srv.Tokens.List(userkey).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve tokens in domain: %v", err)
	}

	report.Xls.File.SetCellValue(report.Xls.Sheet, "A1", "APPLICATION_NAME")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "B1", "CLIENT_ID")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "C1", "ANONYMOUS")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "D1", "NATIVE_APP")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "E1", "SCOPES")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "F1", "USER")

	for _, token := range r.Items {
		s := strconv.FormatInt(itr, 10)

		user, err := srv.Users.Get(token.UserKey).Do()
		if err != nil {
			log.Fatalf("Unable to get user in domain: %v", err)
		}

		report.Xls.File.SetCellValue(report.Xls.Sheet, "A"+s, token.DisplayText)
		report.Xls.File.SetCellValue(report.Xls.Sheet, "B"+s, token.ClientId)
		report.Xls.File.SetCellValue(report.Xls.Sheet, "C"+s, token.Anonymous)
		report.Xls.File.SetCellValue(report.Xls.Sheet, "D"+s, token.NativeApp)
		report.Xls.File.SetCellValue(report.Xls.Sheet, "E"+s, token.Scopes)
		report.Xls.File.SetCellValue(report.Xls.Sheet, "F"+s, user.Name.FullName)

		itr++
	}
}
