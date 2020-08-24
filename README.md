# enigma_bank_api

Don't forget to create a database, you can use one in the db.sql file

Create a file called ".env" without quote
The content is below

DB_USER={your mysql user} <br />
DB_PASSWORD={your mysql password} <br />
DB_HOST=localhost <br />
DB_PORT=3306 <br />
DB_SCHEMA=enigma_bank <br />
DB_DRIVER=mysql <br />

The endpoints are:
1. GET /logins
2. GET /login/{id}
3. POST /login
4. PUT /login/{id}
5. DELETE /login/{id}
6. GET /users
7. GET/user/{id}
8. POST /user
9. PUT /user/{id}
10. DELETE /user/{id}
11. GET /transactions (need some adjustments)
12. GET /transaction/{id} (need some adjustments)
13. GET /transaction/user/{id} (need some adjustments)
14. POST /transaction (need some adjustments)
15. PUT /transaction/{id} (need some adjustments)
16. DELETE /transaction/{id} (need some adjustments)

For Creating and Updating, the body should be like this:

For Login
{ <br />
&nbsp;&nbsp;&nbsp;&nbsp;"username": "username", <br />
&nbsp;&nbsp;&nbsp;&nbsp;"password": "password", <br />
} <br />

For User
{ <br />
&nbsp;&nbsp;&nbsp;&nbsp;"login_owner_id": {login-owner-id in number}, <br />
&nbsp;&nbsp;&nbsp;&nbsp;"balance": {balance in number}, <br />
} <br />

For Transaction (need some atrans_datedjustments)
{ <br />
&nbsp;&nbsp;&nbsp;&nbsp;"user_owner_id": {login-owner-id in number}, <br />
&nbsp;&nbsp;&nbsp;&nbsp;"trans_date": {balance in number}, <br />
&nbsp;&nbsp;&nbsp;&nbsp;"destination": {balance in number}, <br />
&nbsp;&nbsp;&nbsp;&nbsp;"amount": {balance in number}, <br />
&nbsp;&nbsp;&nbsp;&nbsp;"description": {balance in number}, <br />
} <br />

How to run:
1. Linux / OSX, type this in the root directory, NOT main directory. Type the arguments, localhost then port <br />
&nbsp;&nbsp;&nbsp;&nbsp;go run main/app.go localhost 1234
2. Windows, type this in the root directory, NOT main directory. Type the arguments, localhost then port<br />
&nbsp;&nbsp;&nbsp;&nbsp;go run main/app.go localhost 1234
