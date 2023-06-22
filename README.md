# api

## Endpoints

### [GET] /

Returns a simple Hello World message.

### [GET] /ping

Returns a simple pong message. This endpoint is used to check if the server is up and running.

### [GET] /fetch/:domain

Returns a JSON object containing information about the given domain.

> :warning: **The domain needs to be URL encoded.**

Sample response:
```json
{
    "request_timestamp": "2023-06-22T17:56:21.952833733Z",
    "url": {
        "full_url": "https://user:pass@domain.tld",
        "protocol": "https",
        "subdomain": "",
        "hostname": "domain.tld",
        "domain": "domain",
        "tld": "tld",
        "port": 443,
        "path": "",
        "query": null,
        "fragment": "",
        "username": "user",
        "password": "pass",
        "ip_address": "123.456.789.012"
    },
    "http_status_code": 200,
    "tls_certificate": {
        "subject": "domain.tld",
        "issuer": "CN=R3,O=Let's Encrypt,C=US",
        "expires": "1503-12-24 05:04:46",
        "is_valid": true
    },
    "ping_time": 541.6001550,
    "http_response_time": 1515.5849951,
    "dns": {
        "mx": [
          "domain.tld."
        ],
        "txt": [
          "v=spf1 a mx -all"
        ],
        "a": [
          "123.456.789.012"
        ],
        "cname": [
          "domain.tld."
        ],
        "ns": [
            "nameserver1.otherdomain.tld.",
            "nameserver2.otherdomain.tld."
        ]
    }
}
```