build:
	- cd functions/riot-api && \
		go build -o ../netlify/functions/riot-api main.go
