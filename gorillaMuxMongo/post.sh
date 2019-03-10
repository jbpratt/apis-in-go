curl -X POST \
  http://localhost:8000/v1/movies \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{ "name": "The Dark Knight", "year": "2008", "directors": ["Christopher Nolan"], "writers": ["Jonathan Nolan", "Christopher Nolan"], "boxOffice": { "budget": 185000000, "gross": 533316061}}'
