PORT := 50051

run:
	@echo "Checking for existing server on port $(PORT)..."
	-@lsof -ti tcp:$(PORT) | xargs -r kill -9
	@echo "Generating proto code..."
	buf generate proto
	@echo "Starting server and client..."
	# Start server in background
	go run server/main.go 0.0.0.0:$(PORT) &
	SERVER_PID=$$!
	# Start client in background
	go run client/main.go 0.0.0.0:$(PORT) &
	CLIENT_PID=$$!
	# Trap Ctrl+C to kill both processes
	trap 'echo "Stopping..."; kill $$SERVER_PID $$CLIENT_PID; exit 0' INT
	wait

