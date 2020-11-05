# Let's Go!

This is my first Golang project.  
This project is about implementing basic CRUD operations using Golang, which is an assignment of QSC backend team.  
Most of the code segments in `db.go` are provided by @CHN-ChenYi   
The aim of this project is to create a simple server which has the following API:  

## POST /create

Sample request:
```
/create?id=RalXYZ&age=18
```

Sample response:
```
uid = 1, RalXYZ, 18
```

## GET /retrieve

Only supports retrieving by uid

Sample request
```
/retrieve?uid=1
```

Sample response:
```
uid = 1, RalXYZ, 18
```

## PUT /update

Sample request:
```
/update?uid=1&id=RalXYZ&age=17
```

Sample response:
```
uid = 1, RalXYZ, 17
```

## DELETE /delete

Only supports deleting by uid

Sample request
```
/delete?uid=1
```
