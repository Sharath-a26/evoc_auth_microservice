package routes

const (
	BASE = "/api"
)

const (
	TEST     = BASE + "/test"
	REGISTER = BASE + "/register"
	VERIFY   = REGISTER + "/verify"
	LOGIN    = BASE + "/login"

	CREATETEAM = BASE + "/team/create"
	ADDMEMBERS = BASE + "/team/addMembers"
	DELETEMEMBERS = BASE + "/team/deleteMembers"
	GETTEAMS = BASE + "/team/getTeams"
	GETMEMBERS = BASE + "/team/getMembers"
)
