-- Create categories table
CREATE TABLE finance.categories (
    id SERIAL PRIMARY KEY,
    -- ID incremental como chave primária
    user_id INTEGER,
    -- ID do usuário associado, não pode ser nulo
    name VARCHAR(50) NOT NULL,
    icon VARCHAR(50) NOT NULL,
    color VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Add index
CREATE INDEX categories_name_idx ON finance.categories(name);

CREATE INDEX categories_user_id_idx ON finance.categories(user_id);

CREATE INDEX categories_created_at_idx ON finance.categories(created_at);

CREATE INDEX categories_updated_at_idx ON finance.categories(updated_at);

-- Insert default categories
INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Alimentação', 'icon-food', '#FF5733');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (
        null,
        'Compras e Lazer',
        'icon-shopping',
        '#33FF57'
    );

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Moradia', 'icon-home', '#3357FF');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (
        null,
        'Saúde e Bem-estar',
        'icon-health',
        '#FF33A1'
    );

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (
        null,
        'Investimento',
        'icon-investment',
        '#FF8C33'
    );

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (
        null,
        'Transferências',
        'icon-transfer',
        '#33FFF5'
    );

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Empréstimos', 'icon-loan', '#FF3333');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Renda', 'icon-income', '#33FF8C');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Imposto e Taxas', 'icon-tax', '#FF33FF');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Educação', 'icon-education', '#8C33FF');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Transporte', 'icon-transport', '#FF5733');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Streaming', 'icon-streaming', '#33FF57');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (
        null,
        'Assinaturas',
        'icon-subscription',
        '#3357FF'
    );

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Viagem', 'icon-travel', '#FF33A1');

INSERT INTO
    finance.categories (user_id, name, icon, color)
VALUES
    (null, 'Seguros', 'icon-insurance', '#FF8C33');