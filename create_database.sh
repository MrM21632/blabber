# Uncomment if the database also doesn't exist
# createdb -U blabber blabber
psql -U blabber -d blabber -a -f ./blabber.sql
