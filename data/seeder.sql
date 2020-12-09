/**
  * CUSTOMERS
 */
INSERT INTO `customers` (`
id`,`name
`, `email`, `password`, `address`, `created_at`, `updated_at`)
VALUES
(1, 'Alessandro', 'rvalessandro@gmail.com', 'rv123', 'Bali no 2', '2020-12-08', '2020-12-08'),
(2, 'Michelle', 'michele@gmail.com', 'michele123', 'Cianjur no 3', '2020-12-08', '2020-12-08'),
(3, 'Daniel Alexander', 'daniela@gmail.com', 'daniel123', 'Bandung no 4', '2020-12-08', '2020-12-08'),
(4, 'Daniel Christianto', 'christianto@gmail.com', 'dc123', 'Kopo no 4', '2020-12-08', '2020-12-08'),
(5, 'Hanjaya', 'hanjaya@gmail.com', 'hanjaya123', 'Bandung no 2', '2020-12-08', '2020-12-08');

/**
  * CATEGORIES
 */
INSERT INTO categories
        (id, name, created_at, updated_at)
VALUES
        (1, 'Men', '2020-12-08', '2020-12-08');
INSERT INTO categories
        (id, name, created_at, updated_at)
VALUES
        (2, 'Women', '2020-12-08', '2020-12-08');

/**
  * PRODUCTS
 */
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (1, 1, 'White Clothes', 'White clothes with elastic fiber, appropiate for hard sporting, will absorb sweat fast',
                'assets/product/set_1.png', null, '1000000', '2020-12-08', '2020-12-08');
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (2, 1, 'Blue Clothes', 'blue clothes with elastic fiber, appropiate for hard sporting, will absorb sweat fast',
                'assets/product/set_2.png', null, '1500000', '2020-12-08', '2020-12-08');
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (3, 1, 'Blue Mask', 'Blue mask with thick fiber, will always block dust and virus from outside in',
                'assets/product/set_3.png', null, '1250000', '2020-12-08', '2020-12-08');
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (4, 1, 'Purple Mask', 'Purple mask with thick fiber, will always block dust and virus from outside in',
                'assets/product/set_4.png', null, '2000000', '2020-12-08', '2020-12-08');
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (5, 2, 'White fiber cloth',
                'White clothes with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm',
                'assets/product/set_5.png', null, '2500000', '2020-12-08', '2020-12-08');
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (6, 2, 'Blue fiber cloth',
                'blue clothes with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm',
                'assets/product/set_6.png', null, '2500000', '2020-12-08', '2020-12-08');
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (7, 2, 'Red White fiber cloth',
                'Pattern clothes with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm',
                'assets/product/set_7.png', null, '2500000', '2020-12-08', '2020-12-08');
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (8, 2, 'Blue Tiger fiber cloth',
                'Tiger Pattern clothes with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm',
                'assets/product/set_8.png', null, '2500000', '2020-12-08', '2020-12-08');
INSERT INTO millenial_fashion.products
        (id, category_id, name, description, image_url, preview_url, price, created_at,
        updated_at)
VALUES
        (9, 2, 'Red Jacket',
                'Red Jacket with thick fiber, appropiate for soft sporting, will absorb heat and keep you warm',
                'assets/product/set_8.png', null, '4500000', '2020-12-08', '2020-12-08');
