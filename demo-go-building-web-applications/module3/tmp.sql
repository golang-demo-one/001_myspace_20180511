UPDATE mysql.user SET Password = OLD_PASSWORD('ycliu912')  WHERE Host = 'localhost' AND User = 'root';
update mysql.user set authentication_string=password('ycliu912')  where user='root';

INSERT INTO comments set `comment_name`='a',`comment_email`='a@qq.com',`comment_text`='hhh';
ALTER TABLE `comments` CHANGE `page_id` `id` INT(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE
    comments ADD CONSTRAINT `page_id` FOREIGN KEY(page_guid) REFERENCES pages(page_guid)
	
ALTER TABLE  comments ADD CONSTRAINT `page_id` FOREIGN KEY(page_id) REFERENCES pages(page_id)	

INSERT INTO comments SET comment_name='liu', comment_email='liu@qq.com', comment_text='liu', comment_guid='hello-world'

alter table comments modify id int(11) unsigned NOT NULL AUTO_INCREMENT

522:
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'ycliu912';
mysql> ALTER USER 'testuser0912'@'%' IDENTIFIED WITH mysql_native_password BY 'ycliu912';
Query OK, 0 rows affected (0.07 sec)