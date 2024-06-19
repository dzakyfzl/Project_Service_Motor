-- Active: 1716460248431@@127.0.0.1@3306@servicemotor

CREATE servismotor;

USE servismotor;

CREATE TABLE Pelanggan(  
    Id_pelanggan INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    Nama_pelanggan VARCHAR(255),
    Jenis_motor VARCHAR(255),
    Nomor_plat VARCHAR(255)
);

CREATE TABLE Sparepart(
    Id_sparepart VARCHAR(4) PRIMARY KEY,
    Nama_sparepart VARCHAR(255),
    Harga_sparepart INT,
    Jumlah_terjual INT,
    Jenis_motor VARCHAR(255)
);

CREATE TABLE Servis(
    Id_servis VARCHAR(4) NOT NULL PRIMARY KEY,
    Id_pelanggan INT NOT NULL,
    Total_Harga INT,
    Tanggal_kunjungan VARCHAR(10),
    Foreign Key (Id_pelanggan) REFERENCES Pelanggan(Id_pelanggan)
)

CREATE TABLE Memesan(
    Id_servis VARCHAR(4) NOT NULL ,
    Id_sparepart VARCHAR(4) NOT NULL,
    Foreign Key (Id_servis) REFERENCES Servis(Id_servis),
    Foreign Key (Id_sparepart) REFERENCES Sparepart(Id_sparepart)
);

INSERT INTO pelanggan(Id_pelanggan, `Nama_pelanggan`, `Jenis_motor`, `Nomor_plat`) VALUES
    (1,'Dzaky','Jupiter Z1','B6363EH'),
    (2,'Dzakhwan','Aerox','D1233DE'),
    (3,'Tegar','Vario 125','B6100OI'),
    (4,'Vito','Vario 125','D9855EK'),
    (5,'Agung','Vario 160','DK2288AP'),
    (6,'Rifaldi','XSR', 'D2993AE'),
    (7,'Adam','BeAT', 'R8842KH'),
    (8,'Depe','Aerox','B5531RAH'),
    (9,'Jean','MT-25','D2045EZ'),
    (10,'Lucid','ZX10RR','P0450SA'),
    (11,'Avgene','H2','G0011DEN'),
    (12,'Blair','R1M','F1211RK'),
    (13,'Ibnu','Z250','D7939DEH'),
    (14,'Gustino','Xeon','BK2412AL'),
    (15,'Ivan','CBR150R','AG3840DI'),
    (16,'Axel','R15','R2145TR'),
    (17,'Ikhsan','NMAX','B3923AK'),
    (18,'Sonata','TMAX','H3477DL'),
    (19,'Sarah','Sprinter S','H1922RI'),
    (20,'Chase','R1300 GS','DK2952CA');

INSERT INTO sparepart(`Id_sparepart`,`Nama_sparepart`,`Jenis_motor`,`Harga_sparepart`,`Jumlah_terjual`) VALUES
    ('P001','Windshield','Aerox',40000,0),
    ('P002','Windshield','Vario',40000,0),
    ('P003','Windshield','H2',400000,0),
    ('P004','Windshield','MT-25',90000,0),
    ('P005','Set Piston','Aerox',150000,0),
    ('P006','Set Piston','Vario 125',100000,0),
    ('P007','Set Piston','Vario 160',145000,0),
    ('P008','Set Piston','Jupiter Z1',100000,0),
    ('P009','Set Piston','BeAT',100000,0),
    ('P010','Blok Mesin','Sprinter S',500000,0),
    ('P011','Blok Mesin','R1300 GS',10800000,0),
    ('P012','Blok Mesin','Z250',2300000,0),
    ('P013','Blok Mesin','XSR',430000,0),
    ('P014','Kaliper','All 1000cc',1800000,0),
    ('P015','Kaliper','All 100cc',300000,0),
    ('P016','Kaliper','All 250cc',750000,0),
    ('P017','Kaliper','Xeon',150000,0),
    ('P018','Kaliper','TMAX',3250000,0),
    ('P019','Kulit Jok','All Bebek',80000,0),
    ('P020','Kulit Jok','All Matic',80000,0);