# FHIR Rest service, Version 2

![alt text][routearch]

## Service Architecture

![alt text][servicearch]

## Building

### Building Docker container

Building the FHIR API depends on the PDX Contract. 
To enable the container build to pull a private Github repo, you need to pass a personal access token.
See the [Github Documentation](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) for how to generate a token.
No additional scopes are required.
Save this token in an environment variable named `GITTOKEN` or substitute it into the `docker build` call below.

```
> docker build . -f build/Dockerfile -t fhirapi --build-arg gittoken=${GITTOKEN}
```


[routearch]: assets/APIv2.png "API Version 2 routes"
[servicearch]: assets/servicearch.png "Service Architecture"
