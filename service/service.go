package service

import (
	"fmt"
	"github.com/zhongwq/Agenda-Golang/entity"
	"time"
)

type User entity.User
type Meeting entity.Meeting

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
		fmt.Println("Error: incorrect Username or Passwd")
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
	if flag, msg := entity.CreateUser(&user); flag == false {
		fmt.Println(msg)
		return false
	}
	return true
}

/**
 * Logout agenda
 * @return if success, true will be returned
 */
func UserLogout() bool {
	return false
}


/**
 * delete a user
 * @param userName user's username
 * @param password user's password
 * @return if success, true will be returned
 */
func deleteUser(userName string, password string) bool {
	userResult := entity.DeleteUser(func(user *entity.User) bool {
		if user.GetName() == userName && user.GetPassword() == password {
			return true
		}
		return false
	})
	if userResult == 0 {
		fmt.Println("Error: incorrect username or password")
		return false
	}

	entity.UpdateMeeting(
		func(meeting *entity.Meeting) bool {
			return meeting.IsParticipator(userName)
		},
		func(meeting *entity.Meeting) {
			meeting.DeleteParticipator(userName)
		})

	entity.DeleteMeeting(func(meeting *entity.Meeting) bool {
		return meeting.GetSponsor() == userName
	})
	return UserLogout()
}
/**
 * list all users from storage
 * @return a user list result
 */
func listAllUsers() []entity.User {
	return entity.QueryUser(func (user* entity.User) bool {
		return true
	})
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
func createMeeting(userName string, title string, startDate time.Time, endDate time.Time, participator []string) bool {
	if startDate.After(endDate) {
		fmt.Println("Create Meeting: startDate should before endDate")
		return false
	}

	for i := 0; i < len(participator); i++ {
		if (userName == participator[i]) {
			fmt.Println("Create Meeting: participator can;t be sponsor!")
			return false
		}
		userResult := entity.QueryUser(func(user *entity.User) bool {
			return user.GetName() == participator[i]
		})
		if (len(userResult) == 0) {
			fmt.Println("Create Meeting: Can't find user named " + participator[i] + "!")
			return false
		}

		meetingResult := entity.QueryMeeting(func(meeting *entity.Meeting) bool {
			if meeting.GetSponsor() == participator[i] || meeting.IsParticipator(participator[i]) {
				if meeting.GetStartDate().Before(startDate) && meeting.GetEndDate().After(startDate) {
					return true
				}
				if meeting.GetStartDate().Before(endDate) && meeting.GetEndDate().After(endDate) {
					return true
				}
				if meeting.GetStartDate().After(startDate) && meeting.GetEndDate().Before(endDate) {
					return true
				}
			}
			return false
		})
		if len(meetingResult) > 0 {
			fmt.Println("Create Meeting: " + participator[i] + "'s time is conflict!")
			return false
		}

		for j := i+1; j < len(participator); j++ {
			if participator[j] == participator[i] {
				fmt.Println("Create Meeting: duplicate participator named" + participator[i])
				return false
			}
		}
	}

	meetingResult := entity.QueryMeeting(func(meeting *entity.Meeting) bool {
		if meeting.GetSponsor() == userName || meeting.IsParticipator(userName) {
			if meeting.GetStartDate().Before(startDate) && meeting.GetEndDate().After(startDate) {
				return true
			}
			if meeting.GetStartDate().Before(endDate) && meeting.GetEndDate().After(endDate) {
				return true
			}
			if meeting.GetStartDate().After(startDate) && meeting.GetEndDate().Before(endDate) {
				return true
			}
		}
		return false
	})

	if len(meetingResult) > 0 {
		fmt.Println("Create Meeting: sponsor's time is conflict!")
		return false
	}

	if flag, msg := entity.CreateMeeting(&entity.Meeting{userName, title, participator, startDate, endDate}); flag ==false {
		fmt.Println(msg)
		return false
	}

	return true
}

/**
 * search a meeting by username and title
 * @param uesrName the sponsor's userName
 * @param title the meeting's title
 * @return a meeting list result
 */
func meetingQueryWithTitle(userName string, title string) []entity.Meeting {
	return entity.QueryMeeting(func(meeting *entity.Meeting) bool {
		if meeting.GetSponsor() == userName && meeting.GetTitle() == title {
			return true
		}
		return false
	})
}

/**
 * search a meeting by username, time interval
 * @param uesrName the sponsor's userName
 * @param startDate time interval's start date
 * @param endDate time interval's end date
 * @return a meeting list result
 */
func meetingQueryWithDate(userName string, startDate time.Time, endDate time.Time) []entity.Meeting {
	return entity.QueryMeeting(func(meeting *entity.Meeting) bool {
		if meeting.GetSponsor() == userName || meeting.IsParticipator(userName) {
			if meeting.GetStartDate().Before(startDate) && meeting.GetEndDate().After(startDate) {
				return true
			}
			if meeting.GetStartDate().Before(endDate) && meeting.GetEndDate().After(endDate) {
				return true
			}
			if meeting.GetStartDate().After(startDate) && meeting.GetEndDate().Before(endDate) {
				return true
			}
		}
		return false
	})
}

/**
 * list all meetings the user take part in
 * @param userName user's username
 * @return a meeting list result
 */
func listAllMeetings(userName string) []entity.Meeting {
	return entity.QueryMeeting(func(meeting *entity.Meeting) bool {
		return true
	})
}

/**
 * list all meetings the user sponsor
 * @param userName user's username
 * @return a meeting list result
 */
func listAllSponsorMeetings(userName string) []entity.Meeting{
	return entity.QueryMeeting(func(meeting *entity.Meeting) bool {
		if meeting.GetSponsor() == userName {
			return true
		}
		return false
	})
}

/**
 * list all meetings the user take part in and sponsor by other
 * @param userName user's username
 * @return a meeting list result
 */
func listAllParticipateMeetings(userName string) []entity.Meeting{
	return entity.QueryMeeting(func(meeting *entity.Meeting) bool {
		if meeting.IsParticipator(userName) {
			return true
		}
		return false
	})
}

/**
 * delete a meeting by title and its sponsor
 * @param userName sponsor's username
 * @param title meeting's title
 * @return if success, true will be returned
 */
func deleteMeeting(userName string, title string) bool {
	result := entity.DeleteMeeting(func(meeting *entity.Meeting) bool {
		if meeting.GetTitle() == title && meeting.GetSponsor() == userName {
			return true
		}
		return false
	})
	return result > 0
}

/**
 * delete all meetings by sponsor
 * @param userName sponsor's username
 * @return if success, true will be returned
 */
func deleteAllMeetings(userName string) bool {
	return entity.DeleteMeeting(func(meeting *entity.Meeting) bool {
		if  meeting.GetSponsor() == userName {
			return true
		}
		return false
	}) > 0
}