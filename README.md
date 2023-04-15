# PageSpeed Commmand

Get the performance score of a website with the Google PageSpeed API on the command line.

## Requirements

Create a Google PageSpeed API key on your Google Cloud account (https://console.cloud.google.com/apis/credentials)

Set the API key in a `.env` file in the root of the project or in `.env.local` file.

## Usage

```bash
# run with default strategy
$ go run main.go -url https://www.example.com
{"url":"https://www.example.com","strategy":"DESKTOP","score":1,"date":"2023-04-01T01:01:01.352859963+02:00"}

# run with mobile strategy
$ go run main.go -url https://www.example.com -strategy mobile

# build binary in ./build directory
$ make build
$ ./build/pagespeed -url https://www.example.com
```

## Doc

* [Google Insight API doc](https://developers.google.com/speed/docs/insights/v5/get-started)
* [Go Google API v5](https://pkg.go.dev/google.golang.org/api/pagespeedonline/v5)
* [joho/godotenv library](https://pkg.go.dev/github.com/joho/godotenv)
