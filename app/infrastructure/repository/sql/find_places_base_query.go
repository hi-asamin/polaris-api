package sql

func FindPlacesBaseQuery() string {
	return `
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
  WHERE
  (
    -- 検索キーワード条件
    %s
  )
  AND
  (
    -- カーソル条件
    ($2::UUID IS NULL OR "Media".id > $2::UUID) 
  )
  ORDER BY
    "Media".id ASC
  LIMIT $1::INTEGER;
  `
}
