CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;

create table users (
     user_id  serial not null PRIMARY KEY,
     nickname citext UNIQUE ,
     email citext UNIQUE ,
     fullname text,
     about text default ''
);

create table forums(
    forum_id    serial not null PRIMARY KEY,
    slug          citext UNIQUE ,
    threads       int default 0,
    posts         int default 0,
    title         text,
    user_nickname    citext
);

create table threads(
   thread_id serial not null PRIMARY KEY,
   forum_id int,
   title text,
   author citext,
   forum  citext,
   message  text,
   votes int default 0,
   slug  citext default '',
   created TIMESTAMPTZ  DEFAULT NOW()
);

    CREATE UNIQUE INDEX nul_uni_idx ON threads(slug)
        WHERE slug not in ('');

create table votes(
    vote_id serial PRIMARY KEY,
    voice int,
    nickname  citext ,
    thread_id int ,
    UNIQUE (thread_id,nickname)
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
    parent int default NULL,
    parents  bigint[] default array []::INTEGER[],
);

CREATE INDEX forum_slug_idx ON forums (slug);
CREATE INDEX users_nick_idx ON users (nickname);

CREATE INDEX fpi_idx ON posts ((posts.parents[1]), thread_id);
CREATE INDEX pid_idx ON posts ((posts.parents[1]), post_id);
CREATE INDEX parents_idx ON posts ((posts.parents[1]));
CREATE INDEX thread_idx ON posts (thread_id);
CREATE INDEX pare_idx ON posts ((posts.parent));

create index if not exists slug_id on threads (slug);
create index if not exists f_created_idx on threads (forum, created);
create index if not exists t_author_idx on threads (author, forum);

CREATE INDEX on votes(nickname,thread_id);

CREATE OR REPLACE FUNCTION parents_change() RETURNS TRIGGER AS
$parents_change$
DECLARE
    temp_arr     BIGINT[];
    first_parent posts;
BEGIN
    IF (NEW.parent = 0) THEN
        NEW.parents := array_append(new.parents, new.post_id::bigint);
    ELSE
        SELECT parents FROM posts WHERE post_id = new.parent INTO temp_arr;
        SELECT * FROM posts WHERE post_id = temp_arr[1] INTO first_parent;
        IF NOT FOUND OR first_parent.thread_id != NEW.thread_id THEN
            RAISE EXCEPTION 'bad parent' USING ERRCODE = '00409';
        end if;
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


CREATE TRIGGER tri_update_forum after insert
    on threads FOR EACH ROW
EXECUTE PROCEDURE updateforum();


CREATE TABLE forum_users
(
    nickname citext NOT NULL,
    forum    citext NOT NULL,
    UNIQUE (nickname,forum)
);

CREATE OR REPLACE FUNCTION update_Forum() RETURNS TRIGGER AS
$update_users_forum$
BEGIN
    INSERT INTO forum_users (nickname, forum) VALUES (NEW.author, NEW.forum) on conflict do nothing;
    return NEW;
end
$update_users_forum$ LANGUAGE plpgsql;



CREATE TRIGGER thread_insert_user_forum
    AFTER INSERT
    ON threads
    FOR EACH ROW
EXECUTE PROCEDURE update_Forum();

CREATE TRIGGER post_insert_user_forum
    AFTER INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE update_Forum();


CREATE INDEX forum_user_index ON forum_users (forum, lower(nickname));


