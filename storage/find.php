<?php


$set = [
    3 => [
        "id" => 3,
        "name" => "a",
        "quantity" => 5,
        "price" => 44.6
    ],
    9 => [
        "id" => 9,
        "name" => "b",
        "quantity" => 7,
        "price" => 3.50
    ],
    15 => [
        "id" => 15,
        "name" => "c",
        "quantity" => 9,
        "price" => 53.6
    ],
    20 => [
        "id" => 20,
        "name" => "d",
        "quantity" => 477,
        "price" => 365.0
    ],
];

$result = [];
if (isset($_GET['id'])) {
    if (isset($set[$_GET['id']])) {
        $result = $set[$_GET['id']];
    }else{
        http_response_code(404);
        return;
    }
} elseif (isset($_GET['ids']) && is_array($_GET['ids'])) {
    $result["items"] = [];
    foreach ($_GET['ids'] as $id) {
        if (isset($set[$id])) {
            $result["items"][] = $set[$id];
        }
    }
} else {
    http_response_code(400);
    return;
}


header("Content-type:application/json");


echo json_encode($result);
