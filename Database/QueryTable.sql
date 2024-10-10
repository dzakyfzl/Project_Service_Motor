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
);

CREATE TABLE Memesan(
    Id_servis VARCHAR(4) NOT NULL ,
    Id_sparepart VARCHAR(4) NOT NULL,
    Foreign Key (Id_servis) REFERENCES Servis(Id_servis),
    Foreign Key (Id_sparepart) REFERENCES Sparepart(Id_sparepart)
);

--Example Data Entry

INSERT INTO pelanggan(`Id_pelanggan`, `Nama_pelanggan`, `Jenis_motor`, `Nomor_plat`) VALUES
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
    (20,'Chase','R1300 GS','DK2952CA'),
    (21,'Dzaki','Jupiter Z1','B6363EH'),
    (22,'zakhwan','Aerox','D1233DE'),
    (23,'Edgar','Vario 125','B6100OI'),
    (24,'Dyl','Vario 125','D9855EK'),
    (25,'Bagung','Vario 160','DK2288AP'),
    (26,'Faldik','XSR', 'D2993AE'),
    (27,'Madam','BeAT', 'R8842KH'),
    (28,'Tepe','Aerox','B5531RAH'),
    (29,'John','MT-25','D2045EZ'),
    (30,'Lucia','ZX10RR','P0450SA'),
    (31,'Afghan','H2','G0011DEN'),
    (32,'Glair','R1M','F1211RK'),
    (33,'Dinu','Z250','D7939DEH'),
    (34,'Agustino','Xeon','BK2412AL'),
    (35,'Irfan','CBR150R','AG3840DI'),
    (36,'Maxwell','R15','R2145TR'),
    (37,'Ehsan','NMAX','B3923AK'),
    (38,'Renata','TMAX','H3477DL'),
    (39,'Farah','Sprinter S','H1922RI'),
    (40,'Cheese','R1300 GS','DK2952CA'),
    (41,'Zaky','Jupiter Z1','B6363EH'),
    (42,'Dzawan','Aerox','D1233DE'),
    (43,'Regar','Vario 125','B6100OI'),
    (44,'Nita','Vario 125','D9855EK'),
    (45,'Wah','Vario 160','DK2288AP'),
    (46,'Rifal','XSR', 'D2993AE'),
    (47,'Kodam','BeAT', 'R8842KH'),
    (48,'Bepe','Aerox','B5531RAH'),
    (49,'Nianna','MT-25','D2045EZ'),
    (50,'Lucy','ZX10RR','P0450SA');

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
    ('P020','Kulit Jok','All Matic',80000,0),
    ('P021','Head Mesin','Aerox',40000,0),
    ('P022','Head Mesin','Vario',40000,0),
    ('P023','Head Mesin','H2',400000,0),
    ('P024','Head Mesin','MT-25',90000,0),
    ('P025','Busi','Aerox',150000,0),
    ('P026','Busi','Vario 125',100000,0),
    ('P027','Busi','Vario 160',145000,0),
    ('P028','Busi','Jupiter Z1',100000,0),
    ('P029','Busi','BeAT',100000,0),
    ('P030','Spakbor','Sprinter S',500000,0),
    ('P031','Spakbor','R1300 GS',10800000,0),
    ('P032','Spakbor','Z250',2300000,0),
    ('P033','Spakbor','XSR',430000,0),
    ('P034','Velg','All 1000cc',1800000,0),
    ('P035','Velg','All 100cc',300000,0),
    ('P036','Velg','All 250cc',750000,0),
    ('P037','Velg','Xeon',150000,0),
    ('P038','Velg','TMAX',3250000,0),
    ('P039','Disc Brake','All Bebek',80000,0),
    ('P040','Disc Brake','All Matic',80000,0),
    ('P041','Oli 1 Liter','R1300 GS',10800000,0),
    ('P042','Oli 1 Liter','Z250',2300000,0),
    ('P043','Oli 1 Liter','XSR',430000,0),
    ('P044','ShockBreaker','All 1000cc',1800000,0),
    ('P045','ShockBreaker','All 100cc',300000,0),
    ('P046','ShockBreaker','All 250cc',750000,0),
    ('P047','ShockBreaker','Xeon',150000,0),
    ('P048','ShockBreaker','TMAX',3250000,0),
    ('P049','Behel','All Bebek',80000,0),
    ('P050','Behel','All Matic',80000,0);

INSERT INTO Servis(Id_servis,Id_pelanggan,Total_Harga,Tanggal_Kunjungan) VALUES
    ('S001',1,40000,'11-03-2024'),
    ('S002',1,40000,'15-03-2024'),
    ('S003',1,400000,'19-03-2024'),
    ('S004',1,90000,'22-03-2024'),
    ('S005',1,150000,'25-03-2024'),
    ('S006',1,100000,'30-03-2024'),
    ('S007',1,145000,'07-04-2024'),
    ('S008',1,100000,'09-04-2024'),
    ('S009',1,100000,'20-04-2024'),
    ('S010',1,500000,'30-04-2024'),
    ('S011',1,10800000,'01-05-2024'),
    ('S012',1,2300000,'02-05-2024'),
    ('S013',1,430000,'03-05-2024'),
    ('S014',2,1800000,'20-03-2024'),
    ('S015',2,300000,'25-03-2024'),
    ('S016',2,750000,'30-03-2024'),
    ('S017',2,150000,'04-04-2024'),
    ('S018',2,3250000,'20-04-2024'),
    ('S019',3,80000,'05-03-2024'),
    ('S020',3,80000,'08-04-2024'),
    ('S021',4,40000,'05-03-2024'),
    ('S022',4,40000,'25-03-2024'),
    ('S023',4,400000,'30-03-2024'),
    ('S024',4,90000,'15-04-2024'),
    ('S025',5,150000,'02-01-2024'),
    ('S026',5,100000,'07-01-2024'),
    ('S027',5,145000,'20-03-2024'),
    ('S028',5,100000,'29-03-2024'),
    ('S029',5,100000,'30-05-2024'),
    ('S030',6,500000,'25-01-2024'),
    ('S031',6,10800000,'27-02-2024'),
    ('S032',6,2300000,'30-03-2024'),
    ('S033',6,430000,'24-04-2024'),
    ('S034',7,1800000,'01-01-2024'),
    ('S035',7,300000,'20-04-2024'),
    ('S036',7,750000,'25-05-2024'),
    ('S037',7,150000,'30-05-2024'),
    ('S038',7,3250000,'03-06-2024'),
    ('S039',8,80000,'20-03-2024'),
    ('S040',8,80000,'30-04-2024'),
    ('S041',9,10800000,'15-02-2024'),
    ('S042',9,2300000,'10-03-2024'),
    ('S043',9,430000,'20-04-2024'),
    ('S044',10,1800000,'05-01-2024'),
    ('S045',10,300000,'09-03-2024'),
    ('S046',10,750000,'20-03-2024'),
    ('S047',10,150000,'25-04-2024'),
    ('S048',10,3250000,'30-07-2024'),
    ('S049',11,80000,'01-07-2024'),
    ('S050',11,80000,'15-08-2024'),
    ('S051',12,40000,'03-01-2024'),
    ('S052',12,40000,'20-02-2024'),
    ('S053',12,400000,'19-03-2024'),
    ('S054',12,90000,'12-04-2024'),
    ('S055',12,150000,'03-05-2024'),
    ('S056',12,100000,'10-05-2024'),
    ('S057',12,145000,'09-06-2024'),
    ('S058',12,100000,'08-07-2024'),
    ('S059',12,100000,'09-08-2024'),
    ('S060',12,500000,'27-08-2024'),
    ('S061',12,10800000,'18-08-2024'),
    ('S062',12,2300000,'26-08-2024'),
    ('S063',12,430000,'09-09-2024'),
    ('S064',13,1800000,'20-01-2024'),
    ('S065',13,300000,'21-02-2024'),
    ('S066',13,750000,'09-02-2024'),
    ('S067',13,150000,'08-04-2024'),
    ('S068',13,3250000,'17-04-2024'),
    ('S069',14,80000,'06-01-2024'),
    ('S070',14,80000,'30-01-2024'),
    ('S071',15,40000,'29-02-2024'),
    ('S072',15,40000,'18-03-2024'),
    ('S073',15,400000,'10-04-2024'),
    ('S074',15,90000,'20-05-2024'),
    ('S075',16,150000,'24-01-2024'),
    ('S076',16,100000,'30-01-2024'),
    ('S077',16,145000,'25-02-2024'),
    ('S078',16,100000,'17-03-2024'),
    ('S079',16,100000,'28-04-2024'),
    ('S080',17,500000,'01-02-2024'),
    ('S081',17,10800000,'18-03-2024'),
    ('S082',17,2300000,'20-04-2024'),
    ('S083',17,430000,'30-05-2024'),
    ('S084',18,1800000,'27-01-2024'),
    ('S085',18,300000,'28-03-2024'),
    ('S086',18,750000,'12-05-2024'),
    ('S087',18,150000,'19-06-2024'),
    ('S088',18,3250000,'09-08-2024'),
    ('S089',19,80000,'21-03-2024'),
    ('S090',19,80000,'30-04-2024'),
    ('S091',20,10800000,'01-03-2024'),
    ('S092',20,2300000,'08-03-2024'),
    ('S093',20,430000,'07-04-2024'),
    ('S094',21,1800000,'17-02-2024'),
    ('S095',22,300000,'11-04-2024'),
    ('S096',23,750000,'10-03-2024'),
    ('S097',24,150000,'20-03-2024'),
    ('S098',25,3250000,'17-05-2024'),
    ('S099',26,80000,'10-05-2024'),
    ('S100',27,80000,'09-07-2024');



INSERT INTO Memesan(Id_sparepart, Id_servis) VALUES
('P013', 'S015'),
('P028', 'S074'),
('P023', 'S066'),
('P034', 'S009'),
('P016', 'S066'),
('P021', 'S091'),
('P041', 'S034'),
('P046', 'S025'),
('P019', 'S088'),
('P050', 'S002'),
('P025', 'S089'),
('P001', 'S067'),
('P038', 'S027'),
('P014', 'S084'),
('P032', 'S058'),
('P005', 'S046'),
('P009', 'S037'),
('P040', 'S078'),
('P043', 'S018'),
('P011', 'S059'),
('P024', 'S064'),
('P003', 'S008'),
('P048', 'S040'),
('P045', 'S073'),
('P018', 'S052'),
('P010', 'S020'),
('P022', 'S086'),
('P026', 'S010'),
('P031', 'S081'),
('P049', 'S050'),
('P002', 'S033'),
('P042', 'S032'),
('P047', 'S062'),
('P020', 'S043'),
('P006', 'S099'),
('P008', 'S048'),
('P004', 'S035'),
('P035', 'S011'),
('P036', 'S075'),
('P007', 'S041'),
('P030', 'S014'),
('P012', 'S039'),
('P033', 'S080'),
('P027', 'S093'),
('P017', 'S076'),
('P039', 'S022'),
('P029', 'S013'),
('P037', 'S003'),
('P044', 'S023'),
('P015', 'S055'),
('P025', 'S069'),
('P001', 'S092'),
('P038', 'S031'),
('P013', 'S097'),
('P028', 'S065'),
('P023', 'S053'),
('P034', 'S005'),
('P016', 'S090'),
('P021', 'S004'),
('P041', 'S030'),
('P046', 'S006'),
('P019', 'S070'),
('P050', 'S012'),
('P025', 'S047'),
('P001', 'S072'),
('P038', 'S024'),
('P014', 'S079'),
('P032', 'S085'),
('P005', 'S068'),
('P009', 'S063'),
('P040', 'S019'),
('P043', 'S095'),
('P011', 'S044'),
('P024', 'S016'),
('P003', 'S042'),
('P048', 'S038'),
('P045', 'S061'),
('P018', 'S049'),
('P010', 'S098'),
('P022', 'S021'),
('P026', 'S077'),
('P031', 'S026'),
('P049', 'S096'),
('P002', 'S071'),
('P042', 'S054'),
('P047', 'S060'),
('P020', 'S045'),
('P006', 'S094'),
('P008', 'S082'),
('P004', 'S087'),
('P035', 'S017'),
('P036', 'S056'),
('P007', 'S083'),
('P030', 'S057'),
('P012', 'S036'),
('P033', 'S029'),
('P027', 'S100'),
('P017', 'S051'),
('P039', 'S028');