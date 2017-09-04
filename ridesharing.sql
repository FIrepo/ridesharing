-- phpMyAdmin SQL Dump
-- version 4.2.11
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Sep 04, 2017 at 09:39 AM
-- Server version: 5.6.21
-- PHP Version: 5.6.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `ridesharing`
--

-- --------------------------------------------------------

--
-- Table structure for table `driver`
--

CREATE TABLE IF NOT EXISTS `driver` (
`id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(70) NOT NULL,
  `is_visible` tinyint(1) NOT NULL DEFAULT '0',
  `is_connected` tinyint(1) NOT NULL DEFAULT '0',
  `lat` double NOT NULL,
  `lon` double NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `driver`
--

INSERT INTO `driver` (`id`, `name`, `email`, `password`, `is_visible`, `is_connected`, `lat`, `lon`) VALUES
(1, 'degananda', 'degananda.ferdian@gmail.com', '$2a$10$lK6YRpBByAtLw/2VLQ2WduJYEN3PQkgk1IE0JqCcSQNRDoH4SJOWq', 1, 1, 2.54, 5.85);

-- --------------------------------------------------------

--
-- Table structure for table `passenger`
--

CREATE TABLE IF NOT EXISTS `passenger` (
`id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(70) NOT NULL,
  `is_visible` tinyint(1) NOT NULL DEFAULT '0',
  `is_connected` tinyint(1) NOT NULL DEFAULT '0',
  `lat` varchar(50) NOT NULL,
  `lon` varchar(50) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `passenger`
--

INSERT INTO `passenger` (`id`, `name`, `email`, `password`, `is_visible`, `is_connected`, `lat`, `lon`) VALUES
(1, 'ferdian', 'ferdian.degananda@gmail.com', '$2a$10$lK6YRpBByAtLw/2VLQ2WduJYEN3PQkgk1IE0JqCcSQNRDoH4SJOWq', 1, 1, '', '');

-- --------------------------------------------------------

--
-- Table structure for table `request`
--

CREATE TABLE IF NOT EXISTS `request` (
`id` int(11) NOT NULL,
  `id_passenger` int(11) NOT NULL,
  `id_driver` int(11) DEFAULT NULL,
  `lat` varchar(50) NOT NULL,
  `lon` varchar(50) NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  `distance` double DEFAULT NULL,
  `time` varchar(25) DEFAULT NULL
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `request`
--

INSERT INTO `request` (`id`, `id_passenger`, `id_driver`, `lat`, `lon`, `status`, `distance`, `time`) VALUES
(10, 1, 1, '1.4789', '2.555', 4, 13.5, '250'),
(11, 1, 1, '1222', '15555', 1, NULL, NULL),
(12, 1, 1, '2.54', '3.52', 4, 150.8, '12580'),
(13, 1, 1, '2.54', '3.52', 0, NULL, NULL),
(14, 1, 1, '2.54', '3.52', 0, NULL, NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `driver`
--
ALTER TABLE `driver`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `passenger`
--
ALTER TABLE `passenger`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `request`
--
ALTER TABLE `request`
 ADD PRIMARY KEY (`id`), ADD KEY `id_passenger` (`id_passenger`), ADD KEY `id_driver` (`id_driver`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `driver`
--
ALTER TABLE `driver`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `passenger`
--
ALTER TABLE `passenger`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `request`
--
ALTER TABLE `request`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=15;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `request`
--
ALTER TABLE `request`
ADD CONSTRAINT `request_ibfk_1` FOREIGN KEY (`id_passenger`) REFERENCES `passenger` (`id`),
ADD CONSTRAINT `request_ibfk_2` FOREIGN KEY (`id_driver`) REFERENCES `driver` (`id`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
