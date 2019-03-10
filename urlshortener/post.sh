curl -X POST \
  http://localhost:8000/v1/short \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
  "url":"http://meta.stackoverflow.com/questions/118594/data-explorer-truncates-links-after-380-characters"
  }'
