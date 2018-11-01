package entity

import (
	"Agenda-Golang/logInit"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var dirty bool
var userList []User
var meetingList []Meeting
var currentUserName string

var userinfoPath string
var meetinginfoPath string
var stateinfoPath string // 当前用户数据(是否已登陆且记录登陆用户) 应该不用json？

var errorlog *log.Logger

/***
 * Initial variables
 */
func init() {
	errorlog = logInit.Error
	gopath := os.Getenv("GOPATH")
	userinfoPath = filepath.Join(gopath, "/src/Agenda-Golang/data/userdata.json")
	meetinginfoPath = filepath.Join(gopath, "/src/Agenda-Golang/data/meetingdata.json")
	stateinfoPath = filepath.Join(gopath, "/src/Agenda-Golang/data/stateinfo.txt")
	ReadFromFile()
}

/**
 * read file content into memory
 * @return if success, true will be returned
 */
func ReadFromFile() bool {
	// Read user data
	data, err := ioutil.ReadFile(userinfoPath)
	var str = string(data)
	if len(str) != 0 {
		err := json.Unmarshal([]byte(str), &userList)
		if err != nil {
			errorlog.Println("Read user data failed")
		}
	}

	// Read meetingdata
	data, err = ioutil.ReadFile(meetinginfoPath)
	str = string(data)
	if len(str) != 0 {
		err := json.Unmarshal([]byte(str), &meetingList)
		if err != nil {
			errorlog.Println("Read meeting data failed")
		}
	}

	// Read state
	data, err = ioutil.ReadFile(stateinfoPath)
	if err != nil {
		errorlog.Println("Read state data failed")
	}
	str = string(data)
	if len(str) != 0 {
		currentUserName = str
	}
	return true
}

/**
 * write file content from memory
 * @return if success, true will be returned
 */
func WriteToFile() bool {
	// Write user data
	data, _ := json.Marshal(userList)
	if len(userList) != 0 {
		err := ioutil.WriteFile(userinfoPath, data, 0644)
		if err != nil {
			errorlog.Println("Write user data failed")
			return false
		}
	}

	// Write meeting data
	data, _ = json.Marshal(meetingList)
	if len(meetingList) != 0 {
		err := ioutil.WriteFile(meetinginfoPath, data, 0644)
		if err != nil {
			errorlog.Println("Write meeting data failed")
			return false
		}
	}
	return true
}

/**
 * create a user
 * @param a user object
 */
func CreateUser(u *User) (bool, string) {
	var l []User
	l = QueryUser(func(a *User) bool {
		if a.Name == u.Name {
			return true
		}
		return false
	})
	if len(l) == 0 {
		userList = append(userList, *u)
		dirty = true
		Sync()
		return true, "Create successfully"
	}
	return false, "User exist"
}

/**
 * query users
 * @param a lambda function as the filter
 * @return a list of fitted users
 */
func QueryUser(filter func(user *User) bool) []User {
	var filtedUser []User
	for i := 0; i < len(userList); i++ {
		if filter(&userList[i]) {
			filtedUser = append(filtedUser, userList[i])
		}
	}
	return filtedUser
}

/**
 * update users
 * @param a lambda function as the filter
 * @param a lambda function as the method to update the user
 * @return the number of updated users
 */
func UpdateUser(filter func(user *User) bool, switcher func(user *User)) int {
	var updatedCount = 0
	for i := 0; i < len(userList); i++ {
		if filter(&userList[i]) {
			switcher(&userList[i])
			updatedCount++
			dirty = true
			Sync()
		}
	}
	return updatedCount
}

/**
 * delete users
 * @param a lambda function as the filter
 * @return the number of deleted users
 */
func DeleteUser(filter func(user *User) bool) int {
	var deletedCount = 0
	for i := 0; i < len(userList); i++ {
		if filter(&userList[i]) {
			userList = append(userList[:i], userList[i+1:]...)
			i--
			deletedCount++
			dirty = true
			Sync()
		}
	}
	return deletedCount
}

/**
 * GetCurrentUser : get current user
 * @return the current user
 * @return error if current user does not exist
 */
func GetCurrentUser() (User, bool) {
	var l []User
	l = QueryUser(func(a *User) bool {
		if a.Name == currentUserName {
			return true
		}
		return false
	})
	if len(l) == 0 {
		return User{}, false
	} else {
		return l[0], true
	}
}

/***
 * user log in
 * @param user that login
 */
func Signin(user *User) {
	data := []byte(user.Name)
	err := ioutil.WriteFile(stateinfoPath, data, 0644)
	if err != nil {
		errorlog.Println("Write state data failed")
	}
	currentUserName = user.Name
}

/**
 * user log out
 * @return whether have error when sync
 */
func Logout() bool {
	if Sync() != nil {
		return false
	}
	data := []byte("")
	err := ioutil.WriteFile(stateinfoPath, data, 0644)
	if err != nil {
		errorlog.Println("Write state data failed")
	}
	return true
}

/***
 * user log in
 * @param user that login
 */
func CreateMeeting(meeting *Meeting) (bool, string) {
	var l []Meeting
	l = QueryMeeting(func(a *Meeting) bool {
		if a.Title == meeting.Title {
			return true
		}
		return false
	})
	if len(l) == 0 {
		meetingList = append(meetingList, *meeting)
		dirty = true
		Sync()
		return true, "Create successfully"
	}
	return false, "Meeting exist"
}

/***
 * query meetings
 * @param a lambda function as the filter
 * @return a list of fitted meetings
 */
func QueryMeeting(filter func(meeting *Meeting) bool) []Meeting {
	var filtedMeeting []Meeting
	for i := 0; i < len(meetingList); i++ {
		if filter(&meetingList[i]) {
			filtedMeeting = append(filtedMeeting, meetingList[i])
		}
	}
	return filtedMeeting
}

/**
 * update meetings
 * @param a lambda function as the filter
 * @param a lambda function as the method to update the meeting
 * @return the number of updated meetings
 */
func UpdateMeeting(filter func(meeting *Meeting) bool, switcher func(meeting *Meeting)) int {
	var updatedCount = 0
	for i := 0; i < len(meetingList); i++ {
		if filter(&meetingList[i]) {
			switcher(&meetingList[i])
			updatedCount++
			dirty = true
			Sync()
		}
	}
	return updatedCount
}

/**
 * delete meetings
 * @param a lambda function as the filter
 * @return the number of deleted meetings
 */
func DeleteMeeting(filter func(meeting *Meeting) bool) int {
	var deletedCount = 0
	for i := 0; i < len(meetingList); i++ {
		if filter(&meetingList[i]) {
			meetingList = append(meetingList[:i], meetingList[i+1:]...)
			i--
			deletedCount++
			dirty = true
			Sync()
		}
	}
	return deletedCount
}

/**
 * sync with the file
 */
func Sync() error {
	// 同步三个文件
	if WriteToFile() {
		return nil
	}
	return errors.New("error")
}
