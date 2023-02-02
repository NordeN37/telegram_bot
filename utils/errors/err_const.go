package errors

const (
	ErrIdMissing         = Const("не верно указан Id объекта")
	ErrUuidMissing       = Const("не верно указан uuid объекта")
	ErrNumberMissing     = Const("не верно указан номер объекта")
	ErrPayloadExtract    = Const("не удалось извлечь данные тела запроса")
	ErrBadRequest        = Const("ошибка параметров запроса")
	ErrJsonUnMarshal     = Const("не удалось декодировать JSON")
	ErrJsonMarshal       = Const("не удалось упаковать данные в JSON")
	ErrInvalidSecretKey  = Const("неверный ключ")
	ErrParsePathParam    = Const("не удалось проанализировать параметр пути")
	ErrRecordingNotFound = Const("Запись не найдена")
	ErrValidate          = Const("ошибка валидации")
)
