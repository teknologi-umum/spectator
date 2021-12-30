# Logger

Logger provides standard logging capabilities to log data into the InfluxDB database storage.

## Sending a log

`POST /`

Send a log that will be written to the InfluxDB. It is very recommended to add the `data.level` to the log you're sending, that way we can track everything easier.

You can send the request with either JSON or MessagePack by simply specifying the `Content-Type` header.

It will reject the request with a 418 Teapot if the `Content-Type` header is not specified.

Request headers:
```yaml
Content-Type: ["application/json", "application/msgpack"]
```

Request body schema:
```json5
{
    "access_token": "string (required)",
    "data": {
        "request_id": "string (required)",
        "application": "string (required)",
        "message": "string (required)",
        "body": {}, // Any object, optional
        "level": "string (optional)", // defaults to 'debug'
        "environment": "string (optional)",
        "language": "string (optional)",
        "timestamp": "datetime unix number or iso8601 format or rfc3339 format (optional)" // defaults to current time
    }
}
```

Sample request:

Curl:
```sh
curl --request POST \
  --url http://logger-endpoint/ \
  --header 'Content-Type: application/json' \
  --data '{
    "access_token": "string (required)",
    "data": {
        "request_id": "string (required)",
        "application": "string (required)",
        "message": "string (required)",
        "body": {},
        "level": "string (optional)",
        "environment": "string (optional)",
        "language": "string (optional)",
        "timestamp": "datetime unix number or iso8601 format or rfc3339 format (optional)"
    }
}'
```

## Retrieving logs

`GET /`

Read logs that are in InfluxDB with specified filters from the request URL's query parameters.

Available query parameters:
- `level` - Show logs with the specified level
- `request_id` - Show logs with the specified request_id
- `application` - Show logs with the specified application
- `from` - Unix second timestamp, show logs from the specified time
- `to` - Unix second timestamp, show logs up to the specified time

The response body defaults in JSON format, but you can receive a MessagePack format by specifying the `Accept` header. Either it is explicitly `application/json` or `application/msgpack`.

## Healthchecks / Ping

`GET /ping`

Simple healthcheck endpoint to check if the service and database conenction is up and running.

Curl:
```sh
curl --request GET \
  --url http://logger-endpoint/ping
```

C#:
```csharp
var client = new HttpClient();
var request = new HttpRequestMessage
{
    Method = HttpMethod.Get,
    RequestUri = new Uri("http://logger-endpoint/ping"),
};
using (var response = await client.SendAsync(request))
{
    response.EnsureSuccessStatusCode();
    var body = await response.Content.ReadAsStringAsync();
    Console.WriteLine(body);
}
```

Go:
```go
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "http://logger-endpoint/ping"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
```

Sample response:
```
HTTP/1.1 200 OK
Content-Length: 2
Content-Type: text/plain; charset=utf-8
Body:

pass
```