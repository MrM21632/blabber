# Populate tables
psql -U blabber -d blabber -a -f ./blabber_4_users_testdata.sql
psql -U blabber -d blabber -a -f ./blabber_5_posts_testdata.sql
