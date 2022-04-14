package constants

const FetchUserId = `select id from "userAuth" where email = $1 and password = $2`

const UserSignUpQuery = `INSERT INTO "userAuth" (email, firstname, lastname, age, password) VALUES($1, $2, $3 ,$4 , $5)`

const CheckUser = `select password from "userAuth" where password = $1 and email = $2`

const FetchUserInfo = `select balance, rating from "userInfo" where userid = $1`

const GetAllUsers = `select id, email, firstname, lastname, age from "userAuth"`

const GetUserByID = `select id, email, firstname, lastname, age  from "userAuth" where id = $1`

const registerInfo = `insert into  "userInfo" (userid, balance, rating) VALUES ($1, $2, $3)`

const IncreaseRatingByID = `update "userInfo" set rating = $1 where userid = $2`

const IncreaseBalance = `update "userInfo" set balance = $1 where userid = $2`

const RegisterTransaction = `insert into  transactions (id, recipient, amount, currency) VALUES ($1, $2, $3, $4)`
