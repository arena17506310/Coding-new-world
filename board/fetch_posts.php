<?php

$host = 'svc.sel5.cloudtype.app:30279';
$db   = 'alexandria';
$user = 'root';
$pass = 'camel99';

$conn = new mysqli($host, $user, $pass, $db);

if ($conn->connect_error) {
  die('Connection failed:'.$conn->connect_error);
}

$sql = "SELECT title,content FROM posts ORDER BY id DESC";
$result = $conn->query($sql);

if ($result->num_rows > 0) {
  while($row = $result->fetch_assoc()) {
    echo "<h3>" . $row["title"]. "</h3><p>" . $row["content"]. "</p><hr>";
  }
} else {
  echo "No posts found.";
}

$conn->close();

?>
