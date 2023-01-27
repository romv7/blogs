create table if not exists comments (
  id int UNSIGNED auto_increment,
  user_id int UNSIGNED,
  uuid varchar(255) unique,
  comment_text tinytext,
  replies int UNSIGNED,
  edited boolean default false,
  created_at timestamp,
  edited_at timestamp null,
  target_uuid varchar(255),

  primary key(id)
) auto_increment = 200;
