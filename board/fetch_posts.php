<?php

$host = 'svc.sel5.cloudtype.app:30279';
$db   = 'alexandria';
$user = 'root';
$pass = 'camel99';

$conn = new mysqli($host, $user, $pass, $db);

if ($conn->connect_error) {
  die('Connection failed:'.$conn->connect_error);
}

$sql = "SELECT id, title FROM posts ORDER BY id DESC";
$result= $conn ->query($sql);

if ($result ->num_rows >0) {
 while($row= $result ->fetch_assoc()) {

   // 게시물 제목을 하이퍼링크로 출력
   echo "<a href='post_detail.php?id=".$row["id"]."'>".$row["title"]."</a><br>";
 }
} else{
 echo"No posts found.";
}

$conn ->close();

?>