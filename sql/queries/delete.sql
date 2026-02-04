-- name: DeleteUrl :exec
DELETE FROM links WHERE short = $1;