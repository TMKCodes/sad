#!/bin/bash
rm -rf /tmp/sad-temp

NUM_CLIENTS=128
sad --devnet --appdir=/tmp/sad-temp --profile=6061 --rpcmaxwebsockets=$NUM_CLIENTS &
SAD_PID=$!
SAD_KILLED=0
function killSadIfNotKilled() {
  if [ $SAD_KILLED -eq 0 ]; then
    kill $SAD_PID
  fi
}
trap "killSadIfNotKilled" EXIT

sleep 1

rpc-idle-clients --devnet --profile=7000 -n=$NUM_CLIENTS
TEST_EXIT_CODE=$?

kill $SAD_PID

wait $SAD_PID
SAD_EXIT_CODE=$?
SAD_KILLED=1

echo "Exit code: $TEST_EXIT_CODE"
echo "sad exit code: $SAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-idle-clients test: PASSED"
  exit 0
fi
echo "rpc-idle-clients test: FAILED"
exit 1
