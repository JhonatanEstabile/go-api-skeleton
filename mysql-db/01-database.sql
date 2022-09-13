# create databases
CREATE DATABASE IF NOT EXISTS `api`;

# create root user and grant rights
GRANT ALL ON *.* TO 'root'@'%';