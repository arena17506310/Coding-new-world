<?php

// Database connection details from signup.php

$host = 'svc.sel5.cloudtype.app:30279';
$db   = 'alexandria';
$user = 'root';
$pass = 'camel99';

$conn = new mysqli($host, $user, $pass, $db);

if ($conn->connect_error) {
  die('Connection failed:'.$conn->connect_error);
}

$title = $_POST['title'];
$content = $_POST['content'];

$sql = "INSERT INTO posts(title,content) VALUES (?,?)";
$stmt= $conn->prepare($sql);
$stmt->bind_param("ss", $title,$content);

if ($stmt->execute()) {
  echo 'Post created successfully.';
} else {
  echo 'Error creating post.';
}

$conn->close();

?>
