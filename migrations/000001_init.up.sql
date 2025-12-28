-- Extensão para gerar UUID (use pgcrypto)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =========================
-- categories
-- =========================
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(64) NOT NULL,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX IF NOT EXISTS idx_categories_user_id
    ON categories(user_id);

-- =========================
-- transactions
-- =========================
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(64) NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    type SMALLINT NOT NULL, -- 0=Expense, 1=Income
    date DATE NOT NULL,
    description TEXT,
    category_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_transactions_category
      FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE INDEX IF NOT EXISTS idx_transactions_user_date
    ON transactions(user_id, date);

CREATE INDEX IF NOT EXISTS idx_transactions_category
    ON transactions(category_id);

-- =========================
-- budgets
-- =========================
CREATE TABLE IF NOT EXISTS budgets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(64) NOT NULL,
    month DATE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT uq_budgets_owner_month UNIQUE (user_id, month)
);

CREATE INDEX IF NOT EXISTS idx_budgets_owner_month
    ON budgets(user_id, month);

-- =========================
-- budget_members
-- (compartilhamento)
-- =========================
CREATE TABLE IF NOT EXISTS budget_members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID NOT NULL,
    user_id VARCHAR(64) NOT NULL,
    role VARCHAR(20) NOT NULL, -- owner/member
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_budget_members_budget
      FOREIGN KEY (budget_id) REFERENCES budgets(id),

    CONSTRAINT uq_budget_members UNIQUE (budget_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_budget_members_budget
    ON budget_members(budget_id);

CREATE INDEX IF NOT EXISTS idx_budget_members_user
    ON budget_members(user_id);

-- =========================
-- budget_items
-- (itens "templates" do planejamento)
-- =========================
CREATE TABLE IF NOT EXISTS budget_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id VARCHAR(64) NOT NULL,            -- quem criou o item
    description TEXT NOT NULL,
    amount NUMERIC(12,2) NOT NULL,
    due_day SMALLINT NOT NULL CHECK (due_day BETWEEN 1 AND 31),
    category_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_budget_items_category
      FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE INDEX IF NOT EXISTS idx_budget_items_user
    ON budget_items(user_id);

-- =========================
-- budget_plannings
-- (liga budget do mês aos items)
-- =========================
CREATE TABLE IF NOT EXISTS budget_plannings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    budget_id UUID NOT NULL,
    budget_item_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_budget_plannings_budget
      FOREIGN KEY (budget_id) REFERENCES budgets(id),

    CONSTRAINT fk_budget_plannings_item
      FOREIGN KEY (budget_item_id) REFERENCES budget_items(id),

    CONSTRAINT uq_budget_plannings UNIQUE (budget_id, budget_item_id)
);

CREATE INDEX IF NOT EXISTS idx_budget_plannings_budget
    ON budget_plannings(budget_id);
