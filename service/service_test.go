package service

import (
	"testing"
	"time"
)

func TestUserRegister(test *testing.T) {
	if !UserRegister("1", "1", "1", "1") {
		test.Error("error")
	}
	if !UserRegister("2", "2", "2", "2") {
		test.Error("error")
	}
	if !UserRegister("3", "3", "3", "3") {
		test.Error("error")
	}
	if UserRegister("2", "2", "2", "2") {
		test.Error("error")
	}
}

func TestUserLogIn(test *testing.T) {
	if !UserLogIn("1", "1") {
		test.Error("error")
	}
}

func TestGetCurrentUser(test *testing.T) {
	_, judge := GetCurrentUser()
	if !judge {
		test.Error("error")
	}
}

func TestLogOut(test *testing.T) {
	if !UserLogout() {
		test.Error("error")
	}
}

func TestDeleteUser(test *testing.T) {
	UserLogIn("3", "3")
	if !DeleteUser("3", "3") {
		test.Error("error")
	}
}

func TestListAllUser(test *testing.T) {
	users := ListAllUsers()
	if len(users) != 2 {
		test.Error("error")
	}
}

func TestCreateMeeting(test *testing.T) {
	sponsor := "1"
	title := "test"
	participators := []string{"2"}
	startDate, _ := time.Parse("2006-01-02 15:04:05", "2018-11-02 12:00:00")
	endDate, _ := time.Parse("2006-01-02 15:04:05", "2018-11-02 13:00:00")
	if !CreateMeeting(sponsor, title, startDate, endDate, participators) {
		test.Error("error")
	}
}

func TestAddMeetingParticipator(test *testing.T) {
	UserRegister("3", "3", "3", "3")
	if !AddMeetingParticipator("1", "test", "3") {
		test.Error("error")
	}
}

func TestDeleteMeetingParticipator(test *testing.T) {
	if !DeleteMeetingParticipator("1", "test", "3") {
		test.Error("error")
	}
}

func TestQuitMeeting(test *testing.T) {
	UserLogIn("2", "2")
	AddMeetingParticipator("1", "test", "3")
	if !QuitMeeting("2", "test") {
		test.Error("error")
	}
}

func TestMeetingQueryWithTitle(test *testing.T) {
	UserLogIn("1", "1")
	meeting := MeetingQueryWithTitle("1", "test")
	if len(meeting) != 1 {
		test.Error("error")
	}
}

func TestMeetingQueryWithDate(test *testing.T) {
	startDate, _ := time.Parse("2006-01-02 15:04:05", "2018-11-02 12:10:00")
	endDate, _ := time.Parse("2006-01-02 15:04:05", "2018-11-02 12:30:00")
	meeting := MeetingQueryWithDate("1", startDate, endDate)
	if len(meeting) != 1 {
		test.Error("error")
	}
}

func TestListAllMeetings(test *testing.T) {
	meeting := ListAllMeetings("1")
	if len(meeting) != 1 {
		test.Error("error")
	}
}

func TestListAllSponsorMeetings(test *testing.T) {
	meeting := ListAllSponsorMeetings("1")
	if len(meeting) != 1 {
		test.Error("error")
	}
}

func TestListAllParticipateMeetings(test *testing.T) {
	meeting := ListAllParticipateMeetings("1")
	if len(meeting) != 0 {
		test.Error("error")
	}
}

func TestDeleteMeeting(test *testing.T) {
	if !DeleteMeeting("1", "test") {
		test.Error("error")
	}
}

func TestDeleteAllMeetings(test *testing.T) {
	sponsor := "1"
	title := "test"
	participators := []string{"2"}
	startDate, _ := time.Parse("2006-01-02 15:04:05", "2018-11-02 12:00:00")
	endDate, _ := time.Parse("2006-01-02 15:04:05", "2018-11-02 13:00:00")
	CreateMeeting(sponsor, title, startDate, endDate, participators)
	sponsor2 := "1"
	title2 := "test2"
	participators2 := []string{"2"}
	startDate2, _ := time.Parse("2006-01-02 15:04:05", "2018-11-03 12:00:00")
	endDate2, _ := time.Parse("2006-01-02 15:04:05", "2018-11-03 13:00:00")
	CreateMeeting(sponsor2, title2, startDate2, endDate2, participators2)
	if !DeleteAllMeetings("1") {
		test.Error("error")
	}
}
