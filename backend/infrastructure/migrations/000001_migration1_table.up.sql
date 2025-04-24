-- Membuat tabel roles
CREATE TABLE roles (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Membuat tabel users
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role_id INT NOT NULL DEFAULT 2,
    email VARCHAR(100),  -- Menambahkan kolom email
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Menambahkan kolom created_at
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

-- Membuat tabel consumers
CREATE TABLE consumers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,  -- Menggunakan user_id, bukan consumer_id
    nik VARCHAR(16) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    legal_name VARCHAR(100),
    birth_place VARCHAR(50),
    birth_date DATE,
    salary DECIMAL(15, 2),
    id_card_image_url VARCHAR(255),
    selfie_image_url VARCHAR(255),
    FOREIGN KEY (user_id) REFERENCES users(id)  -- Relasi dengan users.id
);

-- Membuat tabel limits
CREATE TABLE limits (
    consumer_id INT NOT NULL,
    limit_1 INT NOT NULL,
    limit_2 INT NOT NULL,
    limit_3 INT NOT NULL,
    limit_6 INT NOT NULL,
    FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);

-- Membuat tabel transactions
CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    consumer_id INT NOT NULL,
    contract_number VARCHAR(50) NOT NULL,
    otr DECIMAL(15, 2) NOT NULL,
    admin_fee DECIMAL(15, 2) NOT NULL,
    installment DECIMAL(15, 2) NOT NULL,
    interest DECIMAL(15, 2) NOT NULL,
    asset_name VARCHAR(100),
    transaction_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);

-- Menambahkan role default
INSERT INTO roles (name) VALUES 
('System administrators'),
('All employees');

-- Menambahkan admin user
-- (gunakan password hash bcrypt yang valid)
INSERT INTO users (username, email, password, role_id, created_at)
VALUES 
('admin', 'admin@example.com', '$2a$10$UBBhm00/TYiXWcVjButJROW0MhzPtwEhiITDotkCdsAQkgDFP5xmq', 1, CURRENT_TIMESTAMP);
