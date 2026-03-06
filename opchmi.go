package main

var qrdata [9]string = [9]string{
	"s=Main_QRdata_FIO",      // 0
	"s=Main_QRdata_ETAP",     // 1
	"s=Main_QRdata_OPERACIA", // 2
	"s=Main_QRdata_RC",       // 3
	"s=Main_QRdata_UCHASTOK", // 4
	"s=Main_QRdata_PLANM",    // 5
	"s=Main_QRdata_FAKTM",    // 6
	"s=Main_QRdata_QREtap",   // 7
	"s=Main_QRdata_QRFio",    // 8
}

var qrstatus [10]string = [10]string{
	"s=Main_QRStatus_QRBa", //0 Окно считывания QR открыто
	"s=Main_QRStatus_QRBb", //1 Считан ФИО
	"s=Main_QRStatus_QRBc", //2 Считан Этап
	"s=Main_QRStatus_QRBd", //3 Выполнение Этапа (true - этап начат, false - этап не начат)
	"s=Main_QRStatus_QRBe", //4 Завершение этапа
	"s=Main_QRStatus_QRBf", //5 Все поля данных заполнены
	"s=Main_QRStatus_QRBg", //6 Данные готовы к записи #6
	"s=Main_QRStatus_QRBh", //7 Реактивация считывания следующего кода (установить в false)
	"s=Main_QRStatus_QRBj", //8 Разрешение работы сканера (true - сканер обрабатывает данные)
	"s=Main_QRStatus_QRN",  //9 Считанный qr
}
