create table if not exists comments (
  uuid varbinary(36),
  user_uuid varbinary(36),
  comment_text tinytext,
  replies int UNSIGNED,
  edited boolean default false,
  created_at timestamp,
  edited_at timestamp null,
  target_uuid varbinary(36),

  primary key(uuid)
);
