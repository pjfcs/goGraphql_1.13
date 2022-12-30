create table meetups(
      id bigserial primary key
    , name varchar(255) unique not null
    , description text
    , user_id bigserial references users(id) on delete cascade not null
);