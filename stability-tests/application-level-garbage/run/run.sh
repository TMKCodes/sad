#!/bin/bash
rm -rf /tmp/sad-temp

sad --devnet --appdir=/tmp/sad-temp --profile=6061 --loglevel=debug &
SAD_PID=$!
SAD_KILLED=0
function killSadIfNotKilled() {
    if [ $SAD_KILLED -eq 0 ]; then
      kill $SAD_PID
    fi
}
trap "killSadIfNotKilled" EXIT

sleep 1

application-level-garbage --devnet -alocalhost:22611 -b blocks.dat --profile=7000
TEST_EXIT_CODE=$?

kill $SAD_PID

wait $SAD_PID
SAD_KILLED=1
SAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "sad exit code: $SAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SAD_EXIT_CODE -eq 0 ]; then
  echo "application-level-garbage test: PASSED"
  exit 0
fi
echo "application-level-garbage test: FAILED"
exit 1
