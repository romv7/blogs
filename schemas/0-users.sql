create table if not exists users (
  id bigint UNSIGNED,
  uuid varbinary(36),
  name varchar(255) unique,
  full_name varchar(255),
  email varchar(255) unique,
  type smallint,
  created_at timestamp,
  updated_at timestamp,
  disabled boolean,

  primary key(id),
  unique key(uuid)
);
