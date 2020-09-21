TODO
- Add preflight interceptor to use in the Handler - Issue also raised as a reminder

Set `X-Proxy-Target` header to the host destination
eg: `X-Proxy-Target: https://jsonplaceholder.typicode.com`

Then use `http://localhost:1412/todos` for example, to return the JSON data from https://jsonplaceholder.typicode.com/todos
