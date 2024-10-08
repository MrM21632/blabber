# Uncomment if the database and user also don't exist
# createuser --interactive blabber
# createdb -U blabber blabber

# Create database schema and tables
psql -U blabber -d blabber -a -f ./blabber.sql

# Populate tables
psql -U blabber -d blabber -a -f ./blabber_users.sql
psql -U blabber -d blabber -a -f ./blabber_posts.sql
