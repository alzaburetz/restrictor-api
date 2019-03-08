CREATE TABLE usergroup (
id integer not null primary key autoincrement,
group_name text )

CREATE TABLE users (
id integer primary key,
userkey text unique,
username text,
group_id int, foreign key (group_id) references usergroup(id) )


CREATE TABLE "restriction" (
id integer primary key autoincrement,
app text,
rule text,
time text,
hours_from int,
hours_to int,
executable text,
user_gr int references usergroup )
