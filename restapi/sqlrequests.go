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
	" VALUES(default,$1,$2,$3,$4,$5,$6,$7,$8) RETURNING thread_id"

const GetExistThreadReuqest = "SELECT thread_id,title,author,message,votes,slug,created,forum FROM threads WHERE slug=$1"

const GetExistThreadToPostReuqest = "SELECT t.thread_id,f.forum_id,f.slug FROM threads as t INNER JOIN forums as f on t.forum_id=f.forum_id WHERE t.slug=$1 OR t.thread_id=$2"

const GetThreadsFromForum = "SELECT id,t.title,author,message,votes,t.slug,created FROM forums as f INNER JOIN threads as t on t.forum_id=f.forum_id WHERE f.slug=$1 limit $2"

const GetExistThreadReuqestToVote = "SELECT thread_id,title,author,message,votes,forum,slug,created FROM threads WHERE slug=$1 OR thread_id=$2"

const CreatePostRequest = "INSERT INTO posts(post_id,parent,author,message,forum,thread_id,forum_id,created,parent_path) VALUES(default,$1,$2,$3,$4,$5,$6,$7,$8) RETURNING post_id"

const GetPrevVote = "SELECT voice FROM votes where nickname=$1 and thread_id=$2"
const AddVote = "INSERT INTO votes(vote_id,voice,nickname,thread_id) VALUES(default,$1,$2,$3)"

const UpdateVote = "Update votes set voice=$1 where nickname=$2 and thread_id=$3"

const UpdateVoteCount = "Update threads Set votes=(CASE WHEN $1 > 0 Then votes + 1 ELSE votes - 1 END) WHERE thread_id=$2 RETURNING votes"
