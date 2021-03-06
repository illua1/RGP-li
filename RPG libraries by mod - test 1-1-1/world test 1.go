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







///рассы: человек, эльф, гном, дворф, демон, ангел, гоблин, гигант, химера, кобальд, грифон, карова, кролик, заяц, мышь, белка, змея, ящерица, зверолюд-пушной, зверолюд-варан, зверолюд-птичий, русалка, рыба, черепаха, ящерица, летучая мышь, лошадь, сова, орёл, кура, голубь, попугай, феникс, дракон, бог, полубог, голем, гамункул, насекомые, свинья, овца, лев, медведь, слизь, фея, дух, скелет, личь, вампир, упырь, зомби, ёжик, лягушки, сабака, лиса, олень.
/// -насекомые: бабочка, муравей, божъя коровка, гусеница, личинка-вредитель.
/// -рыба: окунь, карась, щука, рыба клоун, осъминог, кальмар.
/// -ящеры: варан, саламандра, степная, василиск, акцалоталь.
/// -мыши: мышь полевая, мышь одомашненая, мышь кровавая, крыса сточная, крыса кровавая, крыса чумная.
/// -лошади: лошадь деревенская, жеребец почтовый, конь военный, кабыла тягач, осёл обычный, пигас.
/// -собаки: собака, волк, фенрир.
/// -лисы: лиса цветная, норка, харёк, вуховная лиса. 









///LV ADD   ///################### ---------------------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------







type infoStruct struct{// Для записи данных о чём то
    name string; // Имя
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
type typeIdVec2Struct struct{
    types int;
	id int;
}

func TyRealTime(i float32) float32{
    return  i / 100. / 60.;
}





/// LV 0   ///################### ---------------------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------








type World struct{
	world infoStruct; // описание данного мира
	mobs mobsStuct; // перечень всех существ (в том чисе игроков)
	items itemsStruct; // перечень всех предметов
	maps mapsStruct; // карта мира и постройки на ней, лес, все области, ...
	
	physics physicsStruct; // данные для физики и времени
	time float32; // момент жизни (нынешний сдвиг)
}


func(w *World) FULL_GEN(side float64) {
	fmt.Println( "World: " + " Генератор для '"+w.world.name+"' запущен.");
} 
func(w *World) timeUP(time float32) {
	localTime := time + w.time;
	fmt.Println( "World: " + " ход времени на:", localTime, ".");
	w.physics.t += localTime;
	w.time = 0;
} 



/// LV 1   ///################### ---------------------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------








type mobsStuct struct{
	behavior behaviorListTypeSrtuct; // лист всех повидений с доступом по types/id (то есть AI всех сущностей, в том числе контроллер над игроками)
	race raceListTypeStruct; // описание всех виов сущностей (БЕСТИАРИЙ)
	baf bafListStruct; // все бафы, что есть тут
	
	mob []mobInfoStruct; // узел сборки сущности
}

type itemsStruct struct{ // ----- вместилище types/id
	// -патом- food [0]foodItemStruct; // еда
	// -патом- armor [0]armorItemStruct; // броня
	// ....
}

type mapsStruct struct{
	buildes buildListStruct; // все строения (дома, стены крепостей, укрепления, мосты, подземелья, ...)
	trees treesListStruct; // вся крупная ростительность (многолетняя)
	hole holeListStruct; // ямы, рвы, ущелья (перепады высот (как на нормал мапе, указывается не глубина, а сопративление движению каком то направлении))
	water waterListStruct; // все водные пространства, от луж после дождя, до акеана
	grass grassListStruct; // все травы, урожай на грядке, придорожный сорняк, ....
	way wayListStruct; // дороги ( в виде точек, соединённых кривыми для сглаживания, с указаниме ширины, качественности, ширины вурубленных по краям деревъев, пристижности (того, насколько опастной её считают))
	zoneMaps zoneMapsListStruct; // все зоны, в виде точек с весом, вокруг которых распростроняется их вличние ( как воронойз ) (указываются зоны: города, местность монстров, заны погоды, леса, ....)
}

type physicsStruct struct{
	t float32; // время в мире, с момента создания
}






/// LV 2   ///################### ---------------------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------







type bafListStruct struct{
	// -патом- addXp hillXp[]; // исцелялда
}

type behaviorListTypeSrtuct struct{ // ----- вместилище types/id
	// -патом- players []playerControllersStruct; // управллда (а в нутри неё будет сессия)
	// -патом- zombi []zombiAiStruct; // структура всех *мозгов зомби по types/id
	// ....
}

type raceListTypeStruct struct{ // БЕСТИАРИЙ  // ----- вместилище types/id   
	humas []humaTypeDescriptionStruct; // массив структур (однотипных объектов) всех людей (вименно в отдельном элементе этой структуры описан один человек (внешне))
	// -патом- zombi []zombiTypeDescriptionStruct; // ....   
	// .... и им всем нада свой invent inventListStruct;
}

type mobInfoStruct struct{ // не моба, а сущности. это узел компановки сущностей. тута всё, что омжет иметь сущность. выше их список, ниже деление на нсп, игрока, монстра, слизь, насекомое, ....
	name string; // имя сущности (не ник, он в сессии, имя будет в диалогах, ник в системе, айди для дискорда)
	xp int; // конечное число, олицетворяющее жизнь души, при обнулении, востанавливается толкьо при перерождении
	baf bafListStructl // все эффекты, влияющие на параметры игрока (благословение, проклятие, отравление, эффекты зелий, магия маскировки, состояния (горение, тонет, задыхается, ...) ) 
	
	gameType bool; // player/moba
	
	getRace typeIdVec2Struct; // в списке расс сущностей  
	id int; // номер сущности в списке этой рассы ( там все данный )
	status []string; // перечень всех заметок о состоянии этой сущности

}

type buildListStruct struct{}
type treesListStruct struct{}
type holeListStruct struct{}
type waterListStruct struct{}
type grassListStruct struct{}
type wayListStruct struct{}
type zoneMapsListStruct struct{}








///LV 3   ///################### --------------------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------









type humaTypeDescriptionStruct struct{
    age int;
	maxe_age int;
	growth float32;
	shoulders vector3; // вектор с данными о плечах (в виде бокса плеч)
	lowerBack vector3; // вектор с данными о грудной клетке (в виде бокса грудной клетки)
	ribCage vector3; // вектор с данными о тазе (в виде бокса таза)
	head vector3; // вектор с данными о голове (в виде бокса головы)
	//race descriptionStruct; // описание рассы и элеметов, не предусмотренных основными параметрами
	handPower int; // физическая сила рук
	legStrength int; // физическая сила рук
	power int; // стамина персонажа (выносливость)
	
	generickGet string; // ссылка на папку с итоговыми данными о персонаже (сгенерированная модель, пакет данных для хранения)
}

type discordSessionOfPlayerStruct struct{
    id string; // Идентефикатор пользователя, которому принадлежит этотот представитель структуры Player
}

type inventListStruct struct{ //пока мобы не прописаны, инвенторя не у кого нет, это не используется
    /*left_arm slotItemStruct; // Левая.Рука.Слоты для предметов
	right_arm slotItemStruct; // Правая.Рука.Слоты для предметов
	left_leg slotItemStruct; // Левая.Нога.Слоты для предметов
	right_leg slotItemStruct; // Правая.Нога.Слоты для предметов
	body slotItemStruct; // Тело.Слоты для предметов
	head slotItemStruct; // Слоты для предметов на голове */ //эта модель только для гуманоидов (людей)
	add slotItemStruct; // добавочные слоты (ну типа как хранилище) для предметов
}

type bafListStructl struct{
    getBafType []typeIdVec2Struct; // ссылки на бафы
}








///LV 4   ///################### ---------------------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------







type slotItemStruct struct{
    allItem []itemStruct; // отдел инвенторя 
	maxItemCount int; // Макс. слотов в этом отделе инвенторя
}

/*
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
*/







///LV 5   ///################### ---------------------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------







type itemStruct struct{
    mainNameItemType string; // Имя предмета (может быть изменено)Оригинальное имя (скрыто)
	lastNameItem string; // Имя предмета (может быть изменено)
	getTypes typeIdVec2Struct; // ссылка на предмет вы списке и типе
}








///LV 6   ///################### ---------------------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------
//-------------#-------------@-------------







type foodItemStruct struct{
    xit int; // высстанавливаемая сытость (каждего кусочка)
	getEffect []string; // наименование эффекта, аесли съел
	parts int; // число кусков, на которые можна поделить
	timeStories int; // срок годности (макс время жизни)
	timeLive int; // время жизни (на данный момент)
	// и вот тута уже можна подключать струкуту для передачи визуальных данных, но мне в лом
}





func main(){
    var TheHighestWorld World;
	TheHighestWorld.world.name = "world of sword and magic in the name of a generator"
	
	fmt.Println(TheHighestWorld);
	
	side := 5.2652;
	TheHighestWorld.FULL_GEN(side);
	
	for(false){
		TheHighestWorld.timeUP(1);
		fmt.Print(TyRealTime(TheHighestWorld.physics.t)); // время в мире, в секундах
	}
	TheHighestWorld.timeUP(1);
}

