echo "Creating tables on DB urls 🚀"

chmod +x create_db.sh

set -e

if [-z "$DB_URL"]; then
    echo "No DB_URL found"
    echo 1
fi



goose -dir sql/schema postgres "$DB_URL" up

echo "✅ Tables created"