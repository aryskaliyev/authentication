-- forum.sql

-- PRAGMA foreign_keys = ON;

-- Create a table to store user accounts
CREATE TABLE IF NOT EXISTS useraccount (
	user_id INTEGER NOT NULL,
	username VARCHAR(30) NOT NULL,
	email VARCHAR(255) NOT NULL,
	hashed_password BLOB NOT NULL,
	created DATETIME NOT NULL,
	UNIQUE (username),
	UNIQUE (email),
	PRIMARY KEY (user_id)
);

-- Create a table to store sessions
CREATE TABLE IF NOT EXISTS session (
	user_id INTEGER NOT NULL,
	uuid_token TEXT NOT NULL,
	created DATETIME NOT NULL,
	expires DATETIME NOT NULL,
	UNIQUE (user_id),
	FOREIGN KEY (user_id)
		REFERENCES useraccount(user_id)
		ON DELETE CASCADE,
	PRIMARY KEY (user_id)
);

-- Create a table to store categories
CREATE TABLE IF NOT EXISTS category (
	category_id INTEGER NOT NULL,
	name VARCHAR(20) NOT NULL,
	created DATETIME NOT NULL,
	UNIQUE (name),
	PRIMARY KEY (category_id)	
);

-- Create a table to store posts
CREATE TABLE IF NOT EXISTS post (
	post_id INTEGER NOT NULL,
	--user_id INTEGER NOT NULL,
	title VARCHAR(75) NOT NULL,
	body VARCHAR(500) NOT NULL,
	created DATETIME NOT NULL,
	--FOREIGN KEY (user_id)
	--	REFERENCES useraccount(user_id)
	--	ON DELETE CASCADE,
	PRIMARY KEY (post_id)
);

-- Create a table to store post-votes
CREATE TABLE IF NOT EXISTS post_vote (
	post_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	vote INTEGER NOT NULL CHECK (vote == 1 OR vote == -1),
	FOREIGN KEY (user_id)
		REFERENCES useraccount(user_id)
		ON DELETE CASCADE,
	FOREIGN KEY (post_id)
		REFERENCES post(post_id)
		ON DELETE CASCADE,
	PRIMARY KEY (post_id, user_id)
);

-- Create a table to store post-categories
CREATE TABLE IF NOT EXISTS post_category (
	post_id	INTEGER NOT NULL,
	category_id INTEGER NOT NULL,
	FOREIGN KEY (post_id)
		REFERENCES post(post_id)
		ON DELETE CASCADE,
	FOREIGN KEY (category_id)
		REFERENCES category(category_id)
		ON DELETE CASCADE,
	PRIMARY KEY (post_id, category_id)
);

-- Create a table to store comments
CREATE TABLE IF NOT EXISTS comment (
	comment_id INTEGER NOT NULL,
	body VARCHAR(75) NOT NULL,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	created DATETIME NOT NULL,
	FOREIGN KEY (user_id)
		REFERENCES useraccount(user_id)
		ON DELETE CASCADE,
	FOREIGN KEY (post_id)
		REFERENCES post(post_id)
		ON DELETE CASCADE,
	PRIMARY KEY (comment_id)
);

-- Create a table to store comment-votes
CREATE TABLE IF NOT EXISTS comment_vote (
	comment_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	vote INTEGER NOT NULL CHECK (vote == 1 OR vote == -1),
	FOREIGN KEY (user_id)
		REFERENCES useraccount(user_id)
		ON DELETE CASCADE,
	FOREIGN KEY (comment_id)
		REFERENCES comment(comment_id)
		ON DELETE CASCADE,
	PRIMARY KEY (comment_id, user_id)
);
