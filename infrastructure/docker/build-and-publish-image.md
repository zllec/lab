# Build and Publish a Container Image
1. Build
```bash
docker build -t sample.registry.com/foobar:v1.0.0 .
```
2. Login to registry
```bash
docker login <registry_url>
```
3. Push
```bash
docker push sample.registry.com/foobar:v1.0.0
```
