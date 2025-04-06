# API-Challenges-Points-of-Interest

This repository is part of a set of challenges that are available in the profile <a href="https://github.com/backend-br/desafios/" target="_blank">backend-br</a>.

## Challenge

Implement a service for XY Inc., a company that specializes in producing excellent GPS (Global Positioning System) receivers. The company's management is working hard to launch an innovative device that promises to help people locate points of interest (POIs), and they really need your help. You have been hired to develop the platform that will provide all the intelligence for the device. This platform should be based on REST services, to make integration flexible.


## Routes

### `POST /register`

Registers a new point in the database.

#### Request
```json
{
    "name": "Lanchonete",
    "x": 10,
    "y": 15
}
```

#### Response
```json
{
    "id": 1,
    "name": "Lanchonete",
    "x": 10,
    "y": 15
}
```

### `GET /list`

Lists all registered points. No data to make the request.

#### Response
```json
{
    "success": [
        // list of points
    ]
}
```

### `GET /get-points`

Returns all points within a maximum distance.

#### Request

```json
{
    "d_max": 10, // maximum distance
    "x": 10, // x and y are the reference points
    "y": 15
}
```

#### Response
```json
{
    "success": [
        // points that are within the maximum distance
    ]
}
```