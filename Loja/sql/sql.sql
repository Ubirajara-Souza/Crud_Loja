CREATE TABLE produtos (
	id SERIAL PRIMARY KEY,
	nome VARCHAR,
	descricao VARCHAR, 
	preco DECIMAL,
	quantidade INTEGER
);

SELECT * FROM produtos;