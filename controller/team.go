package controller

import (
	"evolve/util"
	"evolve/util/auth"
	"net/http" 
	"fmt"
	"evolve/modules/team"
)

func CreateTeam(res http.ResponseWriter, req *http.Request) {

	logger := util.SharedLogger
	logger.InfoCtx(req, "CreateTeam API called.")

	token, err := req.Cookie("t")

	if err != nil {
		util.JSONResponse(res, http.StatusUnauthorized, "You got to try way better than that.", nil)
		return
	}

	payLoad, err := auth.ValidateToken(token.Value)

	if err != nil {
		util.JSONResponse(res, http.StatusUnauthorized, "Session Expired.", nil)
		return
	}
	fmt.Println(payLoad)
	if payLoad["purpose"] != "login" {
		util.JSONResponse(res, http.StatusUnauthorized, "Good try.", nil)
		return
	}

	//should change it to data
	data, err := util.Body(req)
	if err != nil {
		util.JSONResponse(res, http.StatusBadRequest, err.Error(), nil)
		return
	}
	
	createTeamReq, err := team.CreateTeamReqFromJson(data)
	if err != nil {
		util.JSONResponse(res, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = createTeamReq.CreateTeam(req.Context(), payLoad)

	if err != nil {
		util.JSONResponse(res, http.StatusBadRequest, err.Error(), nil)
		return
	}

	util.JSONResponse(res, http.StatusOK, "Team Creation Successful.", nil)


}