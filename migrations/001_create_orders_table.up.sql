CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(50) NOT NULL
);

-- Inserir alguns dados de exemplo
INSERT INTO orders (user_id, amount, status) VALUES
(1, 100.00, 'pendente'),
(2, 150.50, 'completo'),
(1, 200.00, 'em processamento');