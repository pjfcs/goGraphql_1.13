create table users(
      id bigserial primary key
    , username varchar(55) unique not null
    , email varchar(255) unique not null
);