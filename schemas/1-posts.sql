create table if not exists posts (
  id int UNSIGNED auto_increment,
  headline_text tinytext unique,
  subject_text varchar(255),
  user_id int UNSIGNED,
  keywords json null,
  content_url varchar(255),
  multimedia_url json null,
  stage smallint UNSIGNED,
  status smallint UNSIGNED,
  
  revised_at timestamp,
  archived_at timestamp,
  published_at timestamp,
  created_at timestamp,

  uuid varchar(255) unique,
  
  primary key(id)
) auto_increment = 300;
