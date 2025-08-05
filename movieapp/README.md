# Go Movie App

## 📝 Project Description

**Go Movie App** is a simple full-stack app to show data from The Movie Database. This project will help you practice your programming skills, including working with APIs, and handling JSON data.

## ✨ Requirements

The application should run from the browser, and be able to pull and show the movies based on the country selected. The user should be able to specify the movies they want to see by a dropdown button.

You can look at the API documentation to understand how to fetch the data for each type of movie.

- [RapidAPI: Streaming-availability](https://rapidapi.com/movie-of-the-night-movie-of-the-night-default/api/streaming-availability/playground/apiendpoint_14b2f4b9-8801-499a-bcb7-698e550f9253)

There are some considerations to keep in mind:

- Handle errors gracefully, such as API failures or network issues.
- Use a programming language of your choice to build this project.
- Make sure to include a README file with instructions on how to run the application and any other relevant information.

## 🛠️ Core Technologies Used

- Go (Golang): The primary programming language.

- Standard Library Packages:

  - fmt: For formatted input and output.

  - io: For basic interfaces to I/O primitives.

  - os: For operating system interactions, including file system operations.

  - net/http: Provides HTTP client and server implementations.

  - github.com/gorilla/mux: Implements a request router and dispatcher for matching incoming requests to their respective handler.

  - html/template: Implements data-driven templates for generating HTML output safe against code injection

  - log: Implements a simple logging package.

  - github.com/joho/godotenv: For loading the env vars from a .env file

  - encoding/json: For marshaling (saving) and unmarshaling (loading) Go structs to/from JSON.

  - strings: For string manipulation (e.g., trimming, splitting, case conversion).

## 🚀 Installation

To get Go Movie App up and running on your local machine, follow these steps:

1. Ensure Go is Installed:
   Make sure you have Go installed (version 1.18 or higher is recommended). You can download it from go.dev/dl/.
   Verify your installation:

```
go version
```

2. Clone the Repository (or create project manually):
   If you're starting from scratch as part of a learning exercise, you'd create the project structure manually as described in the task instructions. If this were a real repository:

```
git clone https://github.com/philipoyelegbin/golang-training.git/movieapp
cd movieapp
```

3. Initialize Go Module (if not already done):

```
go mod init github.com/philipoyelegbin/golang-training.git/movieapp     # Only if you created the project manually
```

4. Build the Executable:
   This command compiles your Go source code into a single executable binary.

```
go build -o movieapp
```

This will create an executable file named movieapp in your project's root directory.

## 💡 Usage

Once built, you can run the CLI commands from your terminal.

**General Usage**

```
./movieapp
```

Load the Home page of the app via the browser

## 📂 Project Structure

```
movieapp/
   ├── main.go    # Main entry point and route functions handling
   ├── .env    # Environmental variables for the app
   functions/
      ├── movie.go      # Helper functions used in the main function
   schema/
      ├── types.go     # Struct declaration used in the main file
   templates/
      ├── index.html     # Landing page for the app
```

## 🤝 Contributing

Contributions are welcome! If you'd like to contribute, please:

1. Fork the repository.

2. Create a new branch (git checkout -b feature/your-feature-name).

3. Make your changes.

4. Commit your changes (git commit -m 'feat: Add new feature').

5. Push to the branch (git push origin feature/your-feature-name).

6. Open a Pull Request.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details.
