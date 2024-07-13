<?php

$mysqli = require __DIR__ . "/config.php";

$sql = sprintf("SELECT * FROM user
                WHERE id = '%d'",
                $mysqli->real_escape_string($_GET["id"]));
                
$result = $mysqli->query($sql);

$is_available = $result->num_rows === 0;

header("Content-Type: application/json");

echo json_encode(["available" => $is_available]);