# Basic Blog Application

This is a basic blog application built using Go for the backend and Alpine.js, HTMX, and Tailwind CSS for the frontend (I am very bad at frontend, so it is what it is).
The application allows users to create, update, delete, and view blog posts. 
The idea was to create a simple restful CRUD API with Go and add some frontend features without using Gin or other frameworks. Later version may be rewritten to use Gin, though.
The blog app can be stateless or stateful with SQLite.

**The default username and password are `admin` and `password` in this demo**


## Features (so far)

- **Authorise**: Login page to authorise the user, just the default admin account works at the moment.
- **Create Post**: Create new blog posts by providing a title and content.
- **View Posts**: View a list of all blog posts.
- **Update Post**: Update the title and content of existing blog posts.
- **Delete Post**: Delete blog posts.
- **Toggle Posts**: Toggle the visibility of the list of all blog posts.

## Technologies Used

- **Go**: Backend server to handle API requests.
- **Alpine.js**: Lightweight JavaScript framework for reactive data binding.
- **HTMX**: Library to extend HTML with AJAX capabilities.
- **Tailwind CSS**: Utility-first CSS framework for styling.

## Getting Started

### Prerequisites

- Go [installed](https://go.dev/doc/install) on your local machine.

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Vadoid/go-blog.git
    ```
2. Navigate to the project directory:
    ```sh
    cd go-blog
    ```
3. If you want to run SQLite persistent version, set value to `true` in [main.go](./main.go)
    ```sh
    var persistent = false 
    ```

### Running the Application

1. Start the Go server:
    ```sh
    go run .
    ```
2. Open your web browser and navigate to:
    ```
    http://localhost:8080
    ```

## Usage (if you're confused)

### Creating a Post

1. Fill in the "Title" and "Content" fields in the "Create Post" form.
2. Click the "Create Post" button to submit the form.

### Viewing Posts

1. Click the "Show All Posts" button to toggle the visibility of the posts list.
2. The button text will change to "Hide All Posts" and the list of posts will be displayed.

### Updating a Post

1. Click the "Edit" button under a post to switch to the editing mode.
2. Modify the "Title" and "Content" fields.
3. Click the "Update" button to save the changes.
4. Click the "Cancel" button to exit the editing mode without saving changes.

### Deleting a Post

1. Click the "Delete" button under a post to remove it from the list.


### CLI Authorisations

1. Issue `TOKEN=$(curl -X POST http://localhost:8080/login -d '{"username":"admin","password":"password"}' -H "Content-Type: application/json" | jq -r .token)` to authorise.
2. To test API use `curl -X GET http://localhost:8080/posts -H "Authorization: Bearer $TOKEN"`

### Debug

1. Install Delve - `go install github.com/go-delve/delve/cmd/dlv@v1.22.1`
2. Start debuggung in VSCode.

## Code Structure

- **_main.go_**: Go server code that handles API requests for creating, reading, updating, and deleting posts.
- **_main_test.go_**: Test for `main.go` triggered automatically via Github Actions (for local test do `go test -v`).
- **_db.go_**: Handles DB operations, only kicks in when the persistent var is set to true.
- **_db_test.go_**: Test of DB deployment and CRUD operations. Triggers only if `persistent = true` in `main.go`.
- **_local_test_and_run.sh_**: Bash script to test and run locally.
- **_app.yaml_**: One liner for Google App Engine deployment.
- **_templates/index.html_**: HTML file with Alpine.js and HTMX for the frontend functionality.
- **_launch.json_**: JSON setup file for debugging.



