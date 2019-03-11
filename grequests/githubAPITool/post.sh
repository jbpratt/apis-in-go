curl -X POST \
  http://api.github.com/gists \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'authToken: xxxxxxxxxxxxxxxxxxxxxxxxxxx' \
  -d ./test.txt
