curl -X POST \
  http://localhost:8000/v1/user \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
    "username": "mister",
    "email_address": "mister@autumn.com",
    "first_name": "autumn",
    "last_name": "mister"
}'
