WITH places_with_distance AS (
  SELECT 
    "Place".id AS pid,
    "Place".name,
    "Media".id AS mid, 
    "Media".media_url AS src,
    "Media".media_type AS type,
    "Media".alt_text AS alt,
    ST_Distance(
      "Place".geometry, 
      ST_SetSRID(ST_MakePoint($1::DOUBLE PRECISION, $2::DOUBLE PRECISION), 4326)
    ) AS distance
  FROM
    "Place"
  INNER JOIN
    "Media" 
    ON
      "Place".id = "Media".place_id
  WHERE
    "Place".geometry IS NOT NULL
)
SELECT *
FROM 
  places_with_distance
WHERE 
  (
    -- 初回リクエスト時またはカーソル条件に該当するデータをフィルタリング
    ($3::DOUBLE PRECISION IS NULL OR distance > $3::DOUBLE PRECISION) OR 
    (distance = $3::DOUBLE PRECISION AND ($4::UUID IS NULL OR pid > $4::UUID)) OR 
    (distance = $3::DOUBLE PRECISION AND pid = $4::UUID AND ($5::UUID IS NULL OR mid > $5::UUID))
  )
ORDER BY
  distance ASC,
  pid ASC,
  mid ASC
LIMIT $6::INTEGER;
