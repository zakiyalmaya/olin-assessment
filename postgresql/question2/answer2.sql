-- Use postgres_fdw (Foreign Data Wrapper) to connect to the second database
-- this is can be used for querying across multiple databases
-- https://www.postgresql.org/docs/current/foreign-data.html

-- Instal the extention
CREATE EXTENSION postgres_fdw;

-- Create the server connection
CREATE SERVER fdw_server FOREIGN DATA WRAPPER postgres_fdw OPTIONS(host 'localhost', dbname 'second', port '5432');

-- Create the user mapping
CREATE USER MAPPING FOR postgres SERVER fdw_server OPTIONS(user 'postgres', password 'postgre');

-- Import the data
IMPORT FOREIGN SCHEMA public FROM SERVER fdw_server INTO public;

-- Query the data
SELECT u.name,
	o.amount,
	o.created_at
FROM users u 
JOIN orders o ON u.id = o.user_id;