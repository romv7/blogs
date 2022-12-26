create table if not exists comments (
  id int UNSIGNED auto_increment,
  user_id int UNSIGNED,
  comment_text tinytext,
  replies int UNSIGNED,
  edited bool default false,
  created_at timestamp,
  edited_at timestamp null,
  target_type smallint UNSIGNED,
  target_id int UNSIGNED,

  primary key(id)
) auto_increment = 200;
