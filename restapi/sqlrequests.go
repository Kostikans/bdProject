package restapi

const AddUserRequest = "INSERT INTO users(user_id,nickname,email,fullname,about) VALUES(default,$1,$2,$3,$4)"

const GetPreviousUsers = "SELECT nickname,fullname,email,about FROM users WHERE nickname=$1 OR email=$2"

const GetUserRequest = "SELECT nickname,fullname,email,about FROM users WHERE nickname=$1"

const UpdateUserRequest = "UPDATE users SET fullname=$2,email=$3,about=$4 WHERE nickname=$1 and ($3 not in('') or $2 not in('') or $4 not in(''))"

const CheckUserExist = "Select nickname from users WHERE nickname=$1"

const PostForumRequest = "INSERT INTO forums(forum_id,slug,title,user_nickname) VALUES(default,$1,$2,$3)"

const GetForumInfoRequest = "SELECT title,user_nickname ,slug,posts,threads FROM forums WHERE slug=$1"

const GetForumUsersRequest = "SELECT nickname, fullname, about, email FROM "

const CheckForumExist = "SELECT slug,forum_id from forums WHERE slug=$1"

const CreateThreadRequest = "INSERT INTO threads(thread_id,title,author,message,created,id,forum_id,forum,slug)" +
	" VALUES(default,$1,$2,$3,$4,$5,$6,$7,$8)"

const GetExistThreadReuqest = "SELECT id,title,author,message,votes,slug,created FROM threads WHERE title=$1"

const GetExistThreadToPostReuqest = "SELECT t.thread_id,f.forum_id,f.slug FROM threads as t INNER JOIN forums as f on t.forum_id=f.forum_id WHERE t.slug=$1 OR t.thread_id=$2"

const GetThreadsFromForum = "SELECT id,t.title,author,message,votes,t.slug,created FROM forums as f INNER JOIN threads as t on t.forum_id=f.forum_id WHERE f.slug=$1 limit $2"

const CreatePostRequest = "INSERT INTO posts(post_id,parent,author,message,forum,thread_id,forum_id,created) VALUES(default,$1,$2,$3,$4,$5,$6,$7)"
