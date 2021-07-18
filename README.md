# Example App For Demoing API Auth

## Build

```
docker-compose build
```

## Run

```
docker-compose up
```

## Acquire Client Certificate

```
# Run following command and make note of container id
docker ps

# Copy certs out of container
docker cp <container-id>:/client-adhoc-client-key.pem .
docker cp <container-id>:/client-adhoc-client.pem .
```

## Testing Client Certificate Authentication

```
# This should fail
curl https://127.0.0.1/api/v1/bar

# This should work
# NOTE: -k is used because we are using a self-signed cert for the server and for testing purposes
curl -k \
  --cert client-adhoc-client.pem \
  --key client-adhoc-client-key.pem \
  "https://127.0.0.1/api/v1/bar"
```

