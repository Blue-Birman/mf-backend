-- phpMyAdmin SQL Dump
-- version 5.0.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 28, 2021 at 08:32 PM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.2.34

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `mf_backend`
--

-- --------------------------------------------------------

--
-- Table structure for table `carts`
--

CREATE TABLE `carts` (
  `customer_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `carts`
--

INSERT INTO `carts` (`customer_id`, `product_id`, `created_at`, `updated_at`) VALUES
(6, 1, '2021-03-19', '2021-03-19');

-- --------------------------------------------------------

--
-- Table structure for table `categories`
--

CREATE TABLE `categories` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `categories`
--

INSERT INTO `categories` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Men', '2020-12-08', '2020-12-08'),
(2, 'Women', '2020-12-08', '2020-12-08');

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `address` varchar(255) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `name`, `email`, `password`, `address`, `created_at`, `updated_at`) VALUES
(1, 'Alessandro', 'rvalessandro@gmail.com', 'rv123', 'Bali no 2', '2020-12-08', '2020-12-08'),
(2, 'Michelle', 'michele@gmail.com', 'michele123', 'Cianjur no 3', '2020-12-08', '2020-12-08'),
(3, 'Daniel Alexander', 'daniela@gmail.com', 'daniel123', 'Bandung no 4', '2020-12-08', '2020-12-08'),
(4, 'Daniel Christianto', 'christianto@gmail.com', 'dc123', 'Kopo no 4', '2020-12-08', '2020-12-08'),
(5, 'Hanjaya', 'hanjaya@gmail.com', 'hanjaya123', 'Bandung no 2', '2020-12-08', '2020-12-08'),
(6, 'Unit Testing', 'unit@testing.com', '123', 'Jl Dipatiukur', '2021-03-19', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `image_url` varchar(255) NOT NULL,
  `preview_url` varchar(255) DEFAULT NULL,
  `price` varchar(255) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `category_id`, `name`, `description`, `image_url`, `preview_url`, `price`, `created_at`, `updated_at`) VALUES
(1, 1, 'White Clothes', 'White clothes with elastic fiber, appropiate for hard sporting, will absorb sweat fast', 'assets/product/set_1.png', '36805', '1000000', '2020-12-08', '2020-12-08'),
(2, 1, 'Blue Clothes', 'blue clothes with elastic fiber, appropiate for hard sporting, will absorb sweat fast', 'assets/product/set_2.png', '38266', '1500000', '2020-12-08', '2020-12-08'),
(3, 1, 'Blue Mask', 'Blue mask with thick fiber, will always block dust and virus from outside in', 'assets/product/set_3.png', '38317', '1250000', '2020-12-08', '2020-12-08'),
(4, 1, 'Purple Mask', 'Purple mask with thick fiber, will always block dust and virus from outside in', 'assets/product/set_4.png', '37283', '2000000', '2020-12-08', '2020-12-08'),
(5, 2, 'White fiber cloth', 'White clothes with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm', 'assets/product/set_5.png', '37682', '2500000', '2020-12-08', '2020-12-08'),
(6, 2, 'Blue fiber cloth', 'blue clothes with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm', 'assets/product/set_6.png', '38100', '2500000', '2020-12-08', '2020-12-08'),
(7, 2, 'Red White fiber cloth', 'Pattern clothes with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm', 'assets/product/set_7.png', '38063', '2500000', '2020-12-08', '2020-12-08'),
(8, 2, 'Blue Tiger fiber cloth', 'Tiger Pattern clothes with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm', 'assets/product/set_8.png', '38085', '2500000', '2020-12-08', '2020-12-08'),
(9, 2, 'Red Jacket', 'Red Jacket with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm', 'assets/product/set_9.png', '38687', '4500000', '2020-12-08', '2020-12-08');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL,
  `date` date NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `customer_id`, `date`, `created_at`, `updated_at`) VALUES
(6, 1, '2020-12-09', '2020-12-09', '2020-12-09'),
(7, 1, '2020-12-09', '2020-12-09', '2020-12-09'),
(8, 1, '2020-12-09', '2020-12-09', '2020-12-09'),
(9, 1, '2020-12-09', '2020-12-09', '2020-12-09'),
(10, 1, '2020-12-10', '2020-12-10', '2020-12-10'),
(11, 1, '2021-02-22', '2021-02-22', '2021-02-22');

-- --------------------------------------------------------

--
-- Table structure for table `transaction_products`
--

CREATE TABLE `transaction_products` (
  `id` int(11) NOT NULL,
  `transaction_id` int(11) DEFAULT NULL,
  `product_id` int(11) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `created_at` date NOT NULL,
  `updated_at` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transaction_products`
--

INSERT INTO `transaction_products` (`id`, `transaction_id`, `product_id`, `qty`, `created_at`, `updated_at`) VALUES
(10, 6, 1, 2, '2020-12-09', '2020-12-09'),
(11, 6, 3, 1, '2020-12-09', '2020-12-09'),
(12, 7, 1, 2, '2020-12-09', '2020-12-09'),
(13, 7, 3, 1, '2020-12-09', '2020-12-09'),
(14, 8, 9, 1, '2020-12-09', '2020-12-09'),
(15, 9, 6, 2, '2020-12-09', '2020-12-09'),
(16, 10, 3, 1, '2020-12-10', '2020-12-10'),
(17, 11, 1, 4, '2021-02-22', '2021-02-22');

-- --------------------------------------------------------

--
-- Table structure for table `wishlists`
--

CREATE TABLE `wishlists` (
  `customer_id` int(11) NOT NULL,
  `product_id` int(11) NOT NULL,
  `created_at` date NOT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `wishlists`
--

INSERT INTO `wishlists` (`customer_id`, `product_id`, `created_at`, `updated_at`) VALUES
(1, 4, '2021-03-02', '2021-03-02');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `carts`
--
ALTER TABLE `carts`
  ADD PRIMARY KEY (`customer_id`,`product_id`),
  ADD KEY `carts_product_id_fk` (`product_id`);

--
-- Indexes for table `categories`
--
ALTER TABLE `categories`
  ADD UNIQUE KEY `category_id_uindex` (`id`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `products_category_id_fk` (`category_id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transaction_products`
--
ALTER TABLE `transaction_products`
  ADD UNIQUE KEY `transaction_products_id_uindex` (`id`),
  ADD KEY `transaction_products_products_id_fk` (`product_id`),
  ADD KEY `transaction_products_transactions_id_fk` (`transaction_id`);

--
-- Indexes for table `wishlists`
--
ALTER TABLE `wishlists`
  ADD PRIMARY KEY (`customer_id`,`product_id`),
  ADD KEY `wishlists_product_id_fk` (`product_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `transaction_products`
--
ALTER TABLE `transaction_products`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `carts`
--
ALTER TABLE `carts`
  ADD CONSTRAINT `carts_customer_id_fk` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`),
  ADD CONSTRAINT `carts_product_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_category_id_fk` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);

--
-- Constraints for table `transaction_products`
--
ALTER TABLE `transaction_products`
  ADD CONSTRAINT `transaction_products_products_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  ADD CONSTRAINT `transaction_products_transactions_id_fk` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);

--
-- Constraints for table `wishlists`
--
ALTER TABLE `wishlists`
  ADD CONSTRAINT `wishlists_customer_id_fk` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`),
  ADD CONSTRAINT `wishlists_product_id_fk` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
