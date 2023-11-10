CREATE TABLE users (
    ID SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    deposit DECIMAL(10, 2) NOT NULL DEFAULT 0
);

CREATE TABLE equipment (
    equipment_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    description TEXT
);

CREATE TABLE checkouts (
    id SERIAL PRIMARY KEY,
    user_id INT,
    equipment_id INT NOT NULL,
    rental_date DATE NOT NULL,
    return_date DATE NOT NULL,
    rental_days Varchar(20),
    total_cost DECIMAL(10,2),
    FOREIGN KEY (user_id) REFERENCES users(ID),
    FOREIGN KEY (equipment_id) REFERENCES equipment(equipment_id)
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    user_id INT,
    checkout_id INT,
    rental_days INT NOT NULL,
    total_cost DECIMAL(10, 2) NOT NULL,
    payment DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(ID),
    FOREIGN KEY (checkout_id) REFERENCES checkouts(id)
);

INSERT INTO products (name, price, description)
VALUES ('Excavator (Bego)', 1000000.00, 'Excavator (bego) merupakan jenis alat berat yang secara umum digunakan untuk melakukan penggalian pada tanah dan memindahkan tanah atau material lainnya ke dalam truk muatan.'),
       ('Bulldozer', 500000.00, '27-Bulldozer merupakan alat berat yang secara umum digunakan untuk mengolah lahan. Umumnya digunakan untuk mendorong material tanah, hasil galian, baik itu ke arah depan, samping, ataupun untuk membuat suatu timbunan material.'),
       ('Wheel Loader', 1000000.00, 'Wheel loader merupakan alat yang memiliki fungsi tak terlalu berbeda dari bulldozer, yaitu digunakan untuk memindahkan material atau barang dari suatu alat atau tempat ke alat atau tempat yang lainnya. Cara kerjanya dengan menggali (melakukan loader).');