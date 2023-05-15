<?php
header("Content-type:application/json");
echo json_encode(
    [
        "page"=>0,
        "total" => 4,
        "resultIds"=>[3,15,9,20]
    ]
);
