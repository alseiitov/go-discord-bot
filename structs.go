package main

import "time"

type user struct {
	Login          string
	FirstName      string
	LastName       string
	InSchoolStatus int
	Last           string
	Avatar         string
	DoneProjects   []string
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
	Exams struct {
		Data struct {
			Object []struct {
				Events []struct {
					CreatedAt time.Time `json:"createdAt"`
					EndAt     time.Time `json:"endAt"`
					ID        int       `json:"id"`
				} `json:"events"`
				ID int `json:"id"`
			} `json:"object"`
		} `json:"data"`
	} `json:"exams"`
	ExamRecords struct {
		Data struct {
			Exam156 []struct {
				Amount int `json:"amount"`
				Attrs  struct {
					EventID  int `json:"eventId"`
					ObjectID int `json:"objectId"`
				} `json:"attrs"`
				CreatedAt time.Time `json:"createdAt"`
			} `json:"exam_156"`
			Exam162 []struct {
				Amount int `json:"amount"`
				Attrs  struct {
					EventID  int `json:"eventId"`
					ObjectID int `json:"objectId"`
				} `json:"attrs"`
				CreatedAt time.Time `json:"createdAt"`
			} `json:"exam_162"`
			Exam165 []struct {
				Amount int `json:"amount"`
				Attrs  struct {
					EventID  int `json:"eventId"`
					ObjectID int `json:"objectId"`
				} `json:"attrs"`
				CreatedAt time.Time `json:"createdAt"`
			} `json:"exam_165"`
			Exam168 []struct {
				Amount int `json:"amount"`
				Attrs  struct {
					EventID  int `json:"eventId"`
					ObjectID int `json:"objectId"`
				} `json:"attrs"`
				CreatedAt time.Time `json:"createdAt"`
			} `json:"exam_168"`
		} `json:"data"`
	} `json:"examRecords"`
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
