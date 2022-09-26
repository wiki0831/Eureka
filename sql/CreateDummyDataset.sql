-- Creating 100,000 dummy polygons
DROP Table if exists public."polygon_no_idx";
CREATE TABLE public."polygon_no_idx" as
select row_number() OVER () AS name,
  ST_Expand(randomPoints.geom::geometry, random() / 10)::geography as geom
from (
    select (
        ST_DumpPoints(ST_GeneratePoints(boundingbox, 100000))
      ).geom::geography
    from ST_GeomFromGeoHash('d') as boundingbox
  ) as randomPoints;
DROP Table if exists public."polygon_gist";
create table public."polygon_gist" as (
  select *
  from public."polygon_no_idx"
);
CREATE INDEX polygon_gist_idx ON public."polygon_gist" USING GIST (geom);

DROP Table if exists public."polygon_spgist";
create table public."polygon_spgist" as (
  select *
  from public."polygon_no_idx"
);
CREATE INDEX polygon_spgist_idx ON public."polygon_spgist" USING spgist (geom);
DROP Table if exists public."polygon_brin";
create table public."polygon_brin" as (
  select *
  from public."polygon_no_idx"
);
CREATE INDEX polygon_brin_idx ON public."polygon_brin" USING brin(geom);
-- Creating 100,000 dummy Points
DROP Table if exists public."point_no_idx";
CREATE TABLE public."point_no_idx" as
select row_number() OVER () AS name,
  randomPoints.geom::geography as geom
from (
    select (
        ST_DumpPoints(ST_GeneratePoints(boundingBox, 100000))
      ).geom::geography
    from ST_GeomFromGeoHash('d') as boundingBox
  ) as randomPoints;
DROP Table if exists public."point_gist";
create table public."point_gist" as (
  select *
  from public."point_no_idx"
);
CREATE INDEX point_gist_idx ON public."point_gist" USING GIST (geom);
DROP Table if exists public."point_spgist";
create table public."point_spgist" as (
  select *
  from public."point_no_idx"
);
CREATE INDEX point_spgist_idx ON public."point_spgist" USING spgist (geom);
DROP Table if exists public."point_brin";
create table public."point_brin" as (
  select *
  from public."point_no_idx"
);
CREATE INDEX point_brin_idx ON public."point_brin" USING brin (geom);
-- Creating 100,000 dummy circle
DROP Table if exists public."circle_no_idx";
CREATE TABLE public."circle_no_idx" as
select row_number() OVER () AS name,
  st_buffer(randomPoints.geom, random() * 4000, 4)::geography as geom
from (
    select (
        ST_DumpPoints(ST_GeneratePoints(boundingbox, 100000))
      ).geom::geography
    from ST_GeomFromGeoHash('d') as boundingbox
  ) as randomPoints;
DROP Table if exists public."circle_gist";
create table public."circle_gist" as (
  select *
  from public."circle_no_idx"
);
CREATE INDEX circle_gist_idx ON public."circle_gist" USING GIST (geom);