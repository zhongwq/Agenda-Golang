package entity

type User struct {
	Name, Password, Email, Phone string
}

/**
 * Constructor of User
 */
func (user User) init(name string, password string, email string, phone string) {
	user.Name = name
	user.Password = password
	user.Email = email
	user.Phone = phone
}

/**
 * @brief copy constructor
 */
func (user User) CopyUser(c_user User) {
	c_user.Name = user.Name
	c_user.Password = user.Password
	user.Email = user.Email
	user.Phone = user.Phone
}

/**
* @brief get the name of the user
* @return   return a string indicate the name of the user
 */
func (user User) GetName() string {
	return user.Name
}

/**
* @brief set the name of the user
* @param   a string indicate the new name of the user
 */
func (user *User) SetName(name string) {
	user.Name = name
}

/**
* @brief get the password of the user
* @return   return a string indicate the password of the user
 */
func (user User) GetPassword() string {
	return user.Password
}

/**
* @brief set the password of the user
* @param   a string indicate the new password of the user
 */
func (user *User) SetPassword(password string) {
	user.Password = password
}

/**
* @brief get the email of the user
* @return   return a string indicate the email of the user
 */
func (user User) GetEmail() string {
	return user.Email
}

/**
* @brief set the email of the user
* @param   a string indicate the new email of the user
 */
func (user *User) SetEmail(email string) {
	user.Email = email
}

/**
* @brief get the phone of the user
* @return   return a string indicate the phone of the user
 */
func (user User) GetPhone() string {
	return user.Phone
}

/**
* @brief set the phone of the user
* @param   a string indicate the new phone of the user
 */
func (user *User) SetPhone(phone string) {
	user.Phone = phone
}
