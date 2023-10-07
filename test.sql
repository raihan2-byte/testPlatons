-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 07, 2023 at 04:30 PM
-- Server version: 10.4.24-MariaDB
-- PHP Version: 8.1.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `platon`
--

-- --------------------------------------------------------

--
-- Table structure for table `platon`
--

CREATE TABLE `platon` (
  `ID` int(11) NOT NULL,
  `OriginLatitude` varchar(255) DEFAULT NULL,
  `OriginLongitude` varchar(255) DEFAULT NULL,
  `OriginPostalCode` varchar(255) DEFAULT NULL,
  `Items` text DEFAULT NULL,
  `Couriers` int(11) DEFAULT NULL,
  `DestinationLatitude` varchar(255) DEFAULT NULL,
  `DestinationLongitude` varchar(255) DEFAULT NULL,
  `PostalCode` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `platon`
--

INSERT INTO `platon` (`ID`, `OriginLatitude`, `OriginLongitude`, `OriginPostalCode`, `Items`, `Couriers`, `DestinationLatitude`, `DestinationLongitude`, `PostalCode`) VALUES
(1, '123.456', '789.012', '12345', 'Example Items', 2, '789.012', '123.456', 54321),
(3, '123.456', '789.012', '12345', 'Example Items', 2, '789.012', '123.456', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `platon`
--
ALTER TABLE `platon`
  ADD PRIMARY KEY (`ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `platon`
--
ALTER TABLE `platon`
  MODIFY `ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
