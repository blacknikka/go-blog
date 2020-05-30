create table articles (
  id serial not null,
  title varchar(256),
  body text,
  created_at date,
  updated_at date,
  deleted_at date
);
