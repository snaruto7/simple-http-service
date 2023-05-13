
# Response Service

This is sample HTTP Response Service in Go which is using Gin for routing.


## Pre-Requisite

This service will be used to track Pod Name and Zone of Caller-Service and also publishes it's own Pod Name and Zone on which it is running. Having GKE Cluster for it's deployment is necessary otherwise `/data` API call will fail and pod will get terminated.
## API Reference

#### Get Service Starting API

```http
  GET /
```
Returns `Hello from Response Service`

#### Get Service Ping 

```http
  GET /ping
```
Returns JSON object 
`{"message":"pong"}`

#### Get Pod and Zone information

```http
  GET /data
```

#### Query Parameters

 Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. Name of the Caller Service Pod|
| `zone` | `string` | **Required**. Zone of the Caller Service Pod|

Returns JSON of **Caller-Service** Pod Name and Zone and **Response-Service** Pod Name and zone
Example
```
{
    "SourcePodName": "",
    "SourceNodeZone": "",
    "DestPodName": "",
    "DestNodeZone": "",
}
```

## Usage

Response service has 3 Apis as mentioned above.
When Caller-Service Pod comes up it will call `/data` API of response service with Query parameters of it's own Pod Name and Zone in every 3 seconds.

Deploy the current service using [ Deployment Manifest ](https://github.com/snaruto7/simple-http-service/blob/master/response-service/deployment.yaml) present in the folder in **GKE Cluster** to get it working.
