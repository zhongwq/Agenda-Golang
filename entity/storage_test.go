package entity

import (
	"testing"
	"time"
)

func TestCreateUser(test *testing.T) {
	var user = User{"lim", "lim", "lim", "lim"}
	judge, str := CreateUser(&user)
	if !judge {
		test.Error(str)
	}
}

func TestQueryUser(test *testing.T) {
	var l []User
	l = QueryUser(func(a *User) bool {
		if a.Name == "lim" {
			return true
		}
		return false
	})
	if len(l) == 0 {
		test.Error("error")
	}
}

func TestUpdateUser(test *testing.T) {
	count := UpdateUser(func(a *User) bool {
		if a.Name == "lim" {
			return true
		}
		return false
	}, func(b *User) {
		b.Password = "1234"
	})
	if count != 1 {
		test.Error("error")
	}
}

func TestSignin(test *testing.T) {
	user := User{"lim", "1234", "lim", "lim"}
	Signin(&user)
	tmp, _ := GetCurrentUser()
	if tmp.Name != "lim" {
		test.Error("error")
	}
}

func TestLogout(test *testing.T) {
	if !Logout() {
		test.Error("error")
	}
}

func TestDeleteUser(test *testing.T) {
	count := DeleteUser(func(a *User) bool {
		if a.Name == "lim" {
			return true
		}
		return false
	})
	if count != 1 {
		test.Error("error")
	}
}

func TestCreateMeeting(test *testing.T) {
	participators := []string{"wilson", "heng"}
	startDate, _ := time.Parse("2006-01-02 15:04:05", "2018-11-02 12:00:00")
	endDate, _ := time.Parse("2006-01-02 15:04:05", "2018-11-02 13:00:00")
	meeting := Meeting{"lim", "test1", participators, startDate, endDate}
	judge, _ := CreateMeeting(&meeting)
	if !judge {
		test.Error("error")
	}
}

func TestQueryMeeting(test *testing.T) {
	meeting := QueryMeeting(func(a *Meeting) bool {
		if a.Sponsor == "lim" {
			return true
		}
		return false
	})
	if len(meeting) == 0 {
		test.Error("error")
	}
}

func TestUpdateMeeting(test *testing.T) {
	count := UpdateMeeting(func(a *Meeting) bool {
		if a.Sponsor == "lim" {
			return true
		}
		return false
	}, func(b *Meeting) {
		b.Sponsor = "lim1"
	})
	if count != 1 {
		test.Error("error")
	}
}

func TestDeleteMeeting(test *testing.T) {
	count := DeleteMeeting(func(a *Meeting) bool {
		if a.Sponsor == "lim1" {
			return true
		}
		return false
	})
	if count != 1 {
		test.Error("error")
	}
}
