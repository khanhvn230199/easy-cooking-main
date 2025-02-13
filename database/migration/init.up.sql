CREATE TABLE recipes
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255)             NOT NULL,
    description TEXT,
    cuisine     VARCHAR(100),
    photo_url   TEXT,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE ingredients
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255)             NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);

CREATE TABLE recipe_ingredients
(
    id            SERIAL PRIMARY KEY,
    recipe_id     INTEGER                  NOT NULL,
    ingredient_id INTEGER                  NOT NULL,
    quantity      VARCHAR(100),
    unit          VARCHAR(50),
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP WITH TIME ZONE NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes (id) ON DELETE CASCADE,
    FOREIGN KEY (ingredient_id) REFERENCES ingredients (id) ON DELETE CASCADE
);

CREATE TABLE instructions
(
    id          SERIAL PRIMARY KEY,
    recipe_id   INTEGER                  NOT NULL,
    step_number INTEGER                  NOT NULL,
    description TEXT                     NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMP WITH TIME ZONE NULL,
    FOREIGN KEY (recipe_id) REFERENCES recipes (id) ON DELETE CASCADE
);

-- Trigger để tự động cập nhật updated_at
CREATE OR REPLACE FUNCTION update_modified_column()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_recipes_modtime
    BEFORE UPDATE
    ON recipes
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_ingredients_modtime
    BEFORE UPDATE
    ON ingredients
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_recipe_ingredients_modtime
    BEFORE UPDATE
    ON recipe_ingredients
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();

CREATE TRIGGER update_instructions_modtime
    BEFORE UPDATE
    ON instructions
    FOR EACH ROW
EXECUTE FUNCTION update_modified_column();
