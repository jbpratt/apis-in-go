curl -X POST \
  http://api.github.com/gists \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'authToken: ec765b468e647af09c2e08033d022044942a6926' \
  -d ./test.txt
