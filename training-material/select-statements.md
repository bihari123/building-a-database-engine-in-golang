#  SELECT statement

The **SELECT** statement is used to select data from a database.

```
SELECT column1, column2, ...
FROM table_name;
```

If you want to select all the fields available in the table, use the following syntax:
```
SELECT * FROM table_name;
```
# The SQL SELECT DISTINCT Statement
The **SELECT DISTINCT** statement is used to return only distinct (different) values.

Inside a table, a column often contains many duplicate values; and sometimes you only want to list the different (distinct) values.

```
SELECT DISTINCT Syntax
SELECT DISTINCT column1, column2, ...
FROM table_name;
```

The following SQL statement lists the number of different (distinct) customer countries:

```
SELECT COUNT(DISTINCT Country) FROM Customers;

or

SELECT Count(*) AS DistinctCountries
FROM (SELECT DISTINCT Country FROM Customers);
```

# The SQL WHERE Clause
The WHERE clause is used to filter records.

It is used to extract only those records that fulfill a specified condition.

**WHERE** syntax

```
SELECT column1, column2, ...
FROM table_name
WHERE condition;
```

# Text Fields vs. Numeric Fields
SQL requires single quotes around text values (most database systems will also allow double quotes).

However, numeric fields should not be enclosed in quotes:

```
SELECT * FROM Customers
WHERE Country='Mexico';

SELECT * FROM Customers
WHERE CustomerID=1;
```
```
|Operator |	Description	Example |
| ------- | ------------------- |
| =	| Equal |	
| > |	Greater than |	
| <	| Less than |	
| >= |	Greater than or equal |	
| <= | Less than or equal	|
| <> or != |	Not equal |	
| BETWEEN	| Between a certain range |	
| LIKE | Search for a pattern	|
| IN	| To specify multiple possible values for a column |

```

# The SQL AND, OR and NOT Operators
The **WHERE** clause can be combined with **AND**, **OR**, and **NOT** operators.

The **AND** and **OR** operators are used to filter records based on more than one condition:

 - The **AND** operator displays a record if all the conditions separated by **AND** are **TRUE**.
 - The **OR** operator displays a record if any of the conditions separated by **OR** is **TRUE**.
The **NOT** operator displays a record if the condition(s) is **NOT TRUE**.

## AND Syntax
```

SELECT column1, column2, ...
FROM table_name
WHERE condition1 AND condition2 AND condition3 ...;
```

## OR Syntax
```
SELECT column1, column2, ...
FROM table_name
WHERE condition1 OR condition2 OR condition3 ...;
```

## NOT Syntax
```

SELECT column1, column2, ...
FROM table_name
WHERE NOT condition;
```

# Combining AND, OR and NOT
You can also combine the **AND**, **OR** and **NOT** operators.

The following SQL statement selects all fields from "Customers" where country is "Germany" AND city must be "Berlin" OR "München" (use parenthesis to form complex expressions):

### Example
```

SELECT * FROM Customers
WHERE Country='Germany' AND (City='Berlin' OR City='München');

SELECT * FROM Customers
WHERE NOT Country='Germany' AND NOT Country='USA';

```

# The SQL ORDER BY Keyword
The **ORDER BY** keyword is used to sort the result-set in ascending or descending order.

The **ORDER BY** keyword sorts the records in ascending order by default. To sort the records in descending order, use the **DESC** keyword.

### ORDER BY syntax
```
SELECT column1, column2, ...
FROM table_name
ORDER BY column1, column2, ... ASC|DESC;
```

## ORDER BY Several Columns 
The following SQL statement selects all customers from the "Customers" table, sorted ascending by the "Country" and descending by the "CustomerName" column:

###Example
```
SELECT * FROM Customers
ORDER BY Country ASC, CustomerName DESC;
```
