create table if not exists users (
  id int UNSIGNED auto_increment,
  name varchar(255) unique,
  full_name varchar(255),
  email varchar(255) unique,
  type smallint,
  picture_url varchar(255) null,
  created_at timestamp,

  primary key(id)
) auto_increment = 100;
