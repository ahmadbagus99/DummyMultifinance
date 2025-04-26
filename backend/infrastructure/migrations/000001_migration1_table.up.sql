CREATE DATABASE IF NOT EXISTS dummy_multifinance_dev;
USE dummy_multifinance_dev;

-- Roles table
CREATE TABLE IF NOT EXISTS roles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- tenors table
CREATE TABLE IF NOT EXISTS tenors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL DEFAULT 2,
    email VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- Consumer table
CREATE TABLE IF NOT EXISTS consumers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    nik VARCHAR(50) NOT NULL UNIQUE,
    full_name VARCHAR(100) NOT NULL,
    legal_name VARCHAR(100),
    birth_place VARCHAR(100),
    birth_date DATE,
    salary DECIMAL(15,2),
    ktp_photo TEXT,
    selfie_photo TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Limits table
CREATE TABLE IF NOT EXISTS limits (
    consumer_id INT NOT NULL,
    tenor_id INT NOT NULL,
    Amount INT NOT NULL,
    FOREIGN KEY (tenor_id) REFERENCES tenors(id),
    FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);

-- Transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,        
    contract_number VARCHAR(255) NOT NULL,      
    consumer_id INT NOT NULL,                  
    otr DECIMAL(15, 2) NOT NULL,               
    admin_fee DECIMAL(15, 2) NOT NULL,         
    installment DECIMAL(15, 2) NOT NULL,       
    interest DECIMAL(15, 2) NOT NULL,          
    asset_name VARCHAR(255) NOT NULL,          
    transaction_date DATETIME NOT NULL,     
    approved BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);

-- Menambahkan role default
INSERT INTO roles (name) VALUES 
('System administrators'),
('All employees');

-- Menambahkan tenor default
INSERT INTO tenors (name) VALUES 
('Tenor 1'),
('Tenor 2'),
('Tenor 3'),
('Tenor 6');

-- Menambahkan admin user
-- (gunakan password hash bcrypt yang valid)
INSERT INTO users (username, email, password, role_id, created_at)
VALUES 
('admin', 'admin@example.com', '$2a$10$UBBhm00/TYiXWcVjButJROW0MhzPtwEhiITDotkCdsAQkgDFP5xmq', 1, CURRENT_TIMESTAMP);

