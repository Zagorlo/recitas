create schema if not exists rec;

create table if not exists rec.users (
  uid uuid not null,
  login text not null,
  password text not null,
  constraint users_pk primary key (uid),
  constraint users_un unique (login)
);

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

create table if not exists rec.recipe_ingredients (
  recipe_id int8 not null,
  name text not null
);
create index recipe_ingredients_name_idx on rec.recipe_ingredients (name);
create index recipe_ingredients_recipe_id_idx on rec.recipe_ingredients (recipe_id);

insert into rec.users (uid,login,password) values
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81','ilia1','$2a$13$TAAtTQKzZfh5JDbELAgEZuliPxGAzRw.D0X62yz3AplsWdSrLd1Aq'),
('c95c630e-33d1-47e3-bc7f-192f180f841d','ilia2','$2a$13$cCraa9HG6yxyQiDqGMpW4u.vUbKgYzM/Q8iUSGfhjD8JFKNSCDCfe'),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166','ilia3','$2a$13$yScKIxqUjRdcMNT2LapXAuhrN9cfjKyHCprGxknBHHeFybNTrz9bu'),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d','ilia4','$2a$13$FKQKoCD7vAdS0.rohIV0p.sUpqfkzjng5Fq48lYgJt9.2.2V./a6C'),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a','ilia5','$2a$13$CIe2AbEmO6e.FaACkLLEx.0B8.urM4QG5f4qzd6l2SA8Q4FaZqI0m');

insert into rec.recipes (user_uid, name, description, steps, total_time) values
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name1', 'descr1', '[{"duration":3600,"description":"descr 60"},{"duration":1800,"description":"descr 30"},{"duration":300,"description":"descr 15"},{"duration":60,"description":"descr 15"}]', 5760),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name2', 'descr2', '[{"duration":15,"description":"descr 1200"},{"duration":1200,"description":"descr 1200"},{"duration":1800,"description":"descr 1800"},{"duration":600,"description":"descr 15"},{"duration":3600,"description":"descr 3600"}]', 7215),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name3', 'descr1', '[{"duration":3600,"description":"descr 1800"},{"duration":300,"description":"descr 1800"},{"duration":300,"description":"descr 60"},{"duration":1200,"description":"descr 600"}]', 5400),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name4', 'descr2', '[{"duration":3600,"description":"descr 1200"},{"duration":30,"description":"descr 300"},{"duration":300,"description":"descr 15"}]', 3930),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name5', 'descr1', '[{"duration":1200,"description":"descr 3600"},{"duration":600,"description":"descr 600"},{"duration":1800,"description":"descr 300"}]', 3600),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name6', 'descr2', '[{"duration":600,"description":"descr 600"},{"duration":300,"description":"descr 300"},{"duration":60,"description":"descr 15"},{"duration":15,"description":"descr 3600"},{"duration":600,"description":"descr 3600"}]', 1575),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name7', 'descr1', '[{"duration":60,"description":"descr 3600"},{"duration":300,"description":"descr 15"},{"duration":3600,"description":"descr 60"},{"duration":1200,"description":"descr 60"}]', 5160),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name8', 'descr2', '[{"duration":1800,"description":"descr 15"},{"duration":300,"description":"descr 600"},{"duration":600,"description":"descr 15"},{"duration":3600,"description":"descr 1200"}]', 6300),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name9', 'descr1', '[{"duration":3600,"description":"descr 3600"},{"duration":1200,"description":"descr 30"},{"duration":300,"description":"descr 1800"},{"duration":300,"description":"descr 15"}]', 5400),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name10', 'descr2', '[{"duration":1200,"description":"descr 15"},{"duration":300,"description":"descr 3600"},{"duration":300,"description":"descr 30"},{"duration":600,"description":"descr 15"},{"duration":60,"description":"descr 15"}]', 2460),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name11', 'descr1', '[{"duration":1800,"description":"descr 300"}]', 1800),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name12', 'descr2', '[{"duration":1800,"description":"descr 60"},{"duration":30,"description":"descr 1200"},{"duration":1200,"description":"descr 300"},{"duration":60,"description":"descr 30"},{"duration":600,"description":"descr 30"}]', 3690),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name13', 'descr1', '[{"duration":60,"description":"descr 15"},{"duration":3600,"description":"descr 30"},{"duration":1800,"description":"descr 600"},{"duration":1200,"description":"descr 30"}]', 6660),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name14', 'descr2', '[{"duration":1800,"description":"descr 60"},{"duration":600,"description":"descr 600"}]', 2400),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name15', 'descr1', '[{"duration":300,"description":"descr 1800"},{"duration":30,"description":"descr 3600"},{"duration":15,"description":"descr 3600"},{"duration":1800,"description":"descr 3600"}]', 2145),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name16', 'descr2', '[{"duration":300,"description":"descr 300"},{"duration":600,"description":"descr 15"},{"duration":600,"description":"descr 300"},{"duration":30,"description":"descr 30"}]', 1530),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name17', 'descr1', '[{"duration":1200,"description":"descr 3600"},{"duration":600,"description":"descr 30"},{"duration":300,"description":"descr 1800"},{"duration":60,"description":"descr 30"}]', 2160),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name18', 'descr2', '[{"duration":300,"description":"descr 600"},{"duration":1200,"description":"descr 300"},{"duration":15,"description":"descr 1800"},{"duration":300,"description":"descr 15"}]', 1815),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name19', 'descr1', '[{"duration":30,"description":"descr 600"}]', 30),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name20', 'descr2', '[{"duration":300,"description":"descr 600"},{"duration":1200,"description":"descr 15"},{"duration":1200,"description":"descr 1200"},{"duration":1200,"description":"descr 1800"}]', 3900),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name21', 'descr1', '[{"duration":60,"description":"descr 600"}]', 60),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name22', 'descr2', '[{"duration":30,"description":"descr 60"}]', 30),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name23', 'descr1', '[{"duration":15,"description":"descr 300"},{"duration":600,"description":"descr 15"},{"duration":1200,"description":"descr 60"},{"duration":1200,"description":"descr 60"}]', 3015),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name24', 'descr2', '[{"duration":15,"description":"descr 30"},{"duration":30,"description":"descr 300"},{"duration":30,"description":"descr 600"}]', 75),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name25', 'descr1', '[{"duration":300,"description":"descr 1800"},{"duration":30,"description":"descr 600"},{"duration":300,"description":"descr 3600"},{"duration":15,"description":"descr 15"},{"duration":1800,"description":"descr 300"}]', 2445),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name26', 'descr2', '[{"duration":30,"description":"descr 1200"},{"duration":1800,"description":"descr 1800"},{"duration":3600,"description":"descr 300"},{"duration":1200,"description":"descr 30"}]', 6630),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name27', 'descr1', '[{"duration":1200,"description":"descr 60"},{"duration":3600,"description":"descr 3600"}]', 4800),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name28', 'descr2', '[{"duration":30,"description":"descr 600"},{"duration":1800,"description":"descr 60"},{"duration":600,"description":"descr 30"},{"duration":30,"description":"descr 300"}]', 2460),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name29', 'descr1', '[{"duration":30,"description":"descr 3600"},{"duration":1800,"description":"descr 3600"},{"duration":600,"description":"descr 3600"}]', 2430),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name30', 'descr2', '[{"duration":3600,"description":"descr 3600"},{"duration":15,"description":"descr 3600"},{"duration":30,"description":"descr 300"},{"duration":15,"description":"descr 15"}]', 3660),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name31', 'descr1', '[{"duration":60,"description":"descr 1200"},{"duration":30,"description":"descr 30"},{"duration":30,"description":"descr 600"},{"duration":3600,"description":"descr 3600"}]', 3720),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name32', 'descr2', '[{"duration":3600,"description":"descr 600"},{"duration":15,"description":"descr 600"},{"duration":60,"description":"descr 300"},{"duration":15,"description":"descr 600"}]', 3690),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name33', 'descr1', '[{"duration":60,"description":"descr 1800"},{"duration":1800,"description":"descr 1200"},{"duration":60,"description":"descr 30"},{"duration":60,"description":"descr 1800"}]', 1980),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name34', 'descr2', '[{"duration":600,"description":"descr 600"},{"duration":3600,"description":"descr 1200"}]', 4200),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name35', 'descr1', '[{"duration":30,"description":"descr 300"},{"duration":1200,"description":"descr 1200"},{"duration":30,"description":"descr 15"},{"duration":60,"description":"descr 1200"},{"duration":300,"description":"descr 1200"}]', 1620),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name36', 'descr2', '[{"duration":300,"description":"descr 1800"},{"duration":15,"description":"descr 30"},{"duration":30,"description":"descr 30"},{"duration":1800,"description":"descr 1200"},{"duration":3600,"description":"descr 30"}]', 5745),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name37', 'descr1', '[{"duration":1200,"description":"descr 60"},{"duration":600,"description":"descr 1800"},{"duration":30,"description":"descr 1200"},{"duration":3600,"description":"descr 30"}]', 5430),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name38', 'descr2', '[{"duration":15,"description":"descr 300"}]', 15),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name39', 'descr1', '[{"duration":300,"description":"descr 600"}]', 300),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name40', 'descr2', '[{"duration":600,"description":"descr 3600"},{"duration":3600,"description":"descr 1200"},{"duration":1800,"description":"descr 15"},{"duration":3600,"description":"descr 600"},{"duration":1200,"description":"descr 1200"}]', 10800),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name41', 'descr1', '[{"duration":1200,"description":"descr 3600"},{"duration":60,"description":"descr 3600"},{"duration":30,"description":"descr 1200"},{"duration":15,"description":"descr 60"}]', 1305),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name42', 'descr2', '[{"duration":1200,"description":"descr 300"}]', 1200),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name43', 'descr1', '[{"duration":600,"description":"descr 3600"},{"duration":1200,"description":"descr 1200"},{"duration":300,"description":"descr 300"},{"duration":600,"description":"descr 3600"},{"duration":1200,"description":"descr 3600"}]', 3900),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name44', 'descr2', '[{"duration":1800,"description":"descr 1200"},{"duration":1800,"description":"descr 1200"},{"duration":3600,"description":"descr 3600"},{"duration":15,"description":"descr 1800"},{"duration":15,"description":"descr 30"}]', 7230),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name45', 'descr1', '[{"duration":1800,"description":"descr 300"}]', 1800),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name46', 'descr2', '[{"duration":1200,"description":"descr 15"}]', 1200),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name47', 'descr1', '[{"duration":3600,"description":"descr 30"},{"duration":600,"description":"descr 600"},{"duration":30,"description":"descr 1200"},{"duration":600,"description":"descr 15"},{"duration":60,"description":"descr 3600"}]', 4890),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name48', 'descr2', '[{"duration":15,"description":"descr 1800"},{"duration":15,"description":"descr 300"},{"duration":1200,"description":"descr 1800"}]', 1230),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name49', 'descr1', '[{"duration":300,"description":"descr 15"},{"duration":1800,"description":"descr 1800"},{"duration":600,"description":"descr 15"}]', 2700),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name50', 'descr2', '[{"duration":1800,"description":"descr 60"},{"duration":3600,"description":"descr 3600"},{"duration":3600,"description":"descr 15"},{"duration":300,"description":"descr 3600"}]', 9300),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name51', 'descr1', '[{"duration":60,"description":"descr 300"},{"duration":1200,"description":"descr 1200"}]', 1260),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name52', 'descr2', '[{"duration":30,"description":"descr 1200"},{"duration":3600,"description":"descr 1800"},{"duration":3600,"description":"descr 600"},{"duration":1200,"description":"descr 1200"},{"duration":300,"description":"descr 15"}]', 8730),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name53', 'descr1', '[{"duration":30,"description":"descr 1200"},{"duration":1800,"description":"descr 30"}]', 1830),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name54', 'descr2', '[{"duration":3600,"description":"descr 1200"}]', 3600),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name55', 'descr1', '[{"duration":1800,"description":"descr 15"},{"duration":300,"description":"descr 30"},{"duration":15,"description":"descr 15"}]', 2115),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name56', 'descr2', '[{"duration":300,"description":"descr 300"},{"duration":3600,"description":"descr 1800"}]', 3900),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name57', 'descr1', '[{"duration":1200,"description":"descr 300"},{"duration":15,"description":"descr 60"},{"duration":30,"description":"descr 1800"},{"duration":30,"description":"descr 300"},{"duration":1200,"description":"descr 60"}]', 2475),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name58', 'descr2', '[{"duration":1200,"description":"descr 600"},{"duration":60,"description":"descr 600"},{"duration":15,"description":"descr 600"}]', 1275),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name59', 'descr1', '[{"duration":1200,"description":"descr 15"},{"duration":1800,"description":"descr 15"}]', 3000),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name60', 'descr2', '[{"duration":1200,"description":"descr 30"},{"duration":600,"description":"descr 300"},{"duration":30,"description":"descr 15"},{"duration":60,"description":"descr 3600"}]', 1890),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name61', 'descr1', '[{"duration":15,"description":"descr 1800"},{"duration":3600,"description":"descr 15"}]', 3615),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name62', 'descr2', '[{"duration":3600,"description":"descr 1200"},{"duration":60,"description":"descr 1800"}]', 3660),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name63', 'descr1', '[{"duration":30,"description":"descr 30"},{"duration":1800,"description":"descr 600"}]', 1830),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name64', 'descr2', '[{"duration":3600,"description":"descr 60"},{"duration":300,"description":"descr 600"},{"duration":15,"description":"descr 600"}]', 3915),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name65', 'descr1', '[{"duration":1200,"description":"descr 15"},{"duration":600,"description":"descr 3600"},{"duration":300,"description":"descr 1200"},{"duration":300,"description":"descr 15"}]', 2400),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name66', 'descr2', '[{"duration":1200,"description":"descr 15"},{"duration":1200,"description":"descr 3600"}]', 2400),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name67', 'descr1', '[{"duration":15,"description":"descr 60"},{"duration":60,"description":"descr 60"},{"duration":1800,"description":"descr 15"},{"duration":1200,"description":"descr 1800"}]', 3075),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name68', 'descr2', '[{"duration":300,"description":"descr 300"},{"duration":3600,"description":"descr 15"},{"duration":3600,"description":"descr 300"},{"duration":1800,"description":"descr 1200"},{"duration":30,"description":"descr 1200"}]', 9330),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name69', 'descr1', '[{"duration":1800,"description":"descr 300"},{"duration":600,"description":"descr 300"},{"duration":30,"description":"descr 1200"}]', 2430),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name70', 'descr2', '[{"duration":60,"description":"descr 1800"},{"duration":1800,"description":"descr 60"},{"duration":3600,"description":"descr 1800"},{"duration":1800,"description":"descr 30"}]', 7260),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name71', 'descr1', '[{"duration":15,"description":"descr 300"},{"duration":30,"description":"descr 30"},{"duration":3600,"description":"descr 1800"},{"duration":300,"description":"descr 300"}]', 3945),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name72', 'descr2', '[{"duration":300,"description":"descr 1200"}]', 300),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name73', 'descr1', '[{"duration":30,"description":"descr 60"},{"duration":3600,"description":"descr 300"},{"duration":1200,"description":"descr 600"}]', 4830),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name74', 'descr2', '[{"duration":3600,"description":"descr 300"},{"duration":15,"description":"descr 600"}]', 3615),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name75', 'descr1', '[{"duration":3600,"description":"descr 30"},{"duration":60,"description":"descr 3600"},{"duration":600,"description":"descr 60"},{"duration":3600,"description":"descr 3600"}]', 7860),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name76', 'descr2', '[{"duration":60,"description":"descr 30"},{"duration":300,"description":"descr 1200"}]', 360),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name77', 'descr1', '[{"duration":3600,"description":"descr 30"},{"duration":60,"description":"descr 60"}]', 3660),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name78', 'descr2', '[{"duration":1800,"description":"descr 60"}]', 1800),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name79', 'descr1', '[{"duration":3600,"description":"descr 30"},{"duration":1200,"description":"descr 3600"},{"duration":600,"description":"descr 600"},{"duration":1800,"description":"descr 600"},{"duration":60,"description":"descr 300"}]', 7260),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name80', 'descr2', '[{"duration":60,"description":"descr 1800"},{"duration":1200,"description":"descr 3600"},{"duration":3600,"description":"descr 1800"},{"duration":600,"description":"descr 60"}]', 5460),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name81', 'descr1', '[{"duration":1200,"description":"descr 15"}]', 1200),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name82', 'descr2', '[{"duration":60,"description":"descr 1800"},{"duration":30,"description":"descr 60"},{"duration":1800,"description":"descr 600"}]', 1890),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name83', 'descr1', '[{"duration":1200,"description":"descr 300"},{"duration":300,"description":"descr 1200"},{"duration":1800,"description":"descr 15"},{"duration":600,"description":"descr 1800"}]', 3900),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name84', 'descr2', '[{"duration":3600,"description":"descr 15"},{"duration":15,"description":"descr 600"},{"duration":600,"description":"descr 15"},{"duration":3600,"description":"descr 1800"}]', 7815),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name85', 'descr1', '[{"duration":600,"description":"descr 60"},{"duration":30,"description":"descr 1200"},{"duration":300,"description":"descr 3600"},{"duration":1800,"description":"descr 3600"},{"duration":1800,"description":"descr 1200"}]', 4530),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name86', 'descr2', '[{"duration":60,"description":"descr 60"},{"duration":1800,"description":"descr 1800"},{"duration":1200,"description":"descr 300"},{"duration":30,"description":"descr 30"},{"duration":60,"description":"descr 60"}]', 3150),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name87', 'descr1', '[{"duration":15,"description":"descr 300"}]', 15),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name88', 'descr2', '[{"duration":3600,"description":"descr 300"},{"duration":300,"description":"descr 3600"}]', 3900),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name89', 'descr1', '[{"duration":15,"description":"descr 60"}]', 15),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name90', 'descr2', '[{"duration":1800,"description":"descr 15"},{"duration":3600,"description":"descr 3600"},{"duration":300,"description":"descr 600"},{"duration":30,"description":"descr 3600"},{"duration":600,"description":"descr 3600"}]', 6330),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name91', 'descr1', '[{"duration":30,"description":"descr 1200"}]', 30),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name92', 'descr2', '[{"duration":60,"description":"descr 15"},{"duration":300,"description":"descr 1800"},{"duration":30,"description":"descr 30"},{"duration":15,"description":"descr 15"},{"duration":3600,"description":"descr 30"}]', 4005),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name93', 'descr1', '[{"duration":15,"description":"descr 1800"},{"duration":30,"description":"descr 3600"},{"duration":1200,"description":"descr 3600"}]', 1245),
('5a2f8e9a-1ced-46b5-be17-a3e76e85e166', 'name94', 'descr2', '[{"duration":1800,"description":"descr 60"},{"duration":300,"description":"descr 1800"}]', 2100),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name95', 'descr1', '[{"duration":1800,"description":"descr 15"},{"duration":30,"description":"descr 300"},{"duration":300,"description":"descr 300"},{"duration":600,"description":"descr 3600"},{"duration":600,"description":"descr 600"}]', 3330),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name96', 'descr2', '[{"duration":30,"description":"descr 30"},{"duration":3600,"description":"descr 1800"},{"duration":300,"description":"descr 300"}]', 3930),
('a5b01cfc-5a37-4c45-a140-64ce9350ba2a', 'name97', 'descr1', '[{"duration":30,"description":"descr 3600"},{"duration":600,"description":"descr 1200"},{"duration":3600,"description":"descr 60"},{"duration":15,"description":"descr 600"},{"duration":3600,"description":"descr 15"}]', 7845),
('c95c630e-33d1-47e3-bc7f-192f180f841d', 'name98', 'descr2', '[{"duration":300,"description":"descr 3600"},{"duration":600,"description":"descr 600"},{"duration":15,"description":"descr 30"}]', 915),
('9ef8f2d4-dfed-4176-9b36-edbc448a7d81', 'name99', 'descr1', '[{"duration":30,"description":"descr 3600"}]', 30),
('4d99794f-65d1-4cb3-a487-cafb1c5f996d', 'name100', 'descr2', '[{"duration":3600,"description":"descr 15"},{"duration":1800,"description":"descr 3600"},{"duration":600,"description":"descr 600"},{"duration":30,"description":"descr 300"},{"duration":30,"description":"descr 300"}]', 6060);

insert into rec.recipe_ingredients (recipe_id, name) values
(1, 'banana'),
(2, 'orange'),
(2, 'potato'),
(3, 'milk'),
(4, 'water'),
(5, 'egg'),
(5, 'papaya'),
(5, 'banana'),
(5, 'orange'),
(6, 'potato'),
(6, 'milk'),
(6, 'water'),
(6, 'egg'),
(6, 'papaya'),
(7, 'banana'),
(8, 'orange'),
(8, 'potato'),
(8, 'milk'),
(9, 'water'),
(9, 'egg'),
(10, 'papaya'),
(11, 'banana'),
(11, 'orange'),
(11, 'potato'),
(11, 'milk'),
(12, 'water'),
(12, 'egg'),
(12, 'papaya'),
(13, 'banana'),
(13, 'orange'),
(13, 'potato'),
(14, 'milk'),
(14, 'water'),
(14, 'egg'),
(15, 'papaya'),
(15, 'banana'),
(16, 'orange'),
(16, 'potato'),
(16, 'milk'),
(16, 'water'),
(17, 'egg'),
(18, 'papaya'),
(18, 'banana'),
(18, 'orange'),
(18, 'potato'),
(19, 'milk'),
(19, 'water'),
(19, 'egg'),
(20, 'papaya'),
(20, 'banana'),
(20, 'orange'),
(21, 'potato'),
(21, 'milk'),
(21, 'water'),
(21, 'egg'),
(21, 'papaya'),
(22, 'banana'),
(22, 'orange'),
(22, 'potato'),
(22, 'milk'),
(23, 'water'),
(23, 'egg'),
(24, 'papaya'),
(25, 'banana'),
(25, 'orange'),
(25, 'potato'),
(26, 'milk'),
(26, 'water'),
(27, 'egg'),
(28, 'papaya'),
(28, 'banana'),
(29, 'orange'),
(29, 'potato'),
(29, 'milk'),
(29, 'water'),
(30, 'egg'),
(30, 'papaya'),
(30, 'banana'),
(31, 'orange'),
(31, 'potato'),
(32, 'milk'),
(32, 'water'),
(32, 'egg'),
(32, 'papaya'),
(32, 'banana'),
(33, 'orange'),
(34, 'potato'),
(34, 'milk'),
(34, 'water'),
(34, 'egg'),
(34, 'papaya'),
(35, 'banana'),
(35, 'orange'),
(35, 'potato'),
(35, 'milk'),
(35, 'water'),
(36, 'egg'),
(36, 'papaya'),
(36, 'banana'),
(36, 'orange'),
(37, 'potato'),
(37, 'milk'),
(37, 'water'),
(38, 'egg'),
(38, 'papaya'),
(38, 'banana'),
(38, 'orange'),
(39, 'potato'),
(39, 'milk'),
(39, 'water'),
(39, 'egg'),
(39, 'papaya'),
(40, 'banana'),
(40, 'orange'),
(40, 'potato'),
(40, 'milk'),
(41, 'water'),
(41, 'egg'),
(42, 'papaya'),
(43, 'banana'),
(43, 'orange'),
(43, 'potato'),
(44, 'milk'),
(44, 'water'),
(44, 'egg'),
(44, 'papaya'),
(44, 'banana'),
(45, 'orange'),
(45, 'potato'),
(45, 'milk'),
(45, 'water'),
(46, 'egg'),
(46, 'papaya'),
(47, 'banana'),
(47, 'orange'),
(47, 'potato'),
(47, 'milk'),
(47, 'water'),
(48, 'egg'),
(48, 'papaya'),
(48, 'banana'),
(48, 'orange'),
(48, 'potato'),
(49, 'milk'),
(49, 'water'),
(49, 'egg'),
(50, 'papaya'),
(50, 'banana'),
(50, 'orange'),
(50, 'potato'),
(51, 'milk'),
(51, 'water'),
(51, 'egg'),
(52, 'papaya'),
(53, 'banana'),
(53, 'orange'),
(53, 'potato'),
(54, 'milk'),
(54, 'water'),
(54, 'egg'),
(55, 'papaya'),
(55, 'banana'),
(55, 'orange'),
(55, 'potato'),
(56, 'milk'),
(56, 'water'),
(56, 'egg'),
(56, 'papaya'),
(57, 'banana'),
(57, 'orange'),
(57, 'potato'),
(58, 'milk'),
(59, 'water'),
(59, 'egg'),
(60, 'papaya'),
(61, 'banana'),
(61, 'orange'),
(62, 'potato'),
(62, 'milk'),
(62, 'water'),
(63, 'egg'),
(63, 'papaya'),
(63, 'banana'),
(63, 'orange'),
(63, 'potato'),
(64, 'milk'),
(64, 'water'),
(64, 'egg'),
(64, 'papaya'),
(65, 'banana'),
(65, 'orange'),
(65, 'potato'),
(66, 'milk'),
(66, 'water'),
(66, 'egg'),
(66, 'papaya'),
(66, 'banana'),
(67, 'orange'),
(67, 'potato'),
(67, 'milk'),
(67, 'water'),
(67, 'egg'),
(68, 'papaya'),
(68, 'banana'),
(68, 'orange'),
(69, 'potato'),
(69, 'milk'),
(69, 'water'),
(70, 'egg'),
(70, 'papaya'),
(71, 'banana'),
(71, 'orange'),
(71, 'potato'),
(71, 'milk'),
(71, 'water'),
(72, 'egg'),
(72, 'papaya'),
(72, 'banana'),
(72, 'orange'),
(73, 'potato'),
(73, 'milk'),
(73, 'water'),
(73, 'egg'),
(74, 'papaya'),
(74, 'banana'),
(74, 'orange'),
(74, 'potato'),
(74, 'milk'),
(75, 'water'),
(75, 'egg'),
(75, 'papaya'),
(76, 'banana'),
(77, 'orange'),
(77, 'potato'),
(77, 'milk'),
(78, 'water'),
(79, 'egg'),
(79, 'papaya'),
(79, 'banana'),
(79, 'orange'),
(79, 'potato'),
(80, 'milk'),
(80, 'water'),
(80, 'egg'),
(80, 'papaya'),
(81, 'banana'),
(81, 'orange'),
(81, 'potato'),
(82, 'milk'),
(82, 'water'),
(82, 'egg'),
(83, 'papaya'),
(83, 'banana'),
(84, 'orange'),
(84, 'potato'),
(84, 'milk'),
(84, 'water'),
(84, 'egg'),
(85, 'papaya'),
(86, 'banana'),
(86, 'orange'),
(86, 'potato'),
(86, 'milk'),
(86, 'water'),
(87, 'egg'),
(87, 'papaya'),
(87, 'banana'),
(87, 'orange'),
(87, 'potato'),
(88, 'milk'),
(88, 'water'),
(89, 'egg'),
(89, 'papaya'),
(89, 'banana'),
(90, 'orange'),
(90, 'potato'),
(90, 'milk'),
(90, 'water'),
(91, 'egg'),
(92, 'papaya'),
(93, 'banana'),
(93, 'orange'),
(93, 'potato'),
(93, 'milk'),
(94, 'water'),
(94, 'egg'),
(95, 'papaya'),
(95, 'banana'),
(95, 'orange'),
(95, 'potato'),
(95, 'milk'),
(96, 'water'),
(96, 'egg'),
(97, 'papaya'),
(97, 'banana'),
(97, 'orange'),
(97, 'potato'),
(97, 'milk'),
(98, 'water'),
(98, 'egg'),
(98, 'papaya'),
(99, 'banana'),
(99, 'orange'),
(99, 'potato'),
(99, 'milk'),
(99, 'water'),
(100, 'egg'),
(100, 'papaya'),
(100, 'banana');