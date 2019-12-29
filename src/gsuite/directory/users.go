package directory

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/occu-io/gsreport/src/utils"
)

// func UpdateOrgUnitPath(client *http.Client, cfg *ini.File) {
// 	// Fill GSuite Organization Unit variables for migration
// 	oldOrgUnitPath := cfg.Section("users").Key("oldOrgUnitPath").String()
// 	newOrgUnitPath := cfg.Section("users").Key("newOrgUnitPath").String()
// 	whitelist := cfg.Section("users").Key("whitelist").Strings(",")

// 	srv := getService(client)
// 	r := getUserList(srv, "orgUnitPath='"+oldOrgUnitPath+"'")
// 	if len(r.Users) == 0 {
// 		fmt.Print("No users found.\n")
// 	} else {
// 		sort.Strings(whitelist)

// 		for _, u := range r.Users {
// 			if utils.BinarySearch(whitelist, u.PrimaryEmail) {
// 				log.Printf("Found %s in whitelist, skipping..\n", u.PrimaryEmail)
// 			} else {
// 				u.OrgUnitPath = newOrgUnitPath
// 				u, err := srv.Users.Update(u.PrimaryEmail, u).Do()
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				log.Printf("%s: user=%s, oldOrgUnitPath=%s, newOrgUnitPath=%s", u.PrimaryEmail, u.Name.FullName, oldOrgUnitPath, newOrgUnitPath)
// 			}
// 		}
// 	}
// }

func GetUsersWithout2Fa(report *utils.Report) {
	report.Xls.File.NewSheet(report.Xls.Sheet)

	// * TODO: prerobit parsing ini section "users"
	whitelist := report.Cfg.Section("users").Key("whitelist").Strings(",")

	var itr int64 = 2

	srv := getService(report.Cl)
	r := getUserList(srv, "isEnrolledIn2Sv=false")

	report.Xls.File.SetCellValue(report.Xls.Sheet, "A1", "E-MAIL")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "B1", "NAME")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "C1", "IS_ENROLLED_IN_2_SV")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "D1", "SUSPENDED")

	if len(r.Users) == 0 {
		fmt.Print("No users found.\n")
	} else {
		sort.Strings(whitelist)

		for _, u := range r.Users {
			if !utils.BinarySearch(whitelist, u.PrimaryEmail) {
				s := strconv.FormatInt(itr, 10)

				report.Xls.File.SetCellValue(report.Xls.Sheet, "A"+s, u.PrimaryEmail)
				report.Xls.File.SetCellValue(report.Xls.Sheet, "B"+s, u.Name.FullName)
				report.Xls.File.SetCellValue(report.Xls.Sheet, "C"+s, u.IsEnrolledIn2Sv)
				report.Xls.File.SetCellValue(report.Xls.Sheet, "D"+s, u.Suspended)

				itr++
			}
		}
	}
}

func GetSuspendedUsers(report *utils.Report) {
	report.Xls.File.NewSheet(report.Xls.Sheet)

	var itr int64 = 2

	srv := getService(report.Cl)
	r := getUserList(srv, "isSuspended=true")

	report.Xls.File.SetCellValue(report.Xls.Sheet, "A1", "E-MAIL")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "B1", "NAME")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "C1", "SUSPENDED")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "D1", "SUSPENSION_REASON")

	if len(r.Users) == 0 {
		fmt.Print("No users found.\n")
	} else {
		for _, u := range r.Users {
			s := strconv.FormatInt(itr, 10)

			report.Xls.File.SetCellValue(report.Xls.Sheet, "A"+s, u.PrimaryEmail)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "B"+s, u.Name.FullName)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "C"+s, u.Suspended)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "D"+s, u.SuspensionReason)

			itr++
		}
	}
}

func GetAdmins(report *utils.Report) {
	report.Xls.File.NewSheet(report.Xls.Sheet)

	var itr int64 = 2

	srv := getService(report.Cl)
	r := getUserList(srv, "isDelegatedAdmin=true")

	report.Xls.File.SetCellValue(report.Xls.Sheet, "A1", "E-MAIL")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "B1", "NAME")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "C1", "IS_ENROLLED_IN_2_SV")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "D1", "SUSPENDED")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "E1", "DELEGATED_ADMIN")
	report.Xls.File.SetCellValue(report.Xls.Sheet, "F1", "ADMIN")

	if len(r.Users) == 0 {
		fmt.Print("No users found.\n")
	} else {
		for _, u := range r.Users {
			s := strconv.FormatInt(itr, 10)

			report.Xls.File.SetCellValue(report.Xls.Sheet, "A"+s, u.PrimaryEmail)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "B"+s, u.Name.FullName)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "C"+s, u.IsEnrolledIn2Sv)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "D"+s, u.Suspended)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "E"+s, u.IsDelegatedAdmin)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "F"+s, u.IsAdmin)

			itr++
		}
	}

	r = getUserList(srv, "isAdmin=true")
	if len(r.Users) == 0 {
		fmt.Print("No users found.\n")
	} else {
		for _, u := range r.Users {
			s := strconv.FormatInt(itr, 10)

			report.Xls.File.SetCellValue(report.Xls.Sheet, "A"+s, u.PrimaryEmail)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "B"+s, u.Name.FullName)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "C"+s, u.IsEnrolledIn2Sv)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "D"+s, u.Suspended)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "E"+s, u.IsDelegatedAdmin)
			report.Xls.File.SetCellValue(report.Xls.Sheet, "F"+s, u.IsAdmin)

			itr++
		}
	}
}

func GetAsps(report *utils.Report) {
	srv := getService(report.Cl)
	//userkey := cfg.Section("users").Key("userkey").String()

	u := getUserList(srv, "")
	if len(u.Users) == 0 {
		fmt.Print("No users found.\n")
	} else {
		for _, user := range u.Users {
			r, err := srv.Asps.List(user.PrimaryEmail).Do()

			if err != nil {
				log.Fatalf("Unable to retrieve users in domain: %v", err)
			}

			if len(r.Items) == 0 {
				fmt.Printf("%s: No asps found.\n", user.PrimaryEmail)
			} else {
				for _, item := range r.Items {
					log.Printf("%s:%s\n", user.PrimaryEmail, item.Name)
				}
			}
		}
	}
}
