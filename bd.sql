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
   slug  citext default '',
   created TIMESTAMPTZ  DEFAULT NOW(),
   CONSTRAINT fk_threads
       FOREIGN KEY(forum_id)
           REFERENCES forums(forum_id)
           ON DELETE CASCADE
);

CREATE UNIQUE INDEX nul_uni_idx ON threads(slug)
    WHERE slug not in ('');

create table votes(
    vote_id serial not null PRIMARY KEY,
    voice int,
    nickname  citext  not null,
    thread_id int not null,
    UNIQUE (thread_id,nickname),
    CONSTRAINT fk_votes
        FOREIGN KEY (thread_id)
            REFERENCES threads(thread_id)
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
    parent int default null,
    parent_path LTREE,
     CONSTRAINT fk_posts
        FOREIGN KEY(thread_id)
            REFERENCES threads(thread_id)
            ON DELETE CASCADE,
    CONSTRAINT fk_forums
            FOREIGN KEY(forum_id)
            REFERENCES forums(forum_id)
            ON DELETE CASCADE
);
CREATE EXTENSION IF NOT EXISTS ltree with schema posts;

CREATE INDEX section_parent_path_idx ON posts USING GIST (parent_path);
CREATE INDEX posts_parent_id_idx ON posts (parent);

CREATE OR REPLACE FUNCTION get_calculated_si_node_path(param_si_id integer)
    RETURNS ltree AS
$$
SELECT  CASE WHEN s.parent = 0 THEN (0::text || '.' || s.post_id::text)::ltree
             ELSE get_calculated_si_node_path(s.parent)  || s.post_id::text  END
FROM posts As s
WHERE s.post_id = $1;
$$
    LANGUAGE sql;

CREATE OR REPLACE FUNCTION trig_update_si_node_path() RETURNS trigger AS
$$
BEGIN
    IF TG_OP = 'INSERT' THEN
        UPDATE posts SET parent_path = get_calculated_si_node_path(NEW.post_id) WHERE posts.post_id = NEW.post_id;
    END IF;
    RETURN NEW;
END
$$
    LANGUAGE 'plpgsql' VOLATILE;

CREATE TRIGGER trig01_update_si_node_path AFTER INSERT OR UPDATE
    ON posts FOR EACH ROW
EXECUTE PROCEDURE trig_update_si_node_path();


