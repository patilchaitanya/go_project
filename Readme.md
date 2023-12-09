# Go Server with Concurrent and Sequential Sorting

This is a simple Go server that provides two endpoints for sorting arrays: one for sequential processing and another for concurrent processing.

## Prerequisites

Make sure you have the following software installed on your machine:

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)

## Getting Started

Follow these steps to run the Go server on your machine:

1. Clone the repository:

    ```bash
    git clone https://github.com/patilchaitanya/go_project.git
    cd go_project
    ```

2. Build the Docker image: (Make sure your Docker Desktop (Docker Nodemon) is Running in Background)

    ```bash
    docker build -t go-server .
    ```
## Getting Started

Follow these steps to run the Go server on your machine:

1. Clone the repository:

    ```bash
    git clone https://github.com/patilchaitanya/go_project.git
    cd go_project
    ```

2. Run the Go server locally:

    ```bash
    go run hello.go
    ```

    Your server will be running at [http://localhost:8000](http://localhost:8000).

3. Test the server locally using cURL or Postman. Here's an example using cURL:

    ```bash
    # Process Single Endpoint
    $body = '{"to_sort": [[10, 9, 8, 7, 6, 5, 4, 3, 2, 1], [20, 19, 18, 17, 16, 15, 14, 13, 12, 11]]}'
    Invoke-RestMethod -Uri 'http://localhost:8000/process-single' -Method Post -Body $body -Headers @{"Content-Type"="application/json"}
    ```

    Replace the `$body` content with your desired input arrays.

    ```bash
    # Process Concurrent Endpoint
    $body = '{"to_sort": [[10, 9, 8, 7, 6, 5, 4, 3, 2, 1], [20, 19, 18, 17, 16, 15, 14, 13, 12, 11]]}'
    Invoke-RestMethod -Uri 'http://localhost:8000/process-concurrent' -Method Post -Body $body -Headers @{"Content-Type"="application/json"}
    ```

    Replace the `$body` content with your desired input arrays.

4. Run the Docker container:

    ```bash
    docker run -p 8000:8000 go-server
    ```

5. The server should now be running. Access the following endpoints:

    - [http://localhost:8000/process-single](http://localhost:8000/process-single) (Sequential Processing)
    - [http://localhost:8000/process-concurrent](http://localhost:8000/process-concurrent)

## Usage

To test the server with your arrays, you can use tools like [cURL](https://curl.se/) or [Postman](https://www.postman.com/). Here's an example using cURL:

### Process Single Endpoint

```bash
curl -X POST -H "Content-Type: application/json" -d '{"to_sort": [[3, 1, 4], [1, 5, 9], [2, 6, 5]]}' http://localhost:8000/process-single
