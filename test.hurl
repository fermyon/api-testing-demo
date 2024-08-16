################################################
# Testing POST
################################################

POST http://localhost:3000/user/user1
HTTP 202

POST http://localhost:3000/user/user2
HTTP 202

POST http://localhost:3000/user/user3
HTTP 202

################################################
# Testing GET
################################################

#-----------------------------------------------
# Listing all users
#-----------------------------------------------

GET http://localhost:3000/all_users
HTTP 200
[Asserts]
jsonpath "$" count == 3
jsonpath "$[0].id" == 1
jsonpath "$[1].id" == 2
jsonpath "$[2].id" == 3
jsonpath "$[0].userName" == "user1"
jsonpath "$[1].userName" == "user2"
jsonpath "$[2].userName" == "user3"
jsonpath "$[0].email" == "user1@example.com"
jsonpath "$[1].email" == "user2@example.com"
jsonpath "$[2].email" == "user3@example.com"


#-----------------------------------------------
# Getting user by ID
#-----------------------------------------------

GET http://localhost:3000/user/1
HTTP 200
[Asserts]
jsonpath "$.id" == 1
jsonpath "$.userName" == "user1"
jsonpath "$.email" == "user1@example.com"

################################################
# Testing DELETE
################################################

#-----------------------------------------------
# Deleting user and validating
#-----------------------------------------------

DELETE http://localhost:3000/user/3
HTTP 202

GET http://localhost:3000/user/3
HTTP 404

################################################
# Testing misc errors
################################################

#-----------------------------------------------
# Creating a duplicate user
#-----------------------------------------------

POST http://localhost:3000/user/user2
HTTP 400

#-----------------------------------------------
# Trying a non-existent route
#-----------------------------------------------

GET http://localhost:3000/not_a_real_route
HTTP 404

#-----------------------------------------------
# Trying a non-approved HTTP method
#-----------------------------------------------

PUT http://localhost:3000/user
HTTP 404

#-----------------------------------------------
# Sending a non-integer ID
#-----------------------------------------------

GET http://localhost:3000/user/asdf
HTTP 400

DELETE http://localhost:3000/user/asdf
HTTP 400