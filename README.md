# Go Commain

### go mod tidy

- It removes unused dependencies.
- It adds any missing module requirements for imports in your code.
- It updates your go.mod and go.sum files to reflect exactly what's needed.

need to install to run env
npm install -g nodemon

npm install --save-dev cross-env

server:
nodemon --watch './\*_/_.go' --signal SIGTERM --exec APP_ENV=dev 'go' run main.goa

## start app

air

it will run and check for file changes

## post with CURL

curl -X POST http://localhost:9000/register -d '{"email": "test@example.com", "password": "password123"}' -H "Content-Type: application/json"

curl -X POST http://localhost:9000/login -d '{"email": "test@example.com", "password": "password123"}' -H "Content-Type: application/json"

# use ari for live file server changes

C:\Users\<YourUsername>\go\bin

## install

go install github.com/cosmtrek/air@latest

C:\Users\<YourUsername>\go\bin

## Step 2: Add Go bin to Windows PATH

## Step 3: Confirm it’s working

air -v

## run air

cd path\to\your\go-project
air
