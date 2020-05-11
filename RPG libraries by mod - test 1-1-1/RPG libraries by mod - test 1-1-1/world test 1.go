package main


import (
    "fmt"
	)

/////////////////////////////////////////:-------
///r/rr////pppp/////gggg/g///////////////:-------
///rr//r///p///p///g////gg///////////////:==-----
///r///////p///p///g/////g///////////////:--==-----   |
///r///////pppp/////ggggg////////////////:-------========
///r///////p/////////g///g///////////////:--==-----   |
///r///////p/////////gggg////////////////:==-----
/////////////////////////////////////////:-------
//By MOD




///LV ADD   ///################### -------------
type infoStruct struct{// Для записи данных о чём то
    name string; // Имя
	id string; // Порядковый номер
	comment string; // Коментарий создателя
}
type vector3 struct{ // XYZ вектор
    x float32; // ось X
	y float32; // ось Y
	z float32; // ось Z
}
type colorStruct struct{ // RGB вектор
    r float32; // Канал цвета R
    g float32; // Канал цвета G
    b float32; // Канал цвета B
}
type posRotSizeColorDataStruct struct{
    pos vector3; // Позиция объекта
	rot vector3; // Угол объекта
	size vector3; // Размер объекта
	color colorStruct; // Основной цвет
	
	add_rot vector3; // Угол изгиба
	add_color colorStruct; // Добавочный цвет
	add_type int; // тип используемого объекта
}

/// LV 0
type World struct{
	world infoStruct; // описание данного мира
	mobs mobsStuct; // перечень всех существ (в том чисе игроков)
	items itemsStruct; // перечень всех предметов
	maps mapsStruct; // карта мира и постройки на ней, лес, все области, ...
}




/// LV 1   ///################### -------------
type MobsStuct struct{
    players [0]playerStruct; // игроки
	pnc [0]npsStruct; // неигровые персонажи
	mob [0]mobStruct; // крупные животные и агрессивные существа
	fonMobs [0]fonMobStruct; //птици на фоне, рыба, мыши, ....
}
type ItemsStruct struct{
	item [0]ItemStruct; // данные об итеме (предмете)
}
type MapsStruct struct{
	buildes buildListStruct; // все строения (дома, стены крепостей, укрепления, мосты, подземелья, ...)
	trees treesListStruct; // вся крупная ростительность (многолетняя)
	hole holeListStruct; // ямы, рвы, ущелья (перепады высот (как на нормал мапе, указывается не глубина, а сопративление движению каком то направлении))
	water waterListStruct; // все водные пространства, от луж после дождя, до акеана
	grass grassListStruct; // все травы, урожай на грядке, придорожный сорняк, ....
	way wayListStruct; // дороги ( в виде точек, соединённых кривыми для сглаживания, с указаниме ширины, качественности, ширины вурубленных по краям деревъев, пристижности (того, насколько опастной её считают))
	zoneMaps zoneMapsListStruct; // все зоны, в виде точек с весом, вокруг которых распростроняется их вличние ( как воронойз ) (указываются зоны: города, местность монстров, заны погоды, леса, ....)
}




/// LV 2   ///################### -------------
type playerStruct struct{
    nickName string; // имя игрока в этом мире
	xp int; // конечное число, олицетворяющее жизнь души, при обнулении, востанавливается толкьо при перерождении
	body bodyInfoStruct; // параметры внешнего вида, расса, параметры(сила, скорость, урон, ...), все о возросте и его исчеслении 
	session discordSessionOfPlayerStruct; // системная информация для связи с дискордом
	invent inventListStruct; // все слоты (для брони, переноска, в руках)
	baf bafListStructl // все эффекты, влияющие на параметры игрока (благословение, проклятие, отравление, эффекты зелий, магия маскировки, состояния (горение, тонет, задыхается, ...) ) 
}
//ЕЁЩ НЕ ОПИСАНО type npsStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type mobStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type fonMobStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type ItemStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type buildListStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type treesListStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type holeListStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type waterListStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type grassListStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type wayListStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type zoneMapsListStruct struct{}




///LV 3   ///################### -------------
type bodyInfoStruct struct{
    age int;
	maxe_age int;
	growth float32;
	shoulders vector3; // вектор с данными о плечах (в виде бокса плеч)
	lower back vector3; // вектор с данными о грудной клетке (в виде бокса грудной клетки)
	rib cage vector3; // вектор с данными о тазе (в виде бокса таза)
	head vector3; // вектор с данными о голове (в виде бокса головы)
	race descriptionStruct; // описание рассы и элеметов, не предусмотренных основными параметрами
	hand power int; // физическая сила рук
	leg strength int; // физическая сила рук
	power int; // стамина персонажа (выносливость)
}
type discordSessionOfPlayerStruct struct{
    id string; // Идентефикатор пользователя, которому принадлежит этотот представитель структуры Player
}
//ЕЁЩ НЕ ОПИСАНО type inventListStruct struct{}
//ЕЁЩ НЕ ОПИСАНО type bafListStructl struct{}





///LV 4   ///################### -------------


type descriptionStruct struct{
    asians int8; // близость к азиатам
    europeans int8; // близость а европейцам
    africans int8; // близость к африканцам
    width int; // уровень жира в щёках/коже
    cheekbones int; // скулы на лице
    nose vector3; //  width, long, down
	nose bool; // тип ушей (человек/зверо человек)
    ears vector3; //  angle, size, elven
    hair []hairDataStruct; // волосы
    eye colorStruct; //  rgb
    skin colorStruct; //  rgb
    scars []scarsDataStruct; // шрамы

    tails []posRotSizeColorDataStruct; // хвосты
    horns []posRotSizeColorDataStruct; // рога
    mustache []posRotSizeColorDataStruct; // усы
}



func main(){
    var TheHighestWorld World;
	fmt.Println(TheHighestWorld);
}

