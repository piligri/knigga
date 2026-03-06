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

type QRdata struct {
	Fio      string  //ФИО исполнителя
	Etap     string  //Номер Этапа
	Operacia string  //Строка операции
	Rc       string  //Наименование рабочего центра
	Uchastok string  //Участок исполнителя
	PlanM    float32 //Метраж из заявки на этап
	FaktM    string  //Фактическая выработка из счетчика линии
	QRFiao   string  //Считанный и обработанный QR Исполнителя
	QREtap   string  //Считанный и обработанный QR Этапа
}

type QRstatus struct {
	InitQRWindow        bool   //Окно считывания QR кода открыто
	FioReadComplite     bool   //Код QR ФИО считан
	EtapReadComplite    bool   //Код QR Этапа считан
	EtapProcess         bool   //Этап в работе - TRUE, этап не начат FALSE
	EtapEnd             bool   //Этап завершен
	AllDataComplite     bool   //Все данные считаны и готовы к записи в 1с
	AllDataReadytoWrite bool   //?????
	NextQRAccept        bool   //Разрешение считывания QR (инверсное поле, FALSE - разрешено чтение)
	ScanerQREnable      bool   //Разрешение работы сканера
	QR                  string //Считанный QR
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
