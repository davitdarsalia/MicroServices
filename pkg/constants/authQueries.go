package constants

const UserSignUpQuery = "INSERT INTO userregistration(email, firstname, lastname, age, password) VALUES($1, $2, $3 ,$4 , $5)"
const CheckUser = "SELECT  email, password FROM userregistration WHERE email = $1 AND  password = $2"
