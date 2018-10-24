CREATE TABLE accounts (
    id varchar(24) NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    verified tinyint(1) NOT NULL,
    created_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_time timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB;
