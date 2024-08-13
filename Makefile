.PHONY: rm run test build
# This stops the Spin app from running in the background and removes the SQLite data contained in the .spin directory
rm:
	pkill -f 'spin up'
	rm -rf .spin

# This runs the Spin app with a sqlite database--using the provided schema--in the background, allowing for other commands to be run in the same shell
run:
	spin up --sqlite @schema.sql > /dev/null 2>&1 &

# This tests the API via HTTP
test:
	hurl --test test.hurl --very-verbose

build:
	spin build