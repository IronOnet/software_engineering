## SQL QUERIES 

### The SELECT statement 

=> The SELECT clause determines which of all possible columns should be included in the query's reqult set

=> You can include in your SELECT clause things as: 
	- Literals, such as numbers or strings 
	- Expressions, such as transaction.amount * -1 
	- Built-in function calls, such as ROUND(transaction.amount, 2) 
	- User-defined function calls 

=> The next query demonstrates the use of a table column, a literal, an expression, and a built-in function 
call in an single query against the employee table 

```sql 

SELECT language_id, 
	'COMMON' language_usage, 
	language_id * 3.1415 lang_pi_value, 
	upper(name) language_name FROM language; 
``` 


### COLUMN ALIASES 

```sql  

SELECT language_id, 'COMMON' language_usage, 
language_id * 3.14 lang_pi_value, 
upper(name) language_name 
FROM language; 
```

In order to make your column aliases stand out even more, you also have the option of using the as 
keyword before the alias name, as in : 

```sql 

SELECT language_id, 
	'COMMON' AS language_usage, 
	language_id * 3.14 AS lang_pi_value, 
	upper(name) AS language_name 
FROM language; 


### FROM clause 

=> The from clause defines the tables used by a query, along with the means of linking the tables together 

#### Tables 

A table is a set of related rows, four different types of tables meet this relaxed definition 

	- Permannent tables (i.e created using the create table statement) 
	- Derived tables (i.e, rows returned by a subquery and held in memory) 
	- Virtual tables (i.e, created using the create view statement) 


##### Derived (subquery-generated) tables 


```sql 

SELECT concat(cust.last_name, ', ', cust.firs_name) full_name 
FROM 
	(SELECT first_name, last_name, email FROM customer WHERE first_name = 'JESSIE') cust;

```

##### Temporary tables 

Data stored in temporary tables will disappear at one point (generally at the end of a transaction or when your database session is closed). 

```sql 

CREATE TEMPORARY TABLE actors_j 
	(actor_id smallint(5), 
	first_name varchar(45), 
	last_name varchar(45) 
);

```

##### Views 

A view is a query that is stored in the data dictionary. It looks and acts like a table, but there is no data 
associate with a view (this is why it's called a virtual table). When you issue a query against a view, your query 
is merged with the view definition to create a final query to be executed.

```sql 

CREATE VIEW cust_vw AS 
SELECT customer_id, first_name, last_name, active 
FROM customer; 

SELECT first_name, last_name 
	FROM cust_vw 
	WHERE active = 0; 

```

##### Table Links 

=> The second deviation from the simple from clause definition is the mandate if more than one table appears 
in the from clause, the conditions used to link the tables must be included as well. this is not a 
requirement  of MYSQL

```sql 

SELECT customer.first_name, customer.last_name, 
	time(rental.rental_data) rental_time 
	FROM customer 
	INNER JOIN rental 
	ON customer.customer_id = rental.customer_d 
	WHERE date(rental.rental_date) = '2005-06-14';
``` 

The previous query displays data from both the customer table (first_name, last_name) and the rental table (rental_date), 
so both tables are included in the from clause. The mechanism for linking the two tables is the customer ID stored 
in both the customer and rental tables


### The "WHERE" clause 

=> THe "WHERE" clause is a mechanism for filtering out unwanted rows from the result set. 

```sql 

FROM film 
WHERE rating = 'G' AND rental_duration >= 7;

``` 

The previous filter contains, two filter conditions, but we can include as many as we want. 

```sql 

SELECT title 
FROM film 
WHERE rating = 'G' OR rental_duran >= 7;


### The "GROUP BY" and "HAVING" clauses 

The "GROUP BY" clause is used to group data by column values. For example

```sql 

SELECT c.first_name c.last_name, count(*) 
FROM customer AS c 
INNER JOIN rental AS r 
ON c.customer_id = r.customer_id 
GROUP BY c.first_name, c.last_name 
HAVING count(*) >= 40;

### The Order By Clause 

=> In general, the rows in a result set returned from a query are not in any particualr order. 
If you want your result set to be sorted, you will need to instruct the server to sort 
the result using the "ORDER BY" clause. 

```sql 
SELECT c.first_name, c.last_name, 
timer(r.rental_date) rental_time 
FROM customer AS c
INNER JOIN rental AS r 
ON c.customer_id = r.customer_id 
WHERE date(r.rental_date) = '2005-06-14'; 

SELECT c.first_name, c.last_name FROM customer AS c 
INNER JOIN rental AS r 
ON c.customer_id = r.customer_id 
WHERE date(r.rental_ate) = '2005-06-14' 
ORDER BY c.last_name; 
```


#### Sorting via numeric placeholders 

=> In the following example, we use an ORDER BY clause specifying a 
descending sort using the third element in the select clause

```sql

SELECT c.first_name, c.last_name, 
time(r.rental_date) as rental_time 
FROM customer as c 
INNER JOIN rental as r 
ON c.customer_id = r.customer_d 
WHERE date(r.rental_date) = '2005-06-14' 
ORDER BY 3 desc;

