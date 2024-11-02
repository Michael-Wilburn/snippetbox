# Snippetbox
Web application called Snippetbox, which lets people paste
and share snippets of text — a bit like Pastebin or GitHub’s Gists.

| **Method** | **URL Pattern**    | **Handler**   | **Action**                 |
|------------|--------------------|---------------|----------------------------|
| ANY        | /                  | home          | Display the home page      |
| ANY        | /snippet/view?id=1 | snippetView   | Display a specific snippet |
| POST       | /snippet/create    | snippetCreate | Create a new snippet       |

# Test the endpoint throw the terminal.
$ curl -i -X POST http://localhost:4000/snippet/create
$ curl -i -x GET  http://localhost:4000/snippet/view?id=123.
