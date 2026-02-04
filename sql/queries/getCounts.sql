-- name: GetCounts :one
SELECT * FROM links WHERE short = $1;
