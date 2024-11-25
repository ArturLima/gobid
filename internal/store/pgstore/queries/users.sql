-- name: CreateUser :one

insert into users("user_name", "email", "password_hash", "bio")
values($1, $2, $3, $4)
  RETURNING id;

-- name: GetUserById :one
SELECT
      id,
      user_name,
      password_hash,
      email,
      bio,
      created_at,
      updated_at
FROM users
where id = $1;

-- name: GetUserByEmail :one
SELECT
      id,
      user_name,
      password_hash,
      email,
      bio,
      created_at,
      updated_at
FROM users
where email = $1;