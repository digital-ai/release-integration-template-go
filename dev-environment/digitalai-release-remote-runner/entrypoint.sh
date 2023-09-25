#!/bin/sh

MAX_WAIT_SECONDS=300
WAIT_INTERVAL=5

start_time=$(date +%s)
end_time=$((start_time + MAX_WAIT_SECONDS))

api_url="http://localhost:5516/tokens/users/remote-runner"
unique_id=$(date +%s%N | md5sum | awk '{print $1}')

while true; do
    current_time=$(date +%s)

    if [ $current_time -ge $end_time ]; then
        echo "Timeout reached. Exiting!"
        exit 1
    fi

    response=$(curl -s -X POST -u admin:admin -H "Content-Type: application/json;charset=UTF-8" -d '{"tokenNote": "'$unique_id'"}' $api_url)

    if [ $? -eq 0 ]; then
        echo "Successful fetched token..."
        export RELEASE_RUNNER_TOKEN=$(echo "$response" | jq -r '.token')
        break
    fi

    echo "Error fetching token! Waiting $WAIT_INTERVAL seconds before trying again..."
    sleep $WAIT_INTERVAL
done

exec /app/release-remote-runner --profiles docker
