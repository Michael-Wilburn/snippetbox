# Snippetbox
Web application called Snippetbox, which lets people paste
and share snippets of text — a bit like Pastebin or GitHub’s Gists.

| **Method** | **URL Pattern**    | **Handler**     | **Action**                   |
|------------|--------------------|-----------------|------------------------------|
| ANY        | /                  | home            | Display the home page        |
| ANY        | /snippet/view?id=1 | snippetView     | Display a specific snippet   |
| POST       | /snippet/create    | snippetCreate   | Create a new snippet         |
| ANY        | /static/           | http.FileServer | Serve a specific static file |

# Test the endpoint throw the terminal.
* $ curl -i -X GET  http://localhost:4000/
* $ curl -i -X POST http://localhost:4000/snippet/create
* $ curl -i -X GET  http://localhost:4000/snippet/view?id=123.
* $ curl -iL -X POST http://localhost:4000/snippet/create


- Using the double arrow >> will append to an existing file, instead of truncating it
when starting the application.
$ go run ./cmd/web >>/tmp/info.log 2>>/tmp/error.log