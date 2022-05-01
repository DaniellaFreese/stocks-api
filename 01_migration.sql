CREATE TABLE IF NOT EXISTS stock(
  id serial NOT NULL,
  ticker varchar(45) NOT NULL,  
  company varchar(45) NOT NULL, 
  price numeric NOT NULL, 
  pe_ratio numeric NOT NULL, 
  PRIMARY KEY(id)
); 

INSERT INTO stock (id, ticker, company, price, pe_ratio)
VALUES (1, 'MSFT', 'Microsoft', 337.41, 37);

INSERT INTO stock (id, ticker, company, price, pe_ratio)
VALUES (2, 'RVN', 'Rivian', 112, 43); 

INSERT INTO stock (id, ticker, company, price, pe_ratio)
VALUES (3, 'JPM', 'JP Morgan', 160, 10); 