package sql

func FindAllPlaces() string {
	return `
  WITH filtered_places AS (
    SELECT DISTINCT "Place".id
    FROM "Place"
    LEFT JOIN "PlaceCategory" ON "Place".id = "PlaceCategory".place_id
    WHERE
      CASE 
        WHEN ARRAY_LENGTH($4::INTEGER[], 1) > 0 THEN
          "PlaceCategory".category_id = ANY($4::INTEGER[])
        ELSE TRUE
      END
  )
  SELECT 
    "Place".id AS pid,
    "Place".name,
    "Place".state,
    "Place".city,
    "Media".id AS mid, 
    "Media".media_url AS src,
    "Media".media_type AS type,
    "Media".alt_text AS alt
  FROM
    "Place"
  INNER JOIN
    "Media" 
    ON
      "Place".id = "Media".place_id
  INNER JOIN
    filtered_places
    ON "Place".id = filtered_places.id
  WHERE 
    (
      -- 初回リクエスト時またはカーソル条件に該当するデータをフィルタリング
      ($1::UUID IS NULL OR "Place".id > $1::UUID) OR 
      ("Place".id = $1::UUID AND ($2::UUID IS NULL OR "Media".id > $2::UUID))
    )
  ORDER BY
    "Place".id ASC,
    "Media".id ASC
  LIMIT $3::INTEGER;
  `
}
