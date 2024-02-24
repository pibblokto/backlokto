# backlokto

Running test postgres in Docker:

```bash
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=backlokto -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=backlokto -d postgres
```

Connecting to test postgres:

```bash
docker exec -it postgres psql -U backlokto -d backlokto
```

Create table for test data:

```sql
CREATE TABLE backlokto_data (
  column1 VARCHAR(255),
  column2 VARCHAR(255),
  column3 INT
); 
```

Insert test data:

```sql
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('FIRST_ROW_1', 'FIRST_ROW_2', 10);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('SECOND_ROW_1', 'SECOND_ROW_2', 15);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('THIRD_ROW_1', 'THIRD_ROW_2', 20);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('FOURTH_ROW_1', 'FOURTH_ROW_2', 25);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('FIFTH_ROW_1', 'FIFTH_ROW_2', 30);
INSERT INTO backlokto_data (column1, column2, column3) VALUES ('SIXTH_ROW_1', 'SIXTH_ROW_2', 35); 
```