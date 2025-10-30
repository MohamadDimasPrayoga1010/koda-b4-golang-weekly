CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

INSERT INTO products (name, price) VALUES
('Bangor Pitik Lava', 29000),
('Bangor Pitik Lava Premium', 32000),
('Bangor Cheese Lava', 31000),
('Bangor Lava Sausage', 27500),
('Bangor Jelata Cheese', 24700),
('Bangor Juragan Cheese', 31700),
('Bangor Ningrat Cheese', 49200),
('Bangor Juragan', 29000),
('Bangor Ningrat', 44200),
('Bangor Sultan', 55500),
('Bangor Fish', 27500),
('Hotdog', 19000),
('Chillidong', 19500),
('Cheese Fries', 19400),
('Tea', 9500),
('Soft Drink', 10500),
('Lemon Tea', 15000);