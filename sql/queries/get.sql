-- name: GetUrl :one
SELECT id,createdAt,updatedAt,longUrl,short FROM links WHERE short = $1;
