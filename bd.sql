create table users (
     user_id  serial not null PRIMARY KEY,
     nickname citext UNIQUE ,
     email citext UNIQUE ,
     fullname text,
     about text default ''
);
CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA users;
CREATE UNIQUE INDEX nickname_unique_idx on users (nickname);
CREATE UNIQUE INDEX email_unique_idx on users (email);

create table forums(
    forum_id    serial not null PRIMARY KEY,
    slug          citext UNIQUE ,
    threads       int default 0,
    posts         int default 0,
    title         text,
    user_nickname     text
);

create table threads(
   thread_id serial not null PRIMARY KEY,
   id int default 42,
   forum_id int,
   title text UNIQUE ,
   author text,
   forum  text,
   message  text,
   votes int default 0,
   slug  text default '',
   created TIMESTAMPTZ  DEFAULT NOW(),
   CONSTRAINT fk_threads
       FOREIGN KEY(forum_id)
           REFERENCES forums(forum_id)
           ON DELETE CASCADE
);

create table posts(
    post_id serial not null PRIMARY KEY,
    thread_id int,
    forum_id int ,
    author text,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    forum text,
    is_edited bool DEFAULT false,
    message text,
    parent int DEFAULT 0,
     CONSTRAINT fk_posts
        FOREIGN KEY(thread_id)
            REFERENCES threads(thread_id)
            ON DELETE CASCADE,
    CONSTRAINT fk_forums
            FOREIGN KEY(forum_id)
            REFERENCES forums(forum_id)
            ON DELETE CASCADE
);