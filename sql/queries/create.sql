-- name: CreateUrl :one
INSERT INTO links(id,createdAt,updatedAt,longUrl,short,accessCount)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    0
)
RETURNING id,createdAt,updatedAt,longUrl,short;