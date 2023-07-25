#### CHARACTER SET

If the value in the fourth column, maxlen is greater than 1, then the
ccharacter set is a multibyte character set

In prior versions of the MYSQL server, the latin1 character set was automatically
chosen as the default chatracter set, but version 8 defaults to utf8mb4

==> When using a floating point type, you can specify a precision (the total number
  of allowable digits both to the left and to the right of the decimal point) and
  a scale (the number of allowable digits to the right of the decimal point)
  but they are not required.

#### Temporal Data

Along with strings and numbers, you will almost certainely be working with
information about dates and /or times. This type of data is referred to as
temporal, and some examples of temporal data in a database include:

  - The future date, that a particular event is expected to happen, such as a
  shipping a customer's order

  - The date that a customer's order was shipped
  - The date and time that a user modified a particular row in a table
  - The date and time that a user modified a particular row in a table
  - An employees birth date
  - The year corresponding to a row in a yearly_sales fact table in a data warehouse
  - The elapsed time needed to complete a wring harness on an automobile assembly line.


### Table Creation

#### Design

A good way to start designing a table is to do a bit of brainstorming to see what
kind of information would be helpful to include.

#### Refinment

Normalization is the process of ensureing that there are no duplicate (other than foreign keys)
or compoound columns in your database design.

###### Table Person

  - person_id smallint(unsigned)
  - first_name varchar(20)
  - last_name varchar(20)
  - eye_color char(2)
  - birth_date date
  - street varchar(30)
  - city varchar(20)
  - state varchar(20)
  - country varchar(20)
  - postal_code varchar(20)

#### Building SQL Schema statments

Create TABLE person
  (person_id SMALLINT UNSIGNED,
    fname VARCHAR(20),
    lname VARCHAR(20),
    eye_color CHAR(2),
    birth_date DATE,
    street VARCHAR(30),
    city VARCHAR(20),
    country VARCHAR(20),
    postal_code VARCHAR(20),
    CONSTRAINT pk_person PRIMARY KEY (person_id));


The last line of the staement is a constraint that defines the primary key of the table.
Another type of constraint called a check constraint  constrains the allowable values for
a particular column. MYSQL allows a check constraint to be attached to a column definition
as in the following:

eye_color CHAR(2) CHECK (eye_color IN ('BR', 'BL', 'GR')),

While check constraints operate as expected on most database servers, the MYSQL
server allows check constraints to be defined but does not enforce them.
However MYSQL does provide another character data type called enum that merges
the check constraint into the data type definition.


```sql

eye_color ENUM('BR', 'BL', 'GR'),
```
Here's how the person table definition looks with an enum data type for the eye_color
column:

```sql
CREATE TABLE person
  (person_id SMALLINT UNSIGNED,
    fname VARCHAR(20),
    lname VARCHAR(20),
    eye_color ENUM('BR', 'BL', 'GR')),
    birth_date DATE,
    street VARCHAR(30),
    city VARCHAR(20),
    state VARCHAR(20),
    country VARCHAR(20),
    postal_code VARCHAR(20),
    CONSTRAINT pk_person PRIMARY KEY (person_id)
);
```

```sql

CREATE TABLE favorite_food
  (person_id SMALLINT UNSIGNED,
  food VARCHAR(20),
  CONSTRAINT pk_favorite_food PRIMARY KEY (person_id, food),
  CONSTRAINT fk_fav_food_person_id FOREIGN KEY (person_id)
  REFERENCES person (person_id)
);
```

Since a person can have more than one favorite food (which is the reason this table
  was created in the first place), it takes more than just the person_id column
  to guarantee uniqueness in the table. This table, therefore has two primary key: person_id and food

- The favorite_food table contains another type of constraint which is called a favorite_food

### Populating and Modifying Tables

#### Inserting Data

Since there is not yet any data in the person and favorite_food tables, the first of the
four SQL data statements to be explored will be the insert statment. There are three
main components to an insert statment.

  - The name of the table into which to add the data
  - The names of the columns in the table to be populated
  - The values with which to populate the columns

### Generating numeric key data

Before inserting data into the person table, it would be useful to discuss how values
are generated for numeric primary keys. Other than picking a number out of thin air
you have a couple of options.

```sql

ALTER TABLE person MODIFY person_id SMALLINT UNSIGNED AUTO_INCREMENT;
```

If you are running these statemtns in your database, you will first need to disable
the foreign key constraint on the favorite_foocd table, and then re-enable the constraints
when finished. T

```sql

SET foreign_key_checks=0;
ALTER TABLE person
  MODIFY person_id SMALLINT UNSIGNED AUTO_INCREMENT;
SET foreign_key_checks=1;
```

```sql

INSERT INTO favorite_food (person_id, food)
-> VALUES (1, 'pizza');
Query OK, 1 row affected (0.01 sec)
mysql> INSERT INTO favorite_food (person_id, food)
-> VALUES (1, 'cookies');
Query OK, 1 row affected (0.00 sec)
mysql> INSERT INTO favorite_food (person_id, food)
-> VALUES (1, 'nachos');

```

Here's a query that retrieves William's favorite foods in alphabetical order using
an order by clause:

```sql

SELECT food FROM favorite_food
WHERE person_id = 1
ORDER BY food;
```

We can execute anothyer statment insert to add SUsan Smith to the person table

```sql

INSERT INTO person
(person_id, fname, lname, eye_color, birth_date, street, city, state, country, postal_code)
VALUES (null, 'Susan', 'Smith', 'BL', '1975-11-02', '23 Mapple St.', 'Arlington', 'VA', 'USA', '20220');
```

#### Deleting Data

```sql

DELETE FROM person WHERE person_id = 2;
```

### When Good Statments Go Bad

There are lots of wasy that you can run afoul when inserting or modifying data.
This section shows you some of the common mistakes that you might come accross
and how the MYSQL server will respond.

#### Nonunique Primary Key

Because the table definitions include the creation of primary key constraints,
MYSQL will make sure that duplicate keys values are not inserted into the tables.
The next statement attempts to bypass the auto-increment feature of the person_id column
and create another row in the person table with the person_id of 1

```sql 

INSERT INTO person (person_id, fname, lname eye_color, birth_date)
VALUES (1, 'Charles', 'Fulton', 'GR', '1968-01-15');

```


### NOnexistant Foreign Key

The table definition of the favorite_food table includes the creation of a
foreign key constraint on the person_id column. This constraint ensures that
all values of person_id entered into the favorite_food table exist in the per
