# Create the user and database for blabber
# Comment these lines out if the user and database already exist
createuser --interactive blabber
createdb -U blabber blabber

# Create database schema and tables
psql -U blabber -d blabber -a -f ./blabber_1_db.sql
psql -U blabber -d blabber -a -f ./blabber_2_users.sql
psql -U blabber -d blabber -a -f ./blabber_3_posts.sql
