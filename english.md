# simurest
Server simulator

[**ENGLISH**](./english.md)

Simulates a REST server: 

```
	simurest [
		[--quiet | -q <don't print the disclaimer (default false)>]
		[--method | -m <http method to be tested (default=GET): GET | POST | HEAD | DELETE | PUT>]
		[--port | -p <tcp port (default=8080)> ]
		[--uri | -u <http URI (default="/")]
		[--status | -s <http Status of the response (default=200)]
		[--body | -b <http response body (default: '{"status": "ok"}')]
	]

 ```

A uri can be exact or have a wildcard at the end (*):
- -u /api/user/1 : Exact uri expected
- -u "/api/user*" : Accepts anything that begins with "/api/user" (If using * don't forget to encluse the uri in double quotes)
