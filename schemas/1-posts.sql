create table if not exists posts (
  id bigint UNSIGNED,
  user_id int UNSIGNED,
  
  uuid varbinary(36),
  headline_text tinytext unique,
  summary_text varchar(255),
  tags varchar(255) null,
  stage smallint UNSIGNED,
  status smallint UNSIGNED,
  
  revised_at timestamp NULL,
  archived_at timestamp NULL,
  published_at timestamp NULL,
  created_at timestamp,

  original_id bigint UNSIGNED NULL,

  primary key(id),
  unique key(uuid)
);
