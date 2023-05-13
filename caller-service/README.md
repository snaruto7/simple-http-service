# Caller Service

This is sample HTTP Caller Service in Go which originates request every 3 seconds to [ Response Service ](https://github.com/snaruto7/simple-http-service/tree/master/response-service)
## Pre-Requisite

This service will be used to fetch Pod Name and Zone using [Google Metadata APIs](https://developers.google.com/analytics/devguides/reporting/metadata/v3) and passes the info to response service. Having GKE Cluster for it's deployment is necessary otherwise `/data` API call will fail and pod will get terminated.
## Usage

When Caller-Service Pod comes up it will call Response Service on `/data`  with Query parameters of it's own Pod Name and Zone every 3 seconds.

Deploy the current service using [ Deployment Manifest ](https://github.com/snaruto7/simple-http-service/blob/master/caller-service/deployment.yaml) present in the folder in **GKE Cluster** to get it working.
