-- name: UpdateUrl :one
UPDATE links
SET
    longUrl = $1
    updatedAt = NOW()
WHERE short = $2
RETURNING id,createdAt,updatedAt,longUrl,short;

