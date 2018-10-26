package entity

import (
	"os"
	"path/filepath"
)

var dirty bool
var userList []string
var meetingList []string
var currentUserName string

var userinfoPath string
var meetinginfoPath string
var stateinfoPath string // 当前用户数据(是否已登陆且记录登陆用户) 应该不用json？


/***
 * Initial variables
 */
func init() {
	gopath := os.Getenv("GOPATH")
	userinfoPath =  filepath.Join(gopath, "/src/github.com/zhongwq/Agenda-Golang/userdata.json")
	meetinginfoPath = filepath.Join(gopath, "/src/github.com/zhongwq/Agenda-Golang/meetingdata.json")
	stateinfoPath = filepath.Join(gopath, "/src/github.com/zhongwq/Agenda-Golang/stateinfo")
}

/**
 * read file content into memory
 * @return if success, true will be returned
 */
func ReadFromFile() bool {
	return false
}

/**
 * write file content from memory
 * @return if success, true will be returned
 */
func WriteToFile() bool {
	return false
}

/**
 * create a user
 * @param a user object
 */
func CreateUser(u* User) (bool,string) {
	return false, ""
}

/**
 * query users
 * @param a lambda function as the filter
 * @return a list of fitted users
 */
func QueryUser(filter func (user* User) bool) []User {
	return nil
}

/**
 * update users
 * @param a lambda function as the filter
 * @param a lambda function as the method to update the user
 * @return the number of updated users
 */
func UpdateUser(filter func (user* User) bool, switcher func (user* User)) int {
	return 0
}

/**
 * delete users
 * @param a lambda function as the filter
 * @return the number of deleted users
 */
func DeleteUser(filter func (user* User) bool) int {
	return 0
}

/***
 * user log in
 * @param user that login
 */
func Signin(user* User){

}

/**
 * user log out
 * @return whether have error when sync
 */
func Logout() {

}

/***
 * user log in
 * @param user that login
 */
func CreateMeeting(meeting* Meeting) (bool,string) {
	return false, ""
}


/***
 * query meetings
 * @param a lambda function as the filter
 * @return a list of fitted meetings
 */
func QueryMeeting(filter func (meeting* Meeting) bool) []Meeting {
	return nil
}

/**
 * update meetings
 * @param a lambda function as the filter
 * @param a lambda function as the method to update the meeting
 * @return the number of updated meetings
 */
func UpdateMeeting(filter func (meeting* Meeting) bool, switcher func (meeting* Meeting)) int {
	return 0
}

/**
 * delete meetings
 * @param a lambda function as the filter
 * @return the number of deleted meetings
 */
func DeleteMeeting(filter func (meeting* Meeting) bool) int {
	return 0
}

/**
 * sync with the file
 */
func Sync() error {
	// 同步三个文件
	return nil
}