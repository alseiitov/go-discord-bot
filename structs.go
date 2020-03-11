package main

type user struct {
	FirstName      string
	LastName       string
	InSchoolStatus int
	Last           string
	Avatar         string
}

type Users struct {
	Data struct {
		EventUser []struct {
			User struct {
				ID          int    `json:"id"`
				GithubLogin string `json:"githubLogin"`
				FirstName   string `json:"firstName"`
				LastName    string `json:"lastName"`
				Xp          struct {
					Aggregate struct {
						Sum struct {
							Amount interface{} `json:"amount"`
						} `json:"sum"`
					} `json:"aggregate"`
				} `json:"xp"`
				Audits struct {
					Aggregate struct {
						Count int `json:"count"`
					} `json:"aggregate"`
				} `json:"audits"`
			} `json:"user"`
		} `json:"event_user"`
	} `json:"data"`
}

type UserInfo struct {
	User struct {
		Data struct {
			User []struct {
				FirstName   string `json:"firstName"`
				GithubLogin string `json:"githubLogin"`
				ID          int    `json:"id"`
				LastName    string `json:"lastName"`
				Tel         string `json:"tel"`
			} `json:"user"`
		} `json:"data"`
	} `json:"user"`
	Attendance struct {
		Data []struct {
			Date      string `json:"date"`
			EmpID     int    `json:"emp_id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Status    int    `json:"status"`
			Time      string `json:"time"`
		} `json:"data"`
		Error     bool        `json:"error"`
		ErrorCode interface{} `json:"errorCode"`
		Message   interface{} `json:"message"`
	} `json:"attendance"`
	Progress struct {
		Data struct {
			User []struct {
				Audits struct {
					Aggregate struct {
						Count int `json:"count"`
					} `json:"aggregate"`
				} `json:"audits"`
				Progresses []struct {
					Object struct {
						Name string `json:"name"`
					} `json:"object"`
				} `json:"progresses"`
				Xp struct {
					Aggregate struct {
						Sum struct {
							Amount int `json:"amount"`
						} `json:"sum"`
					} `json:"aggregate"`
				} `json:"xp"`
			} `json:"user"`
		} `json:"data"`
	} `json:"progress"`
	Image struct {
		Data []struct {
			EmpID     int    `json:"emp_id"`
			Face      string `json:"face"`
			FirstName string `json:"first_name"`
			JobID     int    `json:"job_id"`
			LastName  string `json:"last_name"`
			LocID     int    `json:"loc_id"`
		} `json:"data"`
		Error     bool        `json:"error"`
		ErrorCode interface{} `json:"errorCode"`
		Message   interface{} `json:"message"`
	} `json:"image"`
}

type jwt struct {
	JwtToken string `json:"jwt_token"`
}
