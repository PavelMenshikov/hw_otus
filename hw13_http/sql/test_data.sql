-- Очистка таблиц (если требуется, чтобы избежать дублирования при повторном запуске)
TRUNCATE TABLE order_products RESTART IDENTITY CASCADE;
TRUNCATE TABLE orders RESTART IDENTITY CASCADE;
TRUNCATE TABLE products RESTART IDENTITY CASCADE;
TRUNCATE TABLE users RESTART IDENTITY CASCADE;

-- Вставка пользователей
INSERT INTO users (name, email, password)
VALUES 
  ('Иван Иванов', 'ivan@example.com', 'hashed_password1'),
  ('Мария Петрова', 'maria@example.com', 'hashed_password2'),
  ('Андрей Сидоров', 'andrey@example.com', 'hashed_password3');

-- Вставка товаров (обязательно убедись, что файл сохранён в UTF-8)
INSERT INTO products (name, price)
VALUES 
  ('Ноутбук', 75000.00),
  ('Смартфон', 50000.00),
  ('Наушники', 5000.00);

-- Вставка заказов; здесь user_id=1 и user_id=2 должны существовать
INSERT INTO orders (user_id, total_amount)
VALUES 
  (1, 100000.00),
  (2, 55000.00);

-- Вставка записей в связующую таблицу заказов и товаров;
-- Здесь order_id=1 и order_id=2 должны существовать, а также product_id в диапазоне вставленных товаров.
INSERT INTO order_products (order_id, product_id, quantity)
VALUES 
  (1, 1, 1),
  (1, 3, 2),
  (2, 2, 1);
