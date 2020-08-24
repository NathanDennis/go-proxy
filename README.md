A simple Go proxy server

Set `X-Proxy-Target` header to the host destination
eg: `X-Proxy-Target: https://jsonplaceholder.typicode.com`

Then use `http://localhost:1412/todos` to return the JSON data from https://jsonplaceholder.typicode.com/todos

TODO:
- Add logic to Director to set base path
- Define func to join the URL path
- Flesh out README for clearer overview and examples