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
func readFromFile() bool {
	return false
}

/**
 * write file content from memory
 * @return if success, true will be returned
 */
func writeToFile() bool {
	return false
}

/**
 * create a user
 * @param a user object
 */
func createUser(u* User) {

}

/**
 * query users
 * @param a lambda function as the filter
 * @return a list of fitted users
 */
func queryUser(filter func (user* User) bool) []string {
	return nil
}

/**
 * update users
 * @param a lambda function as the filter
 * @param a lambda function as the method to update the user
 * @return the number of updated users
 */
func updateUser(filter func (user* User) bool, switcher func (user* User)) int{
	return 0
}

/**
 * delete users
 * @param a lambda function as the filter
 * @return the number of deleted users
 */
func deleteUser(filter func (user* User) bool) int {
	return 0
}

/***
 * user log in
 * @param user that login
 */
func Signin(user* User) {

}


/**
 * user log out
 * @return whether have error when sync
 */
func Logout() error {
	return nil
}



/**
 * create a meeting
 * @param a meeting object
 */
func createMeeting(meeting* Meeting) {

}

/**
 * query meetings
 * @param a lambda function as the filter
 * @return a list of fitted meetings
 */
func queryMeeting(filter func (meeting* Meeting) bool) []Meeting {
	return nil
}

/**
 * update meetings
 * @param a lambda function as the filter
 * @param a lambda function as the method to update the meeting
 * @return the number of updated meetings
 */
func updateMeeting(filter func (meeting* Meeting) bool, switcher func (meeting* Meeting)) int {
	return 0
}

/**
 * delete meetings
 * @param a lambda function as the filter
 * @return the number of deleted meetings
 */
func deleteMeeting(filter func (meeting* Meeting) bool) int {
	return 0
}

/**
 * sync with the file
 */
func sync() bool {
	// 同步三个文件
	return false
}