<?php

$host = 'mariadb';
$database = 'app';
$user = 'root';
$pass = 'root';

$pdo = new \PDO(sprintf('mysql:host=%s;dbname=%s', $host, $database), $user, $pass);

$rows=0;
for ($i=0; $i < 1000; $i++) {
    $count = $pdo->exec('INSERT INTO jobs(name) VALUES(UUID())');
    echo sprintf('Added %s record', $rows+=$count), "\n";
}
