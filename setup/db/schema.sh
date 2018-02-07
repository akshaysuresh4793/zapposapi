mysqld_safe --skip-grant-tables &
sleep 5
MYSQL_PWD=password mysql -u root -e "source /tmp/schema.sql"
MYSQL_PWD=password mysql -u root -e "SHOW DATABASES"
