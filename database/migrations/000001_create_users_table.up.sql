CREATE TABLE IF NOT EXISTS users(
   id serial PRIMARY KEY,
   username VARCHAR (50) UNIQUE NOT NULL,
   firstname VARCHAR (20) NOT NULL,
   lastname VARCHAR (20) NOT NULL,
   email VARCHAR (300) UNIQUE NOT NULL
);