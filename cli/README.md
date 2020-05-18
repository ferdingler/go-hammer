# go-hammer Command Line Tool

Example:

```bash
cli run \
--endpoint https://aws.amazon.com \
--method PUT \
--duration 60 \
--tps 1 \
--payload '{"hello":"world"}' \
--headers '{"content-type":"application/json"}'
```