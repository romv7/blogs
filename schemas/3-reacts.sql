create table if not exists reacts (
  user_id int UNSIGNED,
  type smallint UNSIGNED,
  target_uuid varchar(255),
  created_at timestamp
);
