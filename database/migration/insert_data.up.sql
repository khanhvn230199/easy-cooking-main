-- Thêm dữ liệu vào bảng ingredients
INSERT INTO ingredients (name, created_at, updated_at)
VALUES ('Egg', NOW(), NOW()),
       ('Flour', NOW(), NOW()),
       ('Milk', NOW(), NOW()),
       ('Butter', NOW(), NOW()),
       ('Salt', NOW(), NOW()),
       ('Chicken', NOW(), NOW()),
       ('Rice', NOW(), NOW()),
       ('Tomato', NOW(), NOW()),
       ('Onion', NOW(), NOW()),
       ('Garlic', NOW(), NOW()),
       ('Beef', NOW(), NOW()),
       ('Pasta', NOW(), NOW()),
       ('Cheese', NOW(), NOW()),
       ('Olive Oil', NOW(), NOW()),
       ('Pepper', NOW(), NOW());

-- Thêm dữ liệu vào bảng recipes
INSERT INTO recipes (title, description, cuisine, photo_url, created_at, updated_at)
VALUES ('Pancakes', 'Delicious fluffy pancakes', 'American', 'http://example.com/pancakes.jpg', NOW(), NOW()),
       ('Scrambled Eggs', 'Soft and creamy scrambled eggs', 'American', 'http://example.com/scrambled_eggs.jpg', NOW(),
        NOW()),
       ('Chicken Fried Rice', 'Classic Asian stir-fried rice with chicken', 'Asian',
        'http://example.com/chicken_fried_rice.jpg', NOW(), NOW()),
       ('Beef Stir Fry', 'Spicy beef stir-fry with vegetables', 'Asian', 'http://example.com/beef_stir_fry.jpg', NOW(),
        NOW()),
       ('Pasta Carbonara', 'Creamy Italian pasta with eggs and cheese', 'Italian',
        'http://example.com/pasta_carbonara.jpg', NOW(), NOW()),
       ('Tomato Soup', 'Homemade creamy tomato soup', 'European', 'http://example.com/tomato_soup.jpg', NOW(), NOW()),
       ('Grilled Chicken Salad', 'Fresh salad with grilled chicken', 'Mediterranean',
        'http://example.com/chicken_salad.jpg', NOW(), NOW()),
       ('Beef Burger', 'Classic American beef burger', 'American', 'http://example.com/beef_burger.jpg', NOW(), NOW()),
       ('Vegetable Stir Fry', 'Healthy mixed vegetable stir fry', 'Asian', 'http://example.com/vegetable_stir_fry.jpg',
        NOW(), NOW()),
       ('Cheese Omelette', 'Fluffy omelette with melted cheese', 'French', 'http://example.com/cheese_omelette.jpg',
        NOW(), NOW()),
       ('Spaghetti Bolognese', 'Classic Italian pasta with meat sauce', 'Italian',
        'http://example.com/spaghetti_bolognese.jpg', NOW(), NOW()),
       ('Chicken Curry', 'Spicy Indian chicken curry', 'Indian', 'http://example.com/chicken_curry.jpg', NOW(), NOW());

-- Thêm mối quan hệ giữa recipes và ingredients
WITH ingredient_map AS (SELECT name, id
                        FROM ingredients),
     recipe_map AS (SELECT title, id
                    FROM recipes)
INSERT INTO recipe_ingredients (recipe_id, ingredient_id, quantity, unit)
VALUES
    -- Pancakes
    ((SELECT id FROM recipe_map WHERE title = 'Pancakes'), (SELECT id FROM ingredient_map WHERE name = 'Egg'), '2',
     'pieces'),
    ((SELECT id FROM recipe_map WHERE title = 'Pancakes'), (SELECT id FROM ingredient_map WHERE name = 'Flour'), '1',
     'cup'),
    ((SELECT id FROM recipe_map WHERE title = 'Pancakes'), (SELECT id FROM ingredient_map WHERE name = 'Milk'), '1',
     'cup'),
    ((SELECT id FROM recipe_map WHERE title = 'Pancakes'), (SELECT id FROM ingredient_map WHERE name = 'Butter'), '2',
     'tablespoons'),

    -- Scrambled Eggs
    ((SELECT id FROM recipe_map WHERE title = 'Scrambled Eggs'), (SELECT id FROM ingredient_map WHERE name = 'Egg'),
     '3', 'pieces'),
    ((SELECT id FROM recipe_map WHERE title = 'Scrambled Eggs'), (SELECT id FROM ingredient_map WHERE name = 'Butter'),
     '1', 'tablespoon'),
    ((SELECT id FROM recipe_map WHERE title = 'Scrambled Eggs'), (SELECT id FROM ingredient_map WHERE name = 'Salt'),
     '1', 'pinch'),

    -- Chicken Fried Rice
    ((SELECT id FROM recipe_map WHERE title = 'Chicken Fried Rice'),
     (SELECT id FROM ingredient_map WHERE name = 'Chicken'), '200', 'grams'),
    ((SELECT id FROM recipe_map WHERE title = 'Chicken Fried Rice'),
     (SELECT id FROM ingredient_map WHERE name = 'Rice'), '2', 'cups'),
    ((SELECT id FROM recipe_map WHERE title = 'Chicken Fried Rice'),
     (SELECT id FROM ingredient_map WHERE name = 'Onion'), '1', 'piece'),

    -- Beef Stir Fry
    ((SELECT id FROM recipe_map WHERE title = 'Beef Stir Fry'), (SELECT id FROM ingredient_map WHERE name = 'Beef'),
     '300', 'grams'),
    ((SELECT id FROM recipe_map WHERE title = 'Beef Stir Fry'), (SELECT id FROM ingredient_map WHERE name = 'Onion'),
     '1', 'piece'),
    ((SELECT id FROM recipe_map WHERE title = 'Beef Stir Fry'), (SELECT id FROM ingredient_map WHERE name = 'Garlic'),
     '2', 'cloves'),

    -- Pasta Carbonara
    ((SELECT id FROM recipe_map WHERE title = 'Pasta Carbonara'), (SELECT id FROM ingredient_map WHERE name = 'Pasta'),
     '200', 'grams'),
    ((SELECT id FROM recipe_map WHERE title = 'Pasta Carbonara'), (SELECT id FROM ingredient_map WHERE name = 'Egg'),
     '2', 'pieces'),
    ((SELECT id FROM recipe_map WHERE title = 'Pasta Carbonara'), (SELECT id FROM ingredient_map WHERE name = 'Cheese'),
     '50', 'grams');

-- Thêm instructions cho Pancakes
INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 1, 'Mix flour, eggs, and milk in a large bowl until smooth', NOW(), NOW()
FROM recipes
WHERE title = 'Pancakes';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 2, 'Heat a non-stick pan over medium heat', NOW(), NOW()
FROM recipes
WHERE title = 'Pancakes';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 3, 'Pour batter into pan and cook until bubbles form', NOW(), NOW()
FROM recipes
WHERE title = 'Pancakes';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 4, 'Flip and cook other side until golden brown', NOW(), NOW()
FROM recipes
WHERE title = 'Pancakes';

-- Thêm instructions cho Scrambled Eggs
INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 1, 'Beat eggs in a bowl with salt and pepper', NOW(), NOW()
FROM recipes
WHERE title = 'Scrambled Eggs';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 2, 'Melt butter in a non-stick pan over medium heat', NOW(), NOW()
FROM recipes
WHERE title = 'Scrambled Eggs';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 3, 'Pour eggs and stir gently until cooked but still creamy', NOW(), NOW()
FROM recipes
WHERE title = 'Scrambled Eggs';

-- Thêm instructions cho Chicken Fried Rice
INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 1, 'Cook rice according to package instructions', NOW(), NOW()
FROM recipes
WHERE title = 'Chicken Fried Rice';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 2, 'Cut chicken into small pieces and season with salt and pepper', NOW(), NOW()
FROM recipes
WHERE title = 'Chicken Fried Rice';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 3, 'Heat oil and stir-fry chicken until cooked through', NOW(), NOW()
FROM recipes
WHERE title = 'Chicken Fried Rice';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 4, 'Add vegetables and rice, stir-fry until heated through', NOW(), NOW()
FROM recipes
WHERE title = 'Chicken Fried Rice';

-- Thêm instructions cho Beef Stir Fry
INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 1, 'Slice beef thinly against the grain', NOW(), NOW()
FROM recipes
WHERE title = 'Beef Stir Fry';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 2, 'Marinate beef with soy sauce and garlic', NOW(), NOW()
FROM recipes
WHERE title = 'Beef Stir Fry';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 3, 'Stir-fry vegetables in hot wok', NOW(), NOW()
FROM recipes
WHERE title = 'Beef Stir Fry';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 4, 'Add beef and stir-fry until cooked', NOW(), NOW()
FROM recipes
WHERE title = 'Beef Stir Fry';

-- Thêm instructions cho Pasta Carbonara
INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 1, 'Cook pasta in salted boiling water', NOW(), NOW()
FROM recipes
WHERE title = 'Pasta Carbonara';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 2, 'Mix eggs, cheese, and black pepper in a bowl', NOW(), NOW()
FROM recipes
WHERE title = 'Pasta Carbonara';

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 3, 'Toss hot pasta with egg mixture until creamy', NOW(), NOW()
FROM recipes
WHERE title = 'Pasta Carbonara';

-- Thêm instructions cho các recipes còn lại
INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 1, 'Prepare ingredients following recipe specifications', NOW(), NOW()
FROM recipes
WHERE title NOT IN ('Pancakes', 'Scrambled Eggs', 'Chicken Fried Rice', 'Beef Stir Fry', 'Pasta Carbonara');

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 2, 'Cook according to traditional methods', NOW(), NOW()
FROM recipes
WHERE title NOT IN ('Pancakes', 'Scrambled Eggs', 'Chicken Fried Rice', 'Beef Stir Fry', 'Pasta Carbonara');

INSERT INTO instructions (recipe_id, step_number, description, created_at, updated_at)
SELECT id, 3, 'Season to taste and serve hot', NOW(), NOW()
FROM recipes
WHERE title NOT IN ('Pancakes', 'Scrambled Eggs', 'Chicken Fried Rice', 'Beef Stir Fry', 'Pasta Carbonara');