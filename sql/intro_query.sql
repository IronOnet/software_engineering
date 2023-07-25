/* This is a comment*/
SELECT t.txn_id, t.txn_type_cd, t.txn_date, t.amount 
FROM individual i 
    INNER JOIN account a ON i.cust_id = a.cust_id 
    INNER JOIN product p ON p.product_cd = a.product_cd 
    INNER JOIN transaction t ON t.account_id = a.account_id 
WHERE i.fname = 'George' AND i.lname = 'Blake' AND p.name = 'checking account';


SELECT cust_id, fname 
FROM individual 
WHERE lname = 'Smith';

SELECT t.txn_id, t.txn_type_cd, t.txn_date, t.amount 
FROM account a 
    INNER JOIN transaction t ON t.account_id = a.account_id 
WHERE a.cust_id = 8 AND a.product_cd = 'CHK';

/* Along with querying your database, you will most likely be involved with 
populating and modifying the data in your database. Here's a simple example 
of how you would insert a new row into the product table$*/

INSERT INTO product (product_cd, name) 
VALUES ('CD', 'Certificate of Depysit') 

/*A misspel happened, we can update the row with the UPDATE stateemtn*/ 

UPDATE product 
SET name = 'Certificate Of Deposit' 
WHERE product_id = 'CD'; 


/*Create Person Query*/ 
CREATE TABLE person
(person_id SMALLINT UNSIGNED, 
fname VARCHAR(20), 
lname VARCHAR(20), 
eye_color ENUM('BR', 'BL', 'GR'), 
birth_date DATE, 
street VARCHAR(30), 
city VARCHAR(20), 
state VARCHAR(20), 
country VARCHAR(20), 
postal_code VARCHAR(20), 
CONSTRAINT pk_person PRIMARY KEY (person_id) 
);


/** Create a table of favorite food */ 
CREATE TABLE favorite_food 
( person_id SMALLINT UNSIGNED, 
food VARCHAR(20), 
CONSTRAINT pk_favorite_food PRIMARY KEY (person_id, food), 
CONSTRAINT fk_fav_food_person_id FOREIGN KEY (person_id) 
REFERENCES person (person_id)); 


/* Delete a row */ 
DELETE FROM person WHERE person_id = 2; 






