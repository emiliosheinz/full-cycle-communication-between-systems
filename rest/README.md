# Rest

Representational State Transfer (REST) is an architectural style that defines a set of constraints to be used for creating web services. Web Services that conform to the REST architectural style, or RESTful web services, provide interoperability between computer systems on the Internet. REST-compliant web services allow the requesting systems to access and manipulate textual representations of web resources by using a uniform and predefined set of stateless operations.

## Richardson Maturity Model

The Richardson Maturity Model is a model for classifying the maturity of a REST API. The model was introduced by Leonard Richardson in 2008.

### Level 0 - The Swamp of POX

The first level is the swamp of POX (Plain Old XML). This level is not RESTful at all. It uses HTTP as a transport protocol, but it doesn't use any of the other features of HTTP. It uses HTTP as a tunnel to send XML messages. RPC (Remote Procedure Call) is a common example of this level.

### Level 1 - Resources

The second level is the resources level. This level uses HTTP methods to manipulate resources. It uses the HTTP methods GET, POST, PUT and DELETE to manipulate resources. This level is not RESTful yet, because it doesn't use hypermedia.

### Level 2 - HTTP Verbs

The third level is the HTTP verbs level. This level uses HTTP verbs to manipulate resources. It uses the HTTP methods GET, POST, PUT, DELETE and PATCH to manipulate resources. This level is not RESTful yet, because it doesn't use hypermedia. Different from the second level, this level uses the HTTP verbs correctly, in other words it doesn't use POST to GET data for example.

### Level 3 - Hypermedia as the Engine of Application State (HATEOAS)

The fourth level is the HATEOAS level. This level uses hypermedia to manipulate resources. It uses the HTTP methods GET, POST, PUT, DELETE and PATCH to manipulate resources. This level is RESTful, because it uses hypermedia. The client doesn't need to know the URL of the resources, it only needs to know the URL of the entry point and use it to get the URL of the resources.

Example:
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "account": {
      "number": 12345,
      "balance": {
        "currency": "USD",
        "value": 100.00
      },
      "links": {
        "deposit": "/accounts/12345/deposit",
        "withdraw": "/accounts/12345/withdraw",
        "transfer": "/accounts/12345/transfer"
        "close": "/accounts/12345/close"
      }
    }
}
```
