-- If Not exists
CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  username VARCHAR (50) UNIQUE NOT NULL,
  firstname VARCHAR (20) NOT NULL,
  lastname VARCHAR (20) NOT NULL,
  email VARCHAR (300) UNIQUE NOT NULL,
  contact VARCHAR UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT now,
  deleted_at TIMESTAMP DEFAULT now,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- If EXISTS
ALTER TABLE users
ADD COLUMN contact VARCHAR (255) UNIQUE NOT NULL;
ALTER TABLE users
ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT now();
ALTER TABLE users
ADD COLUMN deleted_at TIMESTAMP NOT NULL DEFAULT now();
ALTER TABLE users
ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;