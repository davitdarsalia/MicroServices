package constants

const UserSignUpQuery = `INSERT INTO "userAuth" (email, firstname, lastname, age, password) VALUES($1, $2, $3 ,$4 , $5)`

const CheckUser = `select password from "userAuth" where password = $1 and email = $2`
