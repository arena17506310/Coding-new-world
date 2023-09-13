$(document).ready(function() {
   $("#signup-form").submit(function(e) {
     e.preventDefault();

     var username = $("#username").val();
     var password = $("#password").val();

     $(".message").removeClass("error success");

     if (username === "" || password === "") {
       $(".message").addClass("error").text("사용자명과 비밀번호를 입력하세요.");
       return;
     }

     // 여기서부터 서버와의 통신 등 추가 작업을 진행할 수 있습니다.

    $(".message").addClass("success").text("회원가입 성공!");
   });
});