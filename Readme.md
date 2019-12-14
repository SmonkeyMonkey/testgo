Тестовое задание Go-developer:

Создать api для работы с данными пользователей с использованием баз данных Redis и MongoDb.

Требуемый функционал:

Работа с пользователями (MongoDb):
создать таблицу users и заполнить её тестовыми данными (ссылка ниже) +
разработать функционал добавления новых пользователей с валидацией полей на соответствие формату (email, дата и т.д.) +,without date validation(all data is required)
разработать функционал запроса списка пользователей с постраничной навигацией (с фильтрами по полям) +

Работа с играми (MongoDb и Redis):
создать таблицу user_games и заполнить её тестовыми данными для каждого пользователя (тех что добавили в таблицу users с тестового набора), 
набор тестовых данных по играм можно взять по ссылке ниже на каждого пользователя должно быть рандомное количество игр с набора данных (минимум 5000 записей на пользователя). 

При добавлении игры в таблицу также увеличивается счетчик количества сыгранных игр в таблице в Redis (Sorted Set в котором лежат значения в формате user_id->count_games) -

разработать функционал добавления новых игр +
разработать функционал запроса списка игр с постраничной навигацией (с фильтрами по полям) +,

также нужно разработать функционал запроса статистики по играм (с фильтрами по дате): -
- статистика должна быть сгруппирована по дням (сумма игр за в день)
- также должна быть возможность получить статистику сгруппированную по номерам игр и дням
     -    разработать функционал получения рейтинга пользователя (рейтинг считается по количеству сыгранных пользователем игр) (использовать данные из Redis), api должно отдавать данные с постраничной навигацией.
      В ответе должна быть не только информация о положении пользователя в рейтинге но и информация о нем (имя, фамилия, дата рождения)

Смысл тестового задания в том, чтобы увидеть навыки кандидата по работе со структурами данных, оптимальном хранении данных и оптимизации запросов к БД.
Результаты выполнения разместить на github.

Ссылки с тестовыми наборами:
Пользователи: https://drive.google.com/file/d/1tjubsoSwdzPK553ovvmMZs9qQwMjlKh1/view?usp=sharing
Данные по играм: https://drive.google.com/file/d/1N_6pG7hxMcTJtB2MGAZZGe6_ZRfS21Mr/view?usp=sharing