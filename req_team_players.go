package footballdata

type TeamPlayersRequest struct{ request }

// Do executes the request.
func (r TeamPlayersRequest) Do() (s PlayerList, err error) {
	d, _, err := r.doJson("GET")
	if err != nil {
		return
	}

	err = d.Decode(&s)
	return
}

// PlayersOfTeam prepares a request to fetch a team's players.
func (c *Client) PlayersOfTeam(id uint64) TeamPlayersRequest {
	return TeamPlayersRequest{c.req("teams/%d/players", id)}
}
