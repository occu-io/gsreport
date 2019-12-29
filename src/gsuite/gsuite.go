package gsuite

import (
	dir "github.com/occu-io/gsreport/src/gsuite/directory"
	rep "github.com/occu-io/gsreport/src/gsuite/report"
	"github.com/occu-io/gsreport/src/utils"
)

var Functions = map[string]func(*utils.Report){
	"GetUsersWithout2Fa":      dir.GetUsersWithout2Fa,
	"GetSuspendedUsers":       dir.GetSuspendedUsers,
	"GetAdmins":               dir.GetAdmins,
	"GetAsps":                 dir.GetAsps,
	"GetGroups":               dir.GetGroups,
	"GetGroupMembers":         dir.GetGroupMembers,
	"GetTokensThirdPartyApps": dir.GetTokensThirdPartyApps,
	"GetAuditLogins":          rep.GetAuditLogins,
}
