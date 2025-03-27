
INSERT INTO users (name, email, password)
VALUES ('Иван Иванов', 'ivan@example.com', 'hashed_password');


UPDATE users
SET name = 'Иван Петров'
WHERE id = 1;


DELETE FROM users
WHERE id = 1;


INSERT INTO products (name, price)
VALUES ('Продукт 1', 19.99);


UPDATE products
SET price = 24.99
WHERE id = 1;


DELETE FROM products
WHERE id = 1;


INSERT INTO orders (user_id, total_amount)
VALUES (1, 100.50);


INSERT INTO order_products (order_id, product_id, quantity)
VALUES (1, 1, 2);


DELETE FROM orders
WHERE id = 1;


SELECT * FROM users;


SELECT * FROM products;


SELECT * FROM orders
WHERE user_id = 1;


SELECT u.id, u.name, 
       SUM(o.total_amount) AS total_spent,
       AVG(p.price) AS avg_product_price
FROM users u
JOIN orders o ON u.id = o.user_id
JOIN order_products op ON o.id = op.order_id
JOIN products p ON op.product_id = p.id
WHERE u.id = 1
GROUP BY u.id, u.name;


SELECT * FROM users;


SELECT * FROM products;

SELECT * FROM orders
WHERE user_id = 1;

SELECT 
    u.id AS user_id,
    u.name,
    COALESCE(SUM(o.total_amount), 0) AS total_spent,
    COALESCE(AVG(p.price), 0) AS avg_product_price
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
LEFT JOIN order_products op ON o.id = op.order_id
LEFT JOIN products p ON op.product_id = p.id
WHERE u.id = 1
GROUP BY u.id, u.name;


SELECT 
    u.id AS user_id,
    u.name,
    COALESCE(SUM(o.total_amount), 0) AS total_spent,
    COALESCE(AVG(p.price), 0) AS avg_product_price
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
LEFT JOIN order_products op ON o.id = op.order_id
LEFT JOIN products p ON op.product_id = p.id
GROUP BY u.id, u.name;