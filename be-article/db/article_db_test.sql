-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Sep 23, 2024 at 07:39 PM
-- Server version: 5.7.39
-- PHP Version: 7.4.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `article_db_test`
--

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `id` int(11) NOT NULL,
  `author_id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`id`, `author_id`, `title`, `content`, `image_url`, `created_at`, `updated_at`, `deleted_at`) VALUES
(4, 1, 'new article from dion', '<p>whats up chat</p>', 'https://www.thesaurus.com/e/wp-content/uploads/2021/11/20211104_articles_1000x562-790x310.png', '2024-09-23 21:51:57', '2024-09-24 00:11:03', '2024-09-23 17:11:03'),
(5, 1, 'w article', '<p>spam w in the chat</p>', 'https://www.thesaurus.com/e/wp-content/uploads/2021/11/20211104_articles_1000x562-790x310.png', '2024-09-23 22:13:21', '2024-09-24 00:10:37', NULL),
(6, 1, 'W article in the chat', '<p>whats up chat</p>', 'https://www.thesaurus.com/e/wp-content/uploads/2021/11/20211104_articles_1000x562-790x310.png', '2024-09-23 22:13:51', '2024-09-23 22:21:49', '2024-09-23 15:21:49'),
(7, 1, 'L article', '<p>SPAM L in the chat</p>', 'https://www.thesaurus.com/e/wp-content/uploads/2021/11/20211104_articles_1000x562-790x310.png', '2024-09-24 00:14:26', '2024-09-24 00:14:26', NULL),
(8, 2, 'W Dion Pamungkas', '<p>Spam 2 in the article</p>', 'https://www.thesaurus.com/e/wp-content/uploads/2021/11/20211104_articles_1000x562-790x310.png', '2024-09-24 00:37:23', '2024-09-24 00:37:30', NULL),
(9, 2, 'W Example', '<p>dadas</p>', 'https://res.cloudinary.com/dh9h41hiw/image/upload/v1727114943/test_cloudinary/cykzvwcbkz42zwoqekmg.jpg', '2024-09-24 01:03:36', '2024-09-24 01:09:10', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` enum('visitor','author') NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `reset_token` varchar(255) DEFAULT NULL,
  `reset_token_expires` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `role`, `created_at`, `updated_at`, `reset_token`, `reset_token_expires`) VALUES
(1, 'Dion Arya', 'dionarya.p@gmail.com', '$2a$10$m2xecOAUtDQ5a8vG2zRehuvcXgeVwewT4UJtHobMjEmcJnD5tGtW6', 'author', '2024-09-23 19:49:06', '2024-09-24 02:36:22', NULL, NULL),
(2, 'W speed', 'speedmantap@gmail.com', '$2a$10$s35CSPraMUyLWKRqZpG.0.HFRfZGzIiW9W24z1VoOCw5mzb1UvI7q', 'author', '2024-09-24 00:35:10', '2024-09-24 00:35:10', NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `author_id` (`author_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `articles`
--
ALTER TABLE `articles`
  ADD CONSTRAINT `articles_ibfk_1` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
