package ebarimt

/*
# success PosAPI сан нь хэвийн ажиллах эсэхийг илтгэнэ

	true - ажиллагаа хэвийн
	false - ажиллагаа хэвийн бус

# Database Өгөгдөлийн баазтай харьцахад саад буй эсэхийг тодорхойлно.

	success true - ажиллагаа хэвийн
	false - ажиллагаа хэвийн бус

# Message Хэрэв алдаатай гэж үзвэл алдааны мессежийг буцаана
# Config Тохиргооны мэдээллүүдийг серверээс татаж тохируулсан
эсэхийг тодорхойлоно.

	success
		true - ажиллагаа хэвийн
		false - ажиллагаа хэвийн бус

	message Хэрэв алдаатай гэж үзвэл алдааны мессежийг буцаана

# Network Сүлжээ болон интернет холболт хэвийн байгааг татварын нэгдсэн системрүү холбогдож шалгана.

	Success
		true - ажиллагаа хэвийн
		false - ажиллагаа хэвийн бус

		message Хэрэв алдаатай гэж үзвэл алдааны мессежийг буцаана
*/

type APIOutput struct {
	Success  bool `json:"success"`
	Database struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"database"`
	Config struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"config"`
	Network struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"network"`
}

/*

# RegisterNo 	PosAPI ашиглаж буй татвар суутган төлөгчийн дугаар. ААН бол 7 оронтой тоо, иргэн бол 12 оронтой тоо /2 серийн үсэгийг тоон утгаруу хөрвүүлсэн байдлаар/

# BranchNo 	PosAPI ашиглаж буй татвар суутган төлөгчийн салбарын дугаар. 3 оронтой тоон утгаар /000, 001, 142 гэх мэт/

# PosId 	PosAPI-г системд бүртэгсэн дугаар.

# DbDirPath PosAPI-гийн ашиглаж буй SQLite өгөгдлийн баазыг агуулж буй
directory-гийн байршил

# ExtraInfo PosAPI-аас нэмэлтээр өгөх тайлбар, мэдээллүүд байна
	 countBill Илгээгдээгүй үлдсэн баримтын тоо
*/

type InformationOutput struct {
	RegisterNo string `json:"registerNo"`
	BranchNo   string `json:"branchNo"`
	PosID      string `json:"posId"`
	DBDirPath  string `json:"dbDirPath"`
	ExtraInfo  struct {
		CountBill    string `json:"countBill"`
		CountLottery int    `json:"countLottery"`
		LastSentDate string `json:"lastSentDate"`
		PosVersion   string `json:"posVersion"`
	} `json:"extraInfo"`
}

/*
# Amount 	(2 орны нарийвчлалтай тоо)

	Баримтын гүйлгээний нийт дүн (Бүх төрлийн татвар шингэсэн дүн)

# Vat 	(2 орны нарийвчлалтай тоо)

	Баримтын НӨАТ-ын нийт дүн

# CashAmount (2 орны нарийвчлалтай тоо)

	Бэлэн төлбөрийн дүн (Худалдан авагчаас хүлээн авсан дүн)

# NonCashAmount	(2 орны нарийвчлалтай тоо)

	Бэлэн бус төлбөрийн дүн (Худалдан авагчаас хүлээн авсан дүн)

# CityTax	(2 орны нарийвчлалтай тоо)

	Нийслэл хотын албан татварын нийт дүн

# DistrictCode (2 оронтой бүхэл тоо)

	Баримт хэвэлсэн орон нутгийн код /Татварын албаны орон нутгын код/

# PosNo		(4-6 оронтой бүхэл тоо)

	Тухайн байгууллагын дотоод кассын дугаар

# CustomerNo	(7 оронтой бүхэл тоо эсвэл Монгол улсын иргэний регистер)

	Худалдан авагч байгууллагын ТТД эсвэл Иргэний регистерийн дугаар

# BillType (1 оронтой бүхэл тоо)

	Баримтын төрөл
	1 	Байгууллагаас хувь иргэнд борлуулсан бараа, ажил үйлчилгээний баримт
	3 	Байгууллагаас байгууллагын хооронд борлуулсан бараа, ажил үйлчилгээний баримт
	5 	Нэхэмжлэхээр борлуулсан бараа, ажил, үйлчилгээний баримт

# BillIdSuffix (6 оронтой бүхэл тоо)

	Баримтын ДДТД-ыг давхцуулахгүйн тулд олгох дотоод дугаарлалт. Тухайн өдөртөө дахин давтагдашгүй дугаар байна.

# ReturnBillId (33 оронтой бүхэл тоо)

	Засварлах баримтын ДДТД

# TaxType (1 оронтой бүхэл тоо)

	Татварын төрөл

# InvoiceId (33 оронтой бүхэл тоо)

	Төлбөрийн баримтын харгалзах нэхэмжлэхийн ДДТД

# ReportMonth (yyyy-MM форматтай огноо)

	Баримтын харьяалагдах тайлант сар

# BranchNo (3 оронтой бүхэл тоо)

	Салбарын нэгжийн дугаар
*/
type PutInput struct {
	Amount           string             `json:"amount"`
	Vat              string             `json:"vat"`
	CashAmount       string             `json:"cashAmount"`
	NonCashAmount    string             `json:"nonCashAmount"`
	CityTax          string             `json:"cityTax"`
	DistrictCode     string             `json:"districtCode"`
	PosNo            string             `json:"dosNo"`
	CustomerNo       string             `json:"customerNo,omitempty"`
	BillType         string             `json:"billType,omitempty"`
	BillIdSuffix     string             `json:"billIdSuffix,omitempty"`
	ReturnBillId     string             `json:"returnBillId"`
	TaxType          string             `json:"taxType,omitempty"`
	InvoiceID        string             `json:"invoiceId"`
	ReportMonth      string             `json:"reportMonth"`
	BranchNo         string             `json:"branchNo"`
	Stocks           []Stock            `json:"stocks"`
	BankTransactions []BankTransactions `json:"bankTransactions"`
}

/*
# Success		Баримтыг бүртгэх процесс амжилттай болсон тухай илтгэнэ
	true – амжилттай бүртгэсэн
	false – амжилттай бүртгэж чадаагүй

# RegisterNo 	PosAPI эзэмшигч ААН-ийн ТТД эсвэл Татвар суутган төлөгч иргэний хөрвүүлсэн ТТД

# BillID 		Баримтын ДДТД 33 оронтой тоон утга /НӨАТ-ийн тухай хуулинд зааснаар/

# Date 		Баримт хэвлэсэн огноо

	Формат:
		yyyy-MM-dd hh:mm:ss

# MacAddress 	Баримтыг хэвлэсэн бүртгэлийн машиний MacAddress

# InternalCode 	Баримтын дотоод код

# BillType 		Баримтын төрөл

# QRData 	Баримтын баталгаажуулах Qr кодонд орох нууцлагдсан тоон утга

# Lottery 	Сугалааны дугаар

# LotteryWarningMsg Сугалаа дуусаж буй эсвэл сугалаа хэвлэх боломжгүй болсон талаар мэдээлэл өгөх утга

# ErrorCode Хэрэв алдаа илэрсэн бол уг алдааны код message Алдааны мэдээллийн текстэн утга

*/

type PutOutput struct {
	Success           bool   `json:"success"`
	RegisterNo        string `json:"registerNo"`
	BillID            string `json:"billId"`
	Date              string `json:"date"`
	MacAddress        string `json:"macAddress"`
	InternalCode      string `json:"internalCode"`
	BillType          string `json:"billType"`
	QRData            string `json:"qrData"`
	Lottery           string `json:"lottery"`
	LotteryWarningMsg string `json:"lotteryWarningMsg"`
	ErrorCode         int    `json:"errorCode"`
	Message           string `json:"message"`
}

/*

# Code (Дурын тэмдэгт мөр)
	Бараа, үйлчилгээний код /байгууллагын дотоод код/

# Name (Дурын тэмдэгт мөр)
	Бараа, үйлчилгээний нэр /байгууллагын дотоод нэр/

# MeasureUnit (Дурын тэмдэгт мөр)
	Хэмжих нэгж

# Qty (2 орны нарийвчлалтай тоо)
	Тоо, хэмжээ

# UnitPrice (2 орны нарийвчлалтай тоо)
	Нэгж үнэ (Бүх төрлийн татвар шингэсэн дүн)

# TotalAmount (2 орны нарийвчлалтай тоо)
	Нийт үнэ (Бүх төрлийн татвар шингэсэн дүн)

# CityTax (2 орны нарийвчлалтай тоо)
	Нийслэл хотын албан татварын нийт дүн

# Vat (2 орны нарийвчлалтай тоо)
	НӨАТ-ын нийт дүн

# BarCode (Бүхэл тоо)
	Барааны зураасан код эсвэл бараа,

*/

type Stock struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	MeasureUnit string `json:"measureUnit"`
	Qty         string `json:"qty"`
	UnitPrice   string `json:"unitPrice"`
	TotalAmount string `json:"totalAmount"`
	CityTax     string `json:"cityTax"`
	Vat         string `json:"vat"`
	BarCode     string `json:"barCode"`
}

/*

# Rrn (12 оронтой бүхэл тоо)
	Бэлэн бус гүйлгээний баримтын дугаар

# BankId (2 оронтой бүхэл тоо)
	Пос терминалын эзэмшигч банкны код

#TerminalId (Латин үсэг тоо холилдсон 6 ба түүнээс дээш тэмдэгт бүхий тэмдэгт мөр)
	Пос терминалын дугаар

#ApprovalCode (10 оронтой латин үсэг тооноос бүрдсэн тэмдэгт мөр)
	Бэлэн бус гүйлгээний зөвшөөрлийн код

#Amount (2 орны нарийвчлалтай тоо)
	Бэлэн бус гүйлгээний дүн

*/

type BankTransactions struct {
	Rrn          string `json:"rrn"`
	BankID       string `json:"bankId"`
	TerminalID   string `json:"terminalId"`
	ApprovalCode string `json:"approvalCode"`
	Amount       string `json:"amount"`
}

/*
# ReturnBillId (33 орон бүхий тоон утга)
	Хүчингүй болгох баримтын ДДТД

# Date Формат: yyyy-MM-dd hh:mm:ss
	Баримт хэвлэсэн огноо
*/

type BillInput struct {
	ReturnBillID string `json:"returnBillId"`
	Date         string `json:"date"`
}

type BillOutput struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

type DataOutput struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}
