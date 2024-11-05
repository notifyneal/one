# Golang Web Application Starter Project

This repository contains a sample Golang web application project that serves as a starting point for building your own web applications.

## Features

- Basic HTTP server listening on port 8080 with `Status`, `Echo`, and `Post` handlers, which provide a solid starting point for implementing custom handlers that receive and return JSON data.
- `/status` endpoint for health checks:
  - A simple example of a handler that returns JSON data, which can be used as a starting point for creating more complex handlers that receive and return JSON.
- `/echo` endpoint for echoing request messages:
  - Demonstrates how to read JSON data from a request and return it in the response. This can be adapted for more advanced use cases where processing or modification of the data is needed.
- `/post` endpoint for handling POST requests:
  - Demonstrates how to handle POST requests, log incoming requests, and return the posted message. This can be extended for handling more complex data submissions.


## Getting Started

### Prerequisites

- Golang installed on your system
- Docker installed on your system

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/golang-webapp-starter.git
```
2. Change directory to the project folder:
```bash
cd golang-webapp-starter
```

3. Build the project:

```bash
make build run
```
or
```bash
go build 
```


4. Run the server:
```bash
make  run
```
or
```bash
./server 
```

The server will start listening on port 8080.

## Usage

### Flags

- `--config` or `-c`: Specify the path to the configuration file (default is `$HOME/.server.yaml`)
- `--webroot` or `-w`: Define the web root folder (default is `./web/dist`)

### Endpoints

1. `/status`: Health check endpoint. Returns a JSON object with a `status` key and the value "Ok".

2. `/echo`: Accepts a JSON object with a `message` key and echoes the request message. Logs the entire request and the message.

3. `/post`: Accepts a JSON object with a `message` key via a POST request, logs the request, and returns the message. Useful for testing POST request handling and logging.

## Functions

### Main Functions

- `Execute()`: Executes the root command and initializes the server.
- `initConfig()`: Reads the configuration file and environment variables.

### HTTP Handlers

- `Status()`: HTTP handler for the `/status` endpoint. Returns the health check status.
- `Echo()`: HTTP handler for the `/echo` endpoint. Returns the echoed message from the request.
- `Post()`: HTTP handler for the `/post` endpoint. Handles POST requests and returns the posted message.

## License

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) for CLI support
- [Viper](https://github.com/spf13/viper) for configuration management
