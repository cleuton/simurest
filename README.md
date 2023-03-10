# simurest
Server simulator

[**ENGLISH**](./english.md)

Simula um servidor, por enquanto REST. Você pode especificar: 

```
	simurest [
		[--quiet | -q <não imprimir o disclaimer (default true)>]
		[--method | -m <método http a ser testado (default=GET): GET | POST | HEAD | DELETE | PUT>]
		[--port | -p <tcp port (default=8080)> ]
		[--uri | -u <http URI (default="/")]
		[--status | -s <http Status da response (default=200)]
		[--body | -b <http response body (default: '{"status": "ok"}')]
	]

 ```

Uma uri pode ser exata ou com wildcard (*):
- -u /api/user/1 : A uri tem que ser exata
- -u "/api/user*" : aceita qualquer uri que comece com "/api/user" (se usar * não se esqueça de colocar entre aspas duplas)
