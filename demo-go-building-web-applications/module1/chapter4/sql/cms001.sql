-- phpMyAdmin SQL Dump
-- version 4.8.0.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: 2018-05-18 03:44:31
-- 服务器版本： 8.0.11
-- PHP Version: 7.0.28-0ubuntu0.16.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `cms001`
--

-- --------------------------------------------------------

--
-- 表的结构 `comments`
--

CREATE TABLE `comments` (
  `id` int(11) UNSIGNED NOT NULL,
  `page_id` int(11) NOT NULL,
  `comment_guid` varchar(256) DEFAULT NULL,
  `comment_name` varchar(64) DEFAULT NULL,
  `comment_email` varchar(128) DEFAULT NULL,
  `comment_text` mediumtext,
  `comment_date` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- 表的结构 `pages`
--

CREATE TABLE `pages` (
  `id` int(11) UNSIGNED NOT NULL,
  `page_guid` varchar(256) NOT NULL DEFAULT '',
  `page_title` varchar(256) DEFAULT NULL,
  `page_content` mediumtext,
  `page_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- 转存表中的数据 `pages`
--

INSERT INTO `pages` (`id`, `page_guid`, `page_title`, `page_content`, `page_date`) VALUES
(2, 'hello-world', 'Hello,\r\nWorld', 'I\'m so glad you found this page! It\'s been sitting\r\npatiently on the Internet for some time, just waiting for a\r\nvisitor.', '2018-05-16 09:37:31'),
(3, 'hello-test', 'Hello,\r\nTest', 'I\'m so glad you found this page! \r\nIt\'s been sitting\r\npatiently on the Internet for some time, just waiting for a\r\nvisitor.\r\nHaHa', '2018-05-16 09:37:31'),
(4, 'a-new-blog', 'A New Blog!', 'I\'m so glad you found this page! \r\nIt\'s been sitting\r\npatiently on the Internet for some time, just waiting for a\r\n<i>visitor</i>.\r\nHaHa', '2018-05-17 03:09:18'),
(5, 'two-new-blog', 'Two New Blog!', 'I\'m so glad you found this page! \r\nIt\'s been sitting\r\npatiently on the Internet for some time<i>better</i>, just waiting for a\r\nvisitor.\r\nHaHa\r\nWe can create a new feld within our Page struct and truncate that. But that\'s a little clunky; it requires the feld to always exist within a struct, whether populated with data or not. It\'s much more effcient to expose methods to the template itself.\r\nWe\'ve just scratched the surface of what Go\'s templates can do and we\'ll explore further topics as we continue, but this chapter has hopefully introduced the core\r\nconcepts necessary to start utilizing templates directly.\r\nWe\'ve looked at simple variables, as well as implementing methods within the\r\napplication, within the templates themselves. We\'ve also explored how to bypass\r\ninjection protection for trusted content.\r\nIn the next chapter, we\'ll integrate a backend API for accessing information in a\r\nRESTful way to read and manipulate our underlying data. This will allow us to\r\ndo some more interesting and dynamic things on our templates with Ajax.', '2018-05-17 09:37:31');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `page_id` (`page_id`);

--
-- Indexes for table `pages`
--
ALTER TABLE `pages`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `page_guid` (`page_guid`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `pages`
--
ALTER TABLE `pages`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
