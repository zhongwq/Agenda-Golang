package entity

import (
	"time"
)

type Meeting struct {
	Sponsor, Title     string
	Participators      []string
	StartDate, EndDate time.Time
}

/**
 * @brief constructor with argument
 */
func (meeting Meeting) init(sponsor string, participators []string, startTime time.Time, endTime time.Time, title string) {
	meeting.Sponsor = sponsor
	meeting.Title = title
	meeting.Participators = participators
	meeting.StartDate = startTime
	meeting.EndDate = endTime
}

/**
 * @brief copy constructor of left value
 */
func (meeting Meeting) Copy_Meeting(c_meeting Meeting) {
	c_meeting.Sponsor = meeting.Sponsor
	c_meeting.Title = meeting.Title
	c_meeting.Participators = meeting.Participators
	c_meeting.StartDate = meeting.StartDate
	c_meeting.EndDate = meeting.EndDate
}

/**
 *   @brief get the meeting's sponsor
 *   @return a string indicate sponsor
 */
func (meeting Meeting) GetSponsor() string {
	return meeting.Sponsor
}

/**
 * @brief set the sponsor of a meeting
 * @param  the new sponsor string
 */
func (meeting *Meeting) SetSponsor(sponsor string) {
	meeting.Sponsor = sponsor
}

/**
 * @brief  get the participators of a meeting
 * @return return a string array indicate participators
 */
func (meeting Meeting) GetParticipator() []string {
	return meeting.Participators
}

/**
 *   @brief set the new participators of a meeting
 *   @param the new participators array
 */
func (meeting *Meeting) SetParticipator(participators []string) {
	meeting.Participators = participators
}

/**
 * @brief get the startDate of a meeting
 * @return return a string indicate startDate
 */
func (meeting Meeting) GetStartDate() time.Time {
	return meeting.StartDate
}

/**
 * @brief  set the startDate of a meeting
 * @param  the new startdate of a meeting
 */
func (meeting *Meeting) SetStartDate(startTime time.Time) {
	meeting.StartDate = startTime
}

/**
 * @brief get the endDate of a meeting
 * @return a date indicate the endDate
 */
func (meeting Meeting) GetEndDate() time.Time {
	return meeting.EndDate
}

/**
 * @brief  set the endDate of a meeting
 * @param  the new enddate of a meeting
 */
func (meeting *Meeting) SetEndDate(endTime time.Time) {
	meeting.EndDate = endTime
}

/**
 * @brief get the title of a meeting
 * @return a date title the endDate
 */
func (meeting Meeting) GetTitle() string {
	return meeting.Title
}

/**
 * @brief  set the title of a meeting
 * @param  the new title of a meeting
 */
func (meeting *Meeting) SetTitle(title string) {
	meeting.Title = title
}

/**
 * @brief check if the user has been a participator
 * @param username the source username
 * @return if the user exit this meeting success
 */
func (meeting *Meeting) DeleteParticipator(username string) {
	size := len(meeting.Participators)
	for i := 0; i < size; i++ {
		if meeting.Participators[i] == username {
			meeting.Participators = append(meeting.Participators[:i], meeting.Participators[i+1:]...) //移除i
			break
		}
	}
}

/**
 * @brief check if the user if the sponsor or has been a participator
 * @param username the source username
 * @return if the user take part in this meeting success
 */
func (meeting *Meeting) AddParticipator(username string) bool {
	if meeting.Sponsor == username || meeting.IsParticipator(username) {
		return false
	}
	meeting.Participators = append(meeting.Participators, username)
	return true
}

/**
 * @brief check if the user take part in this meeting
 * @param username the source username
 * @return if the user take part in this meeting
 */
func (meeting Meeting) IsParticipator(username string) bool {
	for _, user := range meeting.Participators {
		if user == username {
			return true
		}
	}
	return false
}
