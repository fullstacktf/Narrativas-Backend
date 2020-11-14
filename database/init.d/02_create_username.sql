CREATE USER IF NOT EXISTS 'rollify'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON rollify.* TO 'rollify'@'%';