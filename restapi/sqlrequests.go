package restapi

const AddUserRequest = "INSERT INTO users VALUES(default,$1,$2,$3,$4)"

const GetPreviousUsers = "SELECT nickname,fullname,email,about FROM users WHERE nickname=$1 OR email=$2"

const GetUserRequest = "SELECT nickname,fullname,email,about FROM users WHERE nickname=$1"

const UpdateUserRequest = "UPDATE users SET fullname=$2,email=$3,about=$4 WHERE nickname=$1"

const PostForumRequest = "INSERT INTO forums(forum_id,slug,title,user_nickname) VALUES(default,$1,$2,$3)"

const GetForumInfoRequest = "SELECT title,user_nickname ,slug,posts,threads FROM forums WHERE slug=$1"

const GetForumUsersRequest = "SELECT nickname, fullname, about, email FROM "
