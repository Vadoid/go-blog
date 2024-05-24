#!/bin/bash

# Check the value of persistent in main.go
if grep -q "var persistent = true" main.go; then
  echo "Persistent mode is enabled. Running all tests, including db_test.go..."
  go test -v -tags=persistent ./...
else
  echo "Persistent mode is disabled. Running tests without db_test.go..."
  go test -v $(find . -name "*.go" ! -name "db_test.go")
fi

if [ $? -ne 0 ]; then
  echo "Tests failed. Aborting."
  exit 1
fi

# Run the application
echo "Starting the application..."
go run main.go db.go
