# Getir Case

Getir Study Case

## Installation

With Docker:

```sh
$ docker build -t getir .
$ docker run -d -p 8080:8080 --name getir getir
```

## What's Added?
- Postman Test's
- Go Test's

## Handler Structure

- `/search  | http://52.214.126.26:8080/search`
- `/holder | http://52.214.126.26:8080/holder`



# Mongo Search 

**URL** : `/search`

**Method** : `POST`

**Data example**

```json
{
  "startDate": "2016-01-21",
  "endDate": "2016-03-02",
  "minCount": 2900,
  "maxCount": 3000
}
```

## Success Response

**Code** : `200`

**Content example**

```json
{
  "code": 0,
  "msg": "Success",
  "records": [
    {
      "createdAt": "2016-02-19T08:35:39.409+02:00",
      "key": "kkxEdhft",
      "totalCount": 2980
    }
  ]
}
```

* [Set Holder]() : `POST /holder`

# Set Data

**URL** : `/holder`

**Method** : `POST`

**Data example**

```json
{
    "key": "example",
    "value": "example"
}
```

## Success Response

**Code** : `201`

**Content example**

```json
{
  "key": "example",
  "value": "example"
}
```
* [Set Holder]() : `GET /holder`

# Get Data
**URL** : `/holder`

**Method** : `GET`

**Data example**

```
Header key : example
```

## Success Response

**Code** : `200`

**Content example**

```json
{
  "key": "example",
  "value": "example"
}
```