# butku - a simple url shortner made using go and redis

## Installation
1. Clone the repository on your machine
2. create a .env file in the root of the cloned directory
3. Add the following entries to your .env file
   1. `REDIS_ADDR` - the address for your redis database along with its port. Example: `localhost:6969`
   2. `REDIS_PASS` - the password to your redis database
   3. `SERVE_PORT` - Port through which the api will be served
4. Make sure to add .env file to your .gitignore file to protect you redis keys
5. build the project using `go build`
6. run using `./butku` on linux or mac and `.\butku.exe` on windows

## Usage
1. to shorten a url send a `POST` request to `/shorten` along with url in JSON format as such -
```json
{"url":"https://example.com"}
```
2. Send a `GET` request to `/{shortURL}` to get redirected