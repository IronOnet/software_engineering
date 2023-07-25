/** Queries for the Sakila Database */ 

SELECT first_name, last_name 
FROM customer 
WHERE last_name = 'ZIEGLER'; 

/** Return every row from the category table */ 
SELECT * FROM category; 

/** Query Clauses */ 
/*
	Several components or clauses make up the select statement. While only one of them is 
	mandatory when using MySQL (the select clause), you will usually include at least 
	two or three of the six available clauses. 

	select ==> Determines which columns to include in the query's result set 
	from ===> identifies the tables from which to retrieve data and how the tables should be joined 
	where ===> Filters out unwanted data 
	groupby ==> Used to group rows together by common column values 
	having ==> filters out unwanted groups 
	order by ===> Sorts the rows of the final result set by one or more columns.

 */


/** The SELECT clause */ 
SELECT * FROM language; 

SELECT name FROM language; 

SELECT language_id, 'COMMON' language_usage, language_id * 3.1415297 lang_pi_value, 
upper(name) language_name FROM langue; 


