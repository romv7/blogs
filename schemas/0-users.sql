create table if not exists users (
  uuid varbinary(36),
  name varchar(255) unique,
  full_name varchar(255),
  email varchar(255) unique,
  type smallint,
  picture_url varchar(255) null,
  created_at timestamp,
  updated_at timestamp,
  disabled boolean,

  primary key(uuid)
);
