1. Развёртывание (запустить docker daemon)

1.1. docker-compose up --build -d --remove-orphans --renew-anon-volumes

1.2. Для остановки: docker-compose down --volumes

1.3. При остановке и запуске заново все новые данные теряются, остаются только предподготовленные

2.0. Регистрация и логин

2.1. Автоматически сгенерированы 5 пользователей с парами логин -- пароль:
ilia1 -- 1111
ilia2 -- 2222
ilia3 -- 3333
ilia4 -- 4444
ilia5 -- 5555

2.2. Для регистрации нового пользователя:
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/Register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "ilia6",
    "password": "6666"
}'

2.3. При регистрации вернётся токен, который нужно будет (у нас нет фронта, так что руками) вставить в куки "user_token":
{
    "token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJfZXhwIjoxNjkyNTY4MjAxLCJfZ2RuIjoiX2Nhc3RpbmdfcmVjaXBlcyIsIl91c2VyX25hbWUiOiI0ZDlhNzE5MC1iN2VlLTQxZmYtODhlOS1iNDIwZGFiYzZjNWIifQ.soQHxgaNmxhV1ToGuESHQz5TEZb1ZAHaFoEJ8r7pJrRkvL83gPJijvFGsG4c4SvTKG9YnIkFn3SoK5teAiAy6A"
}
Этот токен (он потом используется в примерах) был создан для пользователя ilia7 (который не создаётся автоматически). Пользователя нет, а токен рабочий. Так как функционала удаления пользователей нет, то это не баг, а фича.)
Так-то проверку того, что пользователь токена всё ещё существует - не проблема.

2.4. При логине уже существующего пользователя в том же формате возвращается токен. Запрос на логин:
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/Login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "ilia6",
    "password": "6666"
}'

3. P.s., хоть и не в конце: twirp требудет непустое тело запроса (да и ответа тоже), так что запросы даже без "{}" в теле невалидны.

4. Получение рецептов 

4.1. Получение всех рецептов (не требует авторизации):
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/GetAllRecipes' \
--header 'Content-Type: application/json' \
--data-raw '{}'

4.2. Получения списка рецептов авторизованного пользователя (для того, чтобы смотреть/проверять проще было):
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/GetAllRecipesByUser' \
--header 'Content-Type: application/json' \
--header 'Cookie: user_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJfZXhwIjoxNjkyNTY4MjAxLCJfZ2RuIjoiX2Nhc3RpbmdfcmVjaXBlcyIsIl91c2VyX25hbWUiOiI0ZDlhNzE5MC1iN2VlLTQxZmYtODhlOS1iNDIwZGFiYzZjNWIifQ.soQHxgaNmxhV1ToGuESHQz5TEZb1ZAHaFoEJ8r7pJrRkvL83gPJijvFGsG4c4SvTKG9YnIkFn3SoK5teAiAy6A' \
--data-raw '{}'

Тут уже требуется токен в куках

4.3. Получение рецепта по айди (авторизация не требуется):
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/GetRecipe' \
--header 'Content-Type: application/json' \
--data-raw '{"id":77}'

4.4. Фильтрация по ингредиентам, минимальному времени, максимальному времени о сортировка:
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/GetAllRecipes' \
--header 'Content-Type: application/json' \
--data-raw '{
    "min_time": 300,
    "max_time": 1000,
    "time_order": "desc",
    "ingredients":["papaya", "banana"]
}'

4.5. При наличии ингредиентов в запросе проверяется наличие всех-всех-всех ингредиентов (если из примера убрать "banana", то рецептом будет 2 в итоге)

4.6. Если имеется поле "time_order", но там не указана сортировка desc, то будет asc. Без этого поля сортировки по умолчанию нет.

4.7. Минимальное время можно передать большим или равным максимальному. ¯\_(ツ)_/¯ 

5. Неконечный p.s. номер 2: при запросе генерируется валидация тел запросов, при отсутствии id, когда требуется, логина, пароля и пр. будет ошибка валидации запроса

6. Создание рецепта

6.1. Для создания рецепта требуется авторизация:
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/CreateRecipe' \
--header 'Content-Type: application/json' \
--header 'Cookie: user_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJfZXhwIjoxNjkyNTY4MjAxLCJfZ2RuIjoiX2Nhc3RpbmdfcmVjaXBlcyIsIl91c2VyX25hbWUiOiI0ZDlhNzE5MC1iN2VlLTQxZmYtODhlOS1iNDIwZGFiYzZjNWIifQ.soQHxgaNmxhV1ToGuESHQz5TEZb1ZAHaFoEJ8r7pJrRkvL83gPJijvFGsG4c4SvTKG9YnIkFn3SoK5teAiAy6A' \
--data-raw '{
    "name": "my genial recipe",
    "description": "my genial description",
    "ingredients": [
        "taburetka",
        "mylo",
        "verevka"
    ],
    "steps": [
        {
            "description": "my genial step1",
            "duration": 720
        },
        {
            "description": "my genial step 2",
            "duration": 1200
        }
    ]
}'

6.2. При создании название должно быть уникальным:
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/CreateRecipe' \
--header 'Content-Type: application/json' \
--header 'Cookie: user_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJfZXhwIjoxNjkyNTY4MjAxLCJfZ2RuIjoiX2Nhc3RpbmdfcmVjaXBlcyIsIl91c2VyX25hbWUiOiI0ZDlhNzE5MC1iN2VlLTQxZmYtODhlOS1iNDIwZGFiYzZjNWIifQ.soQHxgaNmxhV1ToGuESHQz5TEZb1ZAHaFoEJ8r7pJrRkvL83gPJijvFGsG4c4SvTKG9YnIkFn3SoK5teAiAy6A' \
--data-raw '{
    "name": "my genial recipe also",
    "description": "my genial description",
    "ingredients": [
        "taburetka",
        "mylo",
        "verevka"
    ],
    "steps": [
        {
            "description": "my genial step1",
            "duration": 720
        },
        {
            "description": "my genial step 2",
            "duration": 1200
        }
    ]
}'

6.3. При создании проверяется целостность данных: хотя бы один шаг, хотя бы один ингредиент, название и описание

7. Обновление рецепта

7.1. Для обновления требуется авторизации. Править можно только свои собственные рецепты (ингредиенты в теле запроса тоже поменялись): 
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/UpdateRecipe' \
--header 'Content-Type: application/json' \
--header 'Cookie: user_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJfZXhwIjoxNjkyNTY4MjAxLCJfZ2RuIjoiX2Nhc3RpbmdfcmVjaXBlcyIsIl91c2VyX25hbWUiOiI0ZDlhNzE5MC1iN2VlLTQxZmYtODhlOS1iNDIwZGFiYzZjNWIifQ.soQHxgaNmxhV1ToGuESHQz5TEZb1ZAHaFoEJ8r7pJrRkvL83gPJijvFGsG4c4SvTKG9YnIkFn3SoK5teAiAy6A' \
--data-raw '{
    "id": 101,
    "name": "my genial recipe HAS CHANGED",
    "description": "my genial description HAS CHANGED",
    "ingredients": [
        "mylo",
        "verevka"
    ],
    "steps": [
        {
            "description": "my genial step1",
            "duration": 720
        },
        {
            "description": "my genial step 2",
            "duration": 1200
        }
    ]
}'

7.2. При обновлении всё ещё проверяется целостность данных: хотя бы один шаг, хотя бы один ингредиент, название и описание

8. Удаление рецептов

8.1. Для удаления требуется авторизация. Удалять можно только свои собственные рецепты:
curl --location --request POST 'http://localhost:8000/twirp/api.ApiGateway/DeleteRecipe' \
--header 'Content-Type: application/json' \
--header 'Cookie: user_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJfZXhwIjoxNjkyNTY4MjAxLCJfZ2RuIjoiX2Nhc3RpbmdfcmVjaXBlcyIsIl91c2VyX25hbWUiOiI0ZDlhNzE5MC1iN2VlLTQxZmYtODhlOS1iNDIwZGFiYzZjNWIifQ.soQHxgaNmxhV1ToGuESHQz5TEZb1ZAHaFoEJ8r7pJrRkvL83gPJijvFGsG4c4SvTKG9YnIkFn3SoK5teAiAy6A' \
--data-raw '{
    "id": 101
}'

9. Структура базы:

9.1. Таблица с пользователями:
create table if not exists rec.users (
  uid uuid not null,
  login text not null,
  password text not null,
  constraint users_pk primary key (uid),
  constraint users_un unique (login)
);

9.2. Таблица с рецептами:
create table if not exists rec.recipes (
  user_uid uuid not null,
  id bigserial not null,
  name text not null,
  steps jsonb not null,
  description text not null,
  total_time int8 not null, -- в секундах
  constraint recipes_pk primary key (id),
  constraint recipes_un unique (name)
);
create index recipes_id_idx on rec.recipes (id);

9.3. Таблица с ингредиентами была вынесена отдельно. 
Хоть в рамках данной задачи это и не нужно (можно было бы хранить ингредиенты в массива, так как при единственной в проекте обработке мы их аггрегируем, то есть такое разреженное хранение - даже хуже).
Это задел на будущее, потому что 100% потреубуется какая-то обработка ингредиентов (предложение пользователям не только вбивать свои, но и выбирать среди существующих).

Скрипт: 
create table if not exists rec.recipe_ingredients (
  recipe_id int8 not null,
  name text not null
);
create index recipe_ingredients_name_idx on rec.recipe_ingredients (name); -- это как раз на будущее
create index recipe_ingredients_recipe_id_idx on rec.recipe_ingredients (recipe_id);