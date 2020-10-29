create table users (
     user_id  serial not null PRIMARY KEY,
     nickname VARCHAR (50) UNIQUE,
     email VARCHAR (50) UNIQUE,
     fullname text,
     about text default ''
);

create table forums(
    forum_id    serial not null PRIMARY KEY,
    slug          text UNIQUE ,
    threads       int default 0,
    posts         int default 0,
    title         text,
    user_nickname     text
);
