package team

import (
	"context"
	"encoding/json"
	"fmt"
	"evolve/db/connection"
	"evolve/util"
	
)
type CreateTeamReq struct {
	TeamName string `json:"teamName"`
	TeamDesc    string `json:"teamDesc"`
}

func CreateTeamReqFromJson(data map[string]any) (*CreateTeamReq, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var createTeamReq *CreateTeamReq
	if err := json.Unmarshal(jsonData, &createTeamReq); err != nil {
		return nil, err
	}

	return createTeamReq, nil
}


func (c *CreateTeamReq) CreateTeam(ctx context.Context, user map[string]string) error {
	
	logger := util.SharedLogger

	db, err := connection.PoolConn(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("CreateTeam: failed to get pool connection: %v", err), err)
		return fmt.Errorf("something went wrong")
	}

	// userInfo, err := dbutil.UserById(ctx, user["id"], db)

	if err != nil {
		logger.Error(fmt.Sprintf("CreateTeam: Invalid UserID: %v", err), err)
		return fmt.Errorf("Invalid User")
	}	

	//inserting team info in team table
	var teamID string
	err = db.QueryRow(ctx, "INSERT INTO team (teamName, teamDesc, createdBy) VALUES ($1, $2, $3) RETURNING teamID", c.TeamName, c.TeamDesc, user["id"]).Scan(&teamID)
	
	if err != nil {
		logger.Error(fmt.Sprintf("Createteam: failed to create Team: %v", err), err)
		return fmt.Errorf("something went wrong")
	}

	//inserting the admin into teamMembers table
	_, err = db.Exec(ctx, "INSERT INTO teamMembers (memberId, teamID, role) VALUES ($1, $2, $3)", user["id"], teamID, "Admin")

	if err != nil {
		logger.Error(fmt.Sprintf("Createteam: failed to Insert into teamMembers Table: %v", err), err)
		return fmt.Errorf("something went wrong")
	}
	return nil

}