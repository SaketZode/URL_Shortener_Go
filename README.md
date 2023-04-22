# URL_Shortener_Go
URL shortener and expand shortened URL using golang.

## Instructions to execute the code

1. Clone the repository.
2. cd into the project.
3. Run command `go build`. A binary with name `urlshortener` will be generated.
4. Execute the binary using command `./urlshortener`.


## Instructions to test the code

1. Add some encoded URLs to the application:
`Request type: POST`
`URL: "http://localhost:7070/encode"`
  Sample Request Body:
    ```
    {
        "originalUrl": "https://stackoverflow.com/questions/2377881/how-to-get-a-md5-hash-from-a-string-in-golang"
    }
    ```

2. To decode the URLs encoded in first step:
`Request type: POST`
`URL: http://localhost:7070/decode`
Sample request body:
    ```
    {
        "encodedUrl": "http://short.est/2d4b8956c30260cf9d3ba1c1eacbfa8a"
    }
    ```
