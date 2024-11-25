# Snippetbox
Web application called Snippetbox, which lets people paste
and share snippets of text — a bit like Pastebin or GitHub’s Gists.

| **Method** | **URL Pattern**   | **Handler**       | **Action**                                    |
|------------|-------------------|-------------------|-----------------------------------------------|
| GET        | /                 |  home             | Display the home page                         |
| GET        | /snippet/view/:id | snippetView       | Display a specific snippet                    |
| GET        | /snippet/create   | snippetCreate     | Displa a HTML form for crating a new snippet  |
| POST       | /snippet/create   | snippetCreatePost | Create a new snippet                          |
| GET        | /user/signup      | userSignup        | Display a HTML form for signing up a new user |
| POST       | /user/signupPost  | userSignupPost    | Create a new user                             |
| GET        | /user/login       | userLogin         | Display a HTML form for logging in a user     |
| POST       | /user/login       | userLoginPost     | Authenticate and login the user               |
| POST       | /user/logout      | userLogoutPost    | Logout the user                               |
| GET        | /static/*filepath | http.FileServer   | Serve a specific static file                  |

# Test the endpoint throw the terminal.
* $ curl -i -X GET  http://localhost:4000/
* $ curl -i -X GET  http://localhost:4000/snippet/view/123.
* $ curl -i -X GET http://localhost:4000/snippet/create
* $ curl -iL -X POST http://localhost:4000/snippet/create
* $ curl -i -X GET http://localhost:4000/static/


- Using the double arrow >> will append to an existing file, instead of truncating it
when starting the application.
$ go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log