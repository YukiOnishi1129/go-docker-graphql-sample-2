CREATE TABLE IF NOT EXISTS 20220328_GO_GRAPHQL_DB.todos
(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(50) NOT NULL,
    comment TEXT NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES 20220328_GO_GRAPHQL_DB.users(id) ON UPDATE RESTRICT ON DELETE CASCADE
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

# ALTER TABLE 20220328_GO_GRAPHQL_DB.todos
#     ADD CONSTRAINT fk_users FOREIGN KEY (user_id) REFERENCES 20220328_GO_GRAPHQL_DB.users(id) ON UPDATE RESTRICT ON DELETE CASCADE;