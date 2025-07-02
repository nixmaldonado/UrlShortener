# URL Shortener

This is a simple URL shortener built with Go and the Gin framework.

It allows users to create short URLs that redirect to longer URLs and tracks the number of redirects for each short URL.

The data is stored in a JSON file (`urls.json`).

## Prerequisites

To run this project locally, ensure you have the following installed:

- **Go**: Version 1.16 or higher ([download](https://golang.org/dl/)).

## Project Structure

- `main.go`: The main application file that sets up the Gin server and routes.
- `handlers.go`: Contains the HTTP handler functions, such as redirect and store operations.
- `storage.go`: Manages the storage and retrieval of URL data from the JSON file.
- `helpers.go`: Includes utility functions to support the application's logic.
- `go.mod`: Go module files for dependency management.

A sample `urls.json` is included with the following content:

```json
{
  "1276058": {
    "url": "https://github.com/",
    "redirect_count": 0
  }
}
```

## Setup Instructions

Follow these steps to run the URL shortener locally:

1. **Clone the Repository (if applicable)**:
    ```bash
    git clone git@github.com:nixmaldonado/urlShortener.git
    cd urlShortener
   ```

2. **Install Dependencies (if needed)**:

   The project uses external dependencies (e.g., Gin).

   Run the following command to download and install them:

   ```bash
   go mod tidy
   ```

3. **Start the server by running**:
    ```bash
    go run .
    ```

   This will start the server on localhost:8081 (the default port).
4.  **Test the Redirect**:

    Open a web browser or use a tool like curl to test the short URL:

    ```bash
    curl -L http://localhost:8081/v1/1276058
    ```

    This should redirect you to https://github.com/.

    The **redirect_count** in urls.json will increment each time the URL is accessed.
