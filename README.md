    #How to get started


1. Open terminal and get following packages: 
   ```
   go get github.com/joho/godotenv
   go get github.com/go-chi/chi
   go get github.com/go-chi/cors
   go get github.com/lib/pq
   go get github.com/google/uuid

   go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
   go install github.com/pressly/goose/v3/cmd/goose@latest
   ```

2. Run command:
   ```
   cd sql/schema
   ```

3. Run command:
   ```
   goose postgres postgres://docker:docker@localhost:5433/exampledb?sslmode=disable up
   ```

4. Run command:
   ```
   cd ../../
   ```

5. Run command:
   ```
   sqlc generate
   ```

6. Run command: 
   ```
   docker-compose up -d 
   ```

7. Open pgadmin4 at localhost:5050 (login: pgadmin4@pgadmin.org, password: admin1234) and register server (database.user: docker, database.password: docker, database.dbname: exampledb)

8. Run command:
   ```
   go build && ./messaggio-test
   ```

9. Open an HTTP client(e.g. Thunder Client) to create a message at endpoint: POST http://localhost:8080/v1/messages and then send the following request:
   ```
   ALTER TABLE public.msgs REPLICA IDENTITY FULL;
   ```

10. To set up debezium connector to listen the postgres table changes, run command:
   ```
   curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" 127.0.0.1:8083/connectors/ --data "@debezium.json" 
   ```

11. Open another terminal to tail the kafka topic to see if it's listening to debezium postgres changes. 
    Run here following command:
    ```
    docker run --tty --network messaggio-test_default confluentinc/cp-kafkacat kafkacat -b kafka:9092 -C -s key=s -s value=avro -r http:/schema-registry:8081 -t postgres.public.msgs 
    ```



