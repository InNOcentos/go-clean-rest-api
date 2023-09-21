CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name varchar(255) not null
);
