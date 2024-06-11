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
    Tanggal_kunjungan DATE,
    Foreign Key (Id_pelanggan) REFERENCES Pelanggan(Id_pelanggan)
)

CREATE TABLE Memesan(
    Id_servis VARCHAR(4) NOT NULL ,
    Id_sparepart VARCHAR(4) NOT NULL,
    Foreign Key (Id_servis) REFERENCES Servis(Id_servis),
    Foreign Key (Id_sparepart) REFERENCES Sparepart(Id_sparepart)
);
