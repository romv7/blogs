create table if not exists uuidIndexes (
  id bigint UNSIGNED,
  resource_id bigint UNSIGNED,
  resource_key varchar(512),
  uuid varbinary(36),
  ref varchar(512),
  pem smallint,
  created_at timestamp,

  primary key(id),
  unique key(resource_key)
)