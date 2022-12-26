create table if not exists reacts (
  user_id int UNSIGNED,
  type smallint UNSIGNED,
  target_type smallint UNSIGNED,
  target_id int UNSIGNED
);
