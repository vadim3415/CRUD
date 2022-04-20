# Starting the server and postgreSQL

#
1) docker pull postgres
2) docker run --name=http31-db -e POSTGRES_PASSWORD=qwerty -p 5432:5432 -d --rm

3) exec -it container_id /bin/bash
4) psql -U postgres

5) CREATE TABLE users (
id serial not null unique,
name VARCHAR (50) not null unique,
age VARCHAR (50) not null,
friends text[]
);

6) see table
   \d users

7) go run cmd/main.go

8) see users
   localhost:8083
   you output: {"data":null}

9)create three users:->
   curl -X POST -d "{\"name\":\"Sam\", \"age\": 25}" http://localhost:8083/create

10) make friends:->
    curl -X POST -d "{\"source_id\": 1, \"target_id\": 3}" http://localhost:8083/make_friends

11) age update:->
    curl -X PUT -d "{\"new_age\": 10}" http://localhost:8083/age_updated/1

12) delete user:->
    curl -X DELETE http://localhost:8083/delete/2
13) see friends:->
    localhost:8083/friends/1

