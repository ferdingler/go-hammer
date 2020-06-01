# go-hammer Command Line Tool

Example:

```bash
cli \
--endpoint https://aws.amazon.com \
--method PUT \
--duration 60 \
--tps 1 \
--payload '{"hello":"world"}' \
--headers '{"content-type":"application/json"}'
```

Example passing a scenario defined in JSON

```json
{
    "duration": 60,
    "tps": 1,
    "hammer": "HTTPHammer",
    "hammerConfig": {
        "method": "GET",
        "endpoint": "https://www.google.com",
        "payload": "Hello World",
        "headers": {
            "content-type": "application/json"
        }
    }
}
```

Then run the CLI: 

```bash
cli run scenario --path ./scenario.json
```