package service

import (
	"fmt"
	"github.com/zhongwq/Agenda-Golang/entity"
)

type User entity.User
type Meeting entity.Meeting
type Date entity.Date

/**
 * check if the username match password
 * @param userName the username want to login
 * @param password the password user enter
 * @return if success, true will be returned
 */
func userLogIn(userName string, password string) bool {
	userResult := entity.QueryUser(func(user *entity.User) bool {
		if user.GetName() == userName && user.GetPassword() == password {
			return true
		}
		return false
	})
	if len(userResult) == 0 {
		return false
	}
	entity.Signin(&userResult[0])
	if err := entity.Sync(); err != nil {
		fmt.Println("Login: error occur when sign in")
		return false
	}
	return true
}

/**
 * regist a user
 * @param userName new user's username
 * @param password new user's password
 * @param email new user's email
 * @param phone new user's phone
 * @return if success, true will be returned
 */
func userRegister(userName string, password string, email string, phone string) bool {
	user := entity.User{userName, password, email, phone}
	entity.CreateUser(&user)
	if err := entity.Sync(); err != nil {
		fmt.Println("Login: error occur when sign in")
		return false
	}
	return false
}

/**
 * delete a user
 * @param userName user's username
 * @param password user's password
 * @return if success, true will be returned
 */
func deleteUser(userName string, password string) bool {
	return false
}

/**
 * list all users from storage
 * @return a user list result
 */
func listAllUsers() []User {
	return nil
}


/**
 * create a meeting
 * @param userName the sponsor's userName
 * @param title the meeting's title
 * @param participator the meeting's participator
 * @param startData the meeting's start date
 * @param endData the meeting's end date
 * @return if success, true will be returned
 */
func createMeeting(userName string, title string, startDate Date, endDate Date, participator []string) bool {
	return false
}

/**
 * search a meeting by username and title
 * @param uesrName the sponsor's userName
 * @param title the meeting's title
 * @return a meeting list result
 */
func meetingQueryWithTitle(userName string, title string) []Meeting {
	return nil
}

/**
 * search a meeting by username, time interval
 * @param uesrName the sponsor's userName
 * @param startDate time interval's start date
 * @param endDate time interval's end date
 * @return a meeting list result
 */
func meetingQueryWithDate(userName string, startDate Date, endDate Date) []Meeting {
	return nil
}

/**
 * list all meetings the user take part in
 * @param userName user's username
 * @return a meeting list result
 */
func listAllMeetings(userName string) []Meeting {
	return nil
}

/**
 * list all meetings the user sponsor
 * @param userName user's username
 * @return a meeting list result
 */
func listAllSponsorMeetings(userName string) []Meeting{
	return nil
}

/**
 * list all meetings the user take part in and sponsor by other
 * @param userName user's username
 * @return a meeting list result
 */
func listAllParticipateMeetings(userName string) []Meeting{
	return nil
}

/**
 * delete a meeting by title and its sponsor
 * @param userName sponsor's username
 * @param title meeting's title
 * @return if success, true will be returned
 */
func deleteMeeting(userName string, title string) bool {
	return false
}

/**
 * delete all meetings by sponsor
 * @param userName sponsor's username
 * @return if success, true will be returned
 */
func deleteAllMeetings(userName string) bool {
	return false
}