DROP DATABASE IF EXISTS react_todo_app;
CREATE DATABASE react_todo_app;
USE react_todo_app;

---- drop ----
DROP TABLE IF EXISTS  tasks;

---- create ----
create table tasks
(
 id               INT(20) AUTO_INCREMENT,
 title            VARCHAR(255) NOT NULL,
 body             VARCHAR(255) NOT NULL,
 create_at        DATETIME NOT NULL,
 update_at        DATETIME NOT NULL,
 is_delete        BOOLEAN NOT NULL DEFAULT false,
  PRIMARY KEY (id)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

use react_todo_app;
insert into tasks values(1,'title test','title body',now(),now(),false);