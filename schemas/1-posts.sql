create table if not exists posts (
  id int UNSIGNED auto_increment,
  headline_text tinytext unique,
  subject_text varchar(255),
  user_id int UNSIGNED,

  keywords json null,
  content_url varchar(255),
  blog_url varchar(255) unique,

  multimedia_url json null,

  stage smallint UNSIGNED,
  status smallint UNSIGNED,

  primary key(id)
) auto_increment = 300;
