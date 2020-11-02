create table users (
     user_id  serial not null PRIMARY KEY,
     nickname citext UNIQUE ,
     email citext UNIQUE ,
     fullname text,
     about text default ''
);
CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;
CREATE UNIQUE INDEX nickname_unique_idx on users (nickname);
CREATE UNIQUE INDEX email_unique_idx on users (email);

create table forums(
    forum_id    serial not null PRIMARY KEY,
    slug          citext UNIQUE ,
    threads       int default 0,
    posts         int default 0,
    title         text,
    user_nickname    citext,
    CONSTRAINT fk_threads
        FOREIGN KEY(user_nickname)
        REFERENCES users(nickname)
        ON DELETE CASCADE
);

create table threads(
   thread_id serial not null PRIMARY KEY,
   id int default 42,
   forum_id int,
   title text UNIQUE ,
   author citext,
   forum  citext,
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
        FOREIGN KEY (nickname)
            REFERENCES users(nickname)
            ON DELETE CASCADE
);

create table posts(
    post_id serial not null PRIMARY KEY,
    thread_id int,
    forum_id int ,
    author citext,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    forum citext,
    is_edited bool DEFAULT false,
    message text,
    parent int ,
    parents  bigint[] default array []::INTEGER[],
     CONSTRAINT fk_posts
        FOREIGN KEY(thread_id)
            REFERENCES threads(thread_id)
            ON DELETE CASCADE
);

CREATE INDEX section_parent_path_idx ON posts USING GIST (parents);
CREATE INDEX posts_parent_id_idx ON posts (parent);

CREATE OR REPLACE FUNCTION parents_change() RETURNS TRIGGER AS
$parents_change$
DECLARE
    temp_arr     BIGINT[];
    first_parent posts;
BEGIN
    IF (NEW.parent IS NULL) THEN
        NEW.parents := array_append(new.parents, new.post_id);
    ELSE
        SELECT parents FROM posts WHERE post_id = new.parent INTO temp_arr;
        SELECT * FROM posts WHERE post_id = temp_arr[1] INTO first_parent;

        NEW.parents := NEW.parents || temp_arr || new.post_id::bigint;
    end if;
    RETURN new;
end
$parents_change$ LANGUAGE plpgsql;


CREATE TRIGGER trig01_update_si_node_path before insert
    ON posts FOR EACH ROW
EXECUTE PROCEDURE parents_change();


CREATE OR REPLACE FUNCTION updateForum() RETURNS TRIGGER AS
$updateForum$
BEGIN
    Update forums SET threads=threads+1 WHERE  forum_id=new.forum_id;
    RETURN new;
end
$updateForum$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION updateForum1() RETURNS TRIGGER AS
$updateForum1$
BEGIN
    Update forums SET posts=posts+1 WHERE  forum_id=new.forum_id;
    RETURN new;
end
$updateForum1$ LANGUAGE plpgsql;

CREATE TRIGGER tri_update_forum after insert
    on threads FOR EACH ROW
EXECUTE PROCEDURE updateforum();

CREATE TRIGGER tri1_update_forum after insert
    on posts FOR EACH ROW
EXECUTE PROCEDURE updateforum1();


