echo "Creating tables on DB urls 🚀"

set -e

if [-z "$DB_URL"]; then
    echo "No DB_URL found"
    echo 1
fi

goose -dir sql/schema postgres "$DB_URL" down

echo "Tables created"