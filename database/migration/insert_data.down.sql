-- Xóa mối quan hệ trong bảng recipe_ingredients
DELETE FROM recipe_ingredients WHERE recipe_id IN (SELECT id FROM recipes WHERE title IN ('Pancakes', 'Scrambled Eggs'));

-- Xóa dữ liệu trong bảng recipes
DELETE FROM recipes WHERE title IN ('Pancakes', 'Scrambled Eggs');

-- Xóa dữ liệu trong bảng ingredients
DELETE FROM ingredients WHERE name IN ('Egg', 'Flour', 'Milk', 'Butter', 'Salt');
