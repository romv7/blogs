create table if not exists reacts (
  user_id int UNSIGNED,
  type smallint UNSIGNED,
  target_uuid varbinary(36),
  created_at timestamp
);
