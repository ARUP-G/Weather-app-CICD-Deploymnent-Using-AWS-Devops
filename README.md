# Weather App

This is a simple web application that allows users to check the weather of a city. It consists of a Go backend server that communicates with the OpenWeatherMap API to fetch weather data and serves static HTML files for the frontend.

## Features

- Users can check the current weather of any city.
- Provides a welcome page and a weather form for easy navigation.

## Prerequisites

- Go programming language installed ([Installation guide](https://golang.org/doc/install))
- OpenWeatherMap API key ([Sign up](https://home.openweathermap.org/users/sign_up))

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/ARUP-G/Weather-app.git
    ```

2. Navigate to the project directory:

    ```bash
    cd Weather-app
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

4. Set up your OpenWeatherMap API key:

    Replace the placeholder `"OPENWEATHERMAP_API_KEY"` in `main.go` with your actual API key in the `.env` file.

5. Run the Go server:

    ```bash
    go run main.go
    ```

6. Open a web browser and navigate to `http://localhost:8080/` to access the welcome page.

## Usage

1. Welcome Page:

    - Upon accessing `http://localhost:8080/`, users will see a welcome message and a link to check the weather.

2. Weather Form:

    - Click on the "Check Weather" link to navigate to the weather form (`http://localhost:8080/weather-form.html`).
    - Enter the name of the city you want to check the weather for in the input field and click "Get Weather".
    - The current temperature and description of the weather will be displayed below the form.

## Directory Structure
```
/Weather-app
    ├── static
    │ ├── index.html
    │ └── weather-form.html
    └── main.go
```
## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## License

This project is licensed under the [MIT License](LICENSE).
