package entity

type User struct {
	name, password, email, phone string
}


/**
 * Constructor of User
 */
func (user User) init (name string, password string, email string, phone string)  {
	user.name = name
	user.password = password
	user.email = email
	user.phone = phone
}



/**
  * @brief copy constructor
  */
func (user User) copyUser(c_user User) {
	c_user.name = user.name
	c_user.password = user.password
	user.email = user.email
	user.phone = user.phone
}

/**
* @brief get the name of the user
* @return   return a string indicate the name of the user
*/
func (user User) getName() string {
	return user.name
}

/**
* @brief set the name of the user
* @param   a string indicate the new name of the user
*/
func (user* User) setName(name string) {
	user.name = name
}

/**
* @brief get the password of the user
* @return   return a string indicate the password of the user
*/
func (user User) getPassword() string {
	return user.password
}

/**
* @brief set the password of the user
* @param   a string indicate the new password of the user
*/
func (user* User) setPassword(password string) {
	user.password = password
}

/**
* @brief get the email of the user
* @return   return a string indicate the email of the user
*/
func (user User) getEmail() string {
	return user.email
}

/**
* @brief set the email of the user
* @param   a string indicate the new email of the user
*/
func (user* User) setEmail(email string) {
	user.email = email
}

/**
* @brief get the phone of the user
* @return   return a string indicate the phone of the user
*/
func (user User) getPhone() string {
	return user.phone
}

/**
* @brief set the phone of the user
* @param   a string indicate the new phone of the user
*/
func (user* User) setPhone(phone string) {
	user.phone = phone
}