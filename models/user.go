package models

import "time"

// Modelo de usuário
type UserAcompanhante struct {
	ID             uint       `gorm:"primaryKey"`
	Username       string     `json:"username"`
	CPF            string     `json:"cpf"`
	Password       string     `json:"password"`
	Email          string     `json:"email"`
	BirthDate      *time.Time `json:"birthdata"`
	RenovationDate *time.Time `json:"renovationdate"`
	Gender         string     `json:"gender"`
	Phone          string     `json:"phone"`
	PinCode        string     `json:"pincode"`
	CreatedAt      int64      `gorm:"autoUpdateTime"`
	UpdatedAt      int64      `gorm:"autoCreateTime"`
	//Product      Product?    @relation(fields: [productId], references: [id])
	//productId    Int?        @map("id_product")
	//Address      Address?    @relation(fields: [addressId], references: [id])
	//addressId    Int?        @map("id_address")
	//Checkpoint   Checkpoint? @relation(fields: [checkpointId], references: [id])
	//checkpointId Int?        @map("id_checkpoint")
}

type Anuncio struct {
	ID             uint       `gorm:"primaryKey"`
	LocalService string
	Gender         string     `json:"gender"`
	Title string 
	Describe string
	Ethnicity string
	EyesColor string
	Hair string
	HairSize string
	FootSize string
	Silicone string
	Tatoo string
	Piercings string
	Smoker string
	Ultimo acesso
	 Valor (Opcional)
	 Fotos
	 Sobre mim
	 Meus Serviços
	 Local de encontro
	 Horário
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `json:"username"`
	CPF       string `json:"cpf"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	CreatedAt int64  `gorm:"autoUpdateTime"`
	UpdatedAt int64  `gorm:"autoCreateTime"`
}

// Modelo do carro do usuário
type Carro struct {
	UserID     uint
	ID         uint `gorm:"primaryKey"`
	Modelo     string
	Fabricante string
	Ano        string
	Cor        string
}
