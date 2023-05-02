create table if not exists comments (
  id bigint UNSIGNED,
  uuid varbinary(36),
  user_id bigint UNSIGNED,
  comment_text tinytext,
  replies int UNSIGNED,
  edited boolean default false,
  created_at timestamp,
  edited_at timestamp null,
  target_type smallint UNSIGNED,
  target_uuid varbinary(36),

  primary key(id),
  unique key(uuid)
);
