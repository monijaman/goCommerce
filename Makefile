server:
	npx cross-env APP_ENV=dev nodemon --watch './**/*.go' --signal SIGTERM --exec "go run main.go"
