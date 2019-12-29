package directory

import (
	"log"
	"strconv"

	"github.com/occu-io/gsreport/src/utils"

	admin "google.golang.org/api/admin/directory/v1"
)

func getGroups(srv *admin.Service) []*admin.Group {
	r, err := srv.Groups.List().Customer("my_customer").MaxResults(100).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve groups in domain: %v", err)
	}

	// for _, g := range r.Groups {
	// log.Printf("%s: email=%s", g.Name, g.Email)
	// }

	return r.Groups
}

func GetGroupMembers(report *utils.Report) {
	srv := getService(report.Cl)
	gArr := getGroups(srv)

	report.Xls.File.NewSheet(report.Xls.Sheet)
	var itr int64 = 2

	report.Xls.File.SetCellValue(report.Xls.Sheet, "A1", "E-MAIL")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "B1", "ROLE")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "C1", "GROUP")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "D1", "STATUS")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "E1", "TYPE")

	for _, g := range gArr {
		r, err := srv.Members.List(g.Id).Do()
		if err != nil {
			log.Fatalf("Unable to retrieve members of group: %s in domain: %v", g.Name, err)
		}

		for _, gMem := range r.Members {
			s := strconv.FormatInt(itr, 10)

			if gMem.Email == "" {
				continue
			}

			report.Xls.File.SetCellValue(report.Xls.Sheet, "A"+s, gMem.Email)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "B"+s, gMem.Role)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "C"+s, g.Name)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "D"+s, gMem.Status)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "E"+s, gMem.Type)

			itr++
		}
	}
}

func GetGroups(report *utils.Report) {
	srv := getService(report.Cl)
	gArr := getGroups(srv)

	report.Xls.File.NewSheet(report.Xls.Sheet)
	var itr int64 = 2

	report.Xls.File.SetCellValue(report.Xls.Sheet, "A1", "GROUP")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "B1", "E_MAIL")

	for _, g := range gArr {
		s := strconv.FormatInt(itr, 10)

		report.Xls.File.SetCellValue(report.Xls.Sheet, "A"+s, g.Name)
		report.Xls.File.SetCellValue(report.Xls.Sheet, "B"+s, g.Email)

		itr++
	}

}
