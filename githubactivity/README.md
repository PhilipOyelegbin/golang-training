# Go GitHub User Activity

## 📝 Project Description

**Go GitHub User Activity** is a simple application programming interface (API) to fetch the recent activity and repository of a GitHub user and display it on the browser. This project will help you practice your programming skills, including working with APIs, handling JSON data, and building a simple API application.

## ✨ Requirements

The application should run from the browser, accept the GitHub username as a parameter, fetch the user’s recent activity or repositry based on the route using the GitHub API, and display it on the browser. The user should be able to:

- Provide the GitHub username as a parameter via the browser url.

- Fetch the recent activity or repository of the specified GitHub user using the GitHub API. You can use the following endpoint to fetch the user’s activity or repos:

```
# https://api.github.com/users/<username>/events
# https://api.github.com/users/<username>/repos

# Example: https://api.github.com/users/kamranahmedse/events
# Example: https://api.github.com/users/kamranahmedse/repos
```

- Display the fetched activity or repository on the browser: You can l[earn more about the GitHub API here](https://docs.github.com/en/rest/activity/events?apiVersion=2022-11-28).

- Handle errors gracefully, such as invalid usernames or API failures.

- Use a programming language of your choice to build this project.

- Do not use any external libraries or frameworks to fetch the GitHub activity.

If you are looking to build a more advanced version of this project, you can consider adding features like filtering the activity by event type, displaying the activity in a more structured format, or caching the fetched data to improve performance. You can also explore other endpoints of the GitHub API to fetch additional information about the user or their repositories.

## 🛠️ Core Technologies Used

- Go (Golang): The primary programming language.

- Standard Library Packages:

  - fmt: For formatted input and output.

  - io: For basic interfaces to I/O primitives.

  - net/http: Provides HTTP client and server implementations.

  - github.com/gorilla/mux: Implements a request router and dispatcher for matching incoming requests to their respective handler.

  - html/template: Implements data-driven templates for generating HTML output safe against code injection

## 🚀 Installation

To get Go GitHub User Activity up and running on your local machine, follow these steps:

1. Ensure Go is Installed:
   Make sure you have Go installed (version 1.18 or higher is recommended). You can download it from go.dev/dl/.
   Verify your installation:

```
go version
```

2. Clone the Repository (or create project manually):
   If you're starting from scratch as part of a learning exercise, you'd create the project structure manually as described in the task instructions. If this were a real repository:

```
git clone https://github.com/philipoyelegbin/golang-training.git/githubactivity
cd githubactivity
```

3. Initialize Go Module (if not already done):

```
go mod init github.com/philipoyelegbin/golang-training.git/githubactivity # Only if you created the project manually
```

4. Build the Executable:
   This command compiles your Go source code into a single executable binary.

```
go build -o githubactivity
```

This will create an executable file named githubactivity in your project's root directory.

## 💡 Usage

Once built, you can run the CLI commands from your terminal.

**General Usage**

```
./githubactivity
```

**Endpints**

1. `/` - Home page of the API

2. `/events/{username}` - Get user event activities

3. `/repos/{username}` - Get user repositories

## 📂 Project Structure

```
githubactivity/
    ├── main.go     # Main entry point and route functions handling
    ├── index.html     # Landing page for the API
```

**File Breakdown:**

- `main.go`:

  - Contains the main function, which is the entry point of the application.

  - Includes some route functions to fetch users from events, and repositories.

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
