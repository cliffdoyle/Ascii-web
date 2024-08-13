# ASCII Art Web

## Table of Contents

- Description
- Usage
- Tests
- Authors
- References


## Description

The ASCII Art Web Generator is a web application that converts user input text into ASCII art. Users can enter text, select a font style, and view the generated ASCII art directly on the web page. The application is built using Go and serves the web pages and handles requests efficiently.


## Usage

### How to Run

1. Clone the repository:
    ```sh
    git clone https://learn.zone01kisumu.ke/git/masman/ascii-art-web.git
    cd ascii-art-web
    ```

2. Install Go and set up your Go environment if you haven't already. You can download and install Go from the official [Go website](https://golang.org/dl/).

3. Run the application:
    ```sh
    go run main.go
    ```

4. Open your web browser and navigate to `http://localhost:5000/`.

### Implementation Details: Algorithm

The ASCII Art Generator utilizes a series of Go packages to manage different aspects of the application. Here's an overview of the implementation:

1. **Main Server Setup**:
    - The main function sets up an HTTP server that listens on port 5000. It defines several route handlers:
        - `/` for the homepage
        - `/submit` for handling form submissions

2. **Frontend HTML**:
    - The HTML file serves as the user interface for the application. It includes a form where users can input text and select a font file. The result area displays the generated ASCII art.

3. **ASCII Art Generation**:
    - The core logic for converting text into ASCII art is handled in the `functions` package. The `Ascii` function reads the font data, processes the input text, and generates the corresponding ASCII art.

4. **Backend Handlers**:
    - The backend handlers manage the routing and processing of user requests. The `HandleRequest` function handles form submissions, generates ASCII art, and returns the result to be displayed on the web page.

```With these details, you should be able to understand and run the ASCII Art Web Generator application effectively. Feel free to explore the code and modify it according to your needs.```

## Tests

- Test file is provided that tests the functionality of the function used to generate ascii art characters from the font files.
- `` go test`` to run the test file

## Authors

- Malika
- Cliff

## References

``ASCII MANUAL and HTML``
