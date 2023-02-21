## Multithreading

Get the fastest result between two different APIs. The requests will be made simultaneously to the following APIs:
* `https://cdn.apicep.com/file/apicep/" + zip + ".json`
* `http://viacep.com.br/ws/" + zip + "/json/`

**Acceptance Criteria:**

1. Accept the API that provides the fastest response and discard the slowest response.
2. The result of the request must be displayed on the command line, as well as which API sent it.
3. Limit response time to 1 second. Otherwise, the timeout error should be displayed.
