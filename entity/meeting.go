package entity

import (
	"time"
)

type Date time.Time //定义Date类型为time.Time

type Meeting struct {
	sponsor, title string
	participators []string
	startDate, endDate Date
}

/**
 * @brief constructor with argument
 */
func (meeting Meeting) init(sponsor string, participators []string, startTime Date, endTime Date, title string) {
	meeting.sponsor = sponsor
	meeting.title = title
	meeting.participators = participators
	meeting.startDate = startTime
	meeting.endDate = endTime
}

/**
 * @brief copy constructor of left value
 */
func (meeting Meeting) copy_User(c_meeting Meeting) {
	c_meeting.sponsor = meeting.sponsor
	c_meeting.title = meeting.title
	c_meeting.participators = meeting.participators
	c_meeting.startDate = meeting.startDate
	c_meeting.endDate = meeting.endDate
}

/**
 *   @brief get the meeting's sponsor
 *   @return a string indicate sponsor
 */
func (meeting Meeting) getSponsor() string {
	return meeting.sponsor
}

/**
 * @brief set the sponsor of a meeting
 * @param  the new sponsor string
 */
func (meeting* Meeting) setSponsor(sponsor string) {
	meeting.sponsor = sponsor
}

/**
 * @brief  get the participators of a meeting
 * @return return a string array indicate participators
 */
func (meeting Meeting) getParticipator() []string {
	return meeting.participators
}

/**
 *   @brief set the new participators of a meeting
 *   @param the new participators array
 */
func (meeting* Meeting) setParticipator(participators []string) {
	meeting.participators = participators
}

/**
 * @brief get the startDate of a meeting
 * @return return a string indicate startDate
 */
func (meeting Meeting) getStartDate() Date {
	return meeting.startDate
}

/**
 * @brief  set the startDate of a meeting
 * @param  the new startdate of a meeting
 */
func (meeting* Meeting) setStartDate(startTime Date) {
	meeting.startDate = startTime
}

/**
 * @brief get the endDate of a meeting
 * @return a date indicate the endDate
 */
func (meeting Meeting) getEndDate() Date {
	return meeting.endDate
}

/**
 * @brief  set the endDate of a meeting
 * @param  the new enddate of a meeting
 */
func (meeting* Meeting) setEndDate(endTime Date) {
	meeting.endDate = endTime
}

/**
 * @brief get the title of a meeting
 * @return a date title the endDate
 */
func (meeting Meeting) getTitle() string {
	return meeting.title
}

/**
 * @brief  set the title of a meeting
 * @param  the new title of a meeting
 */
func (meeting* Meeting) setTitle(title string) {
	meeting.title = title
}

/**
 * @brief check if the user has been a participator
 * @param username the source username
 * @return if the user exit this meeting success
 */
func (meeting* Meeting) deleteParticipator(username string) {
	size := len(meeting.participators)
	for i := 0; i< size; i++ {
		if meeting.participators[i] == username {
			meeting.participators = append(meeting.participators[:i], meeting.participators[i+1:]...)//移除i
			break
		}
	}
}

/**
 * @brief check if the user if the sponsor or has been a participator
 * @param username the source username
 * @return if the user take part in this meeting success
 */
func (meeting* Meeting) addParticipator(username string) bool {
	if meeting.sponsor == username || meeting.isParticipator(username) {
		return false
	}
	meeting.participators = append(meeting.participators, username)
	return true
}

/**
 * @brief check if the user take part in this meeting
 * @param username the source username
 * @return if the user take part in this meeting
 */
func (meeting Meeting) isParticipator(username string) bool{
	for _, user := range meeting.participators {
		if (user == username) {
			return true
		}
	}
	return false
}