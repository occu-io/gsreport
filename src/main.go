package main

import (
	"flag"
	"log"

	"gopkg.in/ini.v1"

	"github.com/occu-io/gsreport/src/gsuite"
	"github.com/occu-io/gsreport/src/output"
	"github.com/occu-io/gsreport/src/utils"
)

func main() {
	var configPath string
	var report *utils.Report = &utils.Report{}

	flag.StringVar(&configPath, "config", "conf/svc.ini", "path to config file")
	flag.Parse()

	report.Cfg = utils.IniParse(configPath)

	// Fill authentication variables
	credentialsFile := report.Cfg.Section("global").Key("credentialsFile").String()
	tokenFile := report.Cfg.Section("global").Key("tokenFile").String()

	// authenticate & update orgUnitPath
	report.Cl = gsuite.GetClient(credentialsFile, tokenFile)

	report.Xls = output.NewExcel()

	for _, section := range report.Cfg.Sections() {
		report.Xls.Sheet = section.Name()

		if section.Name() == ini.DEFAULT_SECTION || section.Name() == "global" {
			continue
		}

		task := section.Key("task").In("GetUsersWithout2Fa", []string{
			"GetUsersWithout2Fa",
			"GetSuspendedUsers",
			"GetAdmins",
			"GetAsps",
			"GetGroups",
			"GetGroupMembers",
			"GetAuditLogins",
			"GetTokensThirdPartyApps",
		})

		gsuite.Functions[task](report)
	}

	xlsReportFilePath := report.Cfg.Section("global").Key("output").String()
	report.Xls.File.SaveAs(xlsReportFilePath)

	log.Printf("Done!\n")
}
