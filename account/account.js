$(document).ready(function () {
  $("#signup-form").submit(function (e) {
    e.preventDefault();

    var username = $("#username").val(); // username 필드 추가
    var password = $("#password").val(); // password 필드 추가
    var phoneNum = $("#phoneNum").val();
    var schoolNum = $("#schoolNum").val();

    $(".message").removeClass("error success");

    if (username === "" || password === "") {
      $(".message")
        .addClass("error")
        .text("빈 칸에 정보를 모두 기입하여 주세요.");
      return;
    }
    // 하시발살려줘
    // AJAX 요청
    $.ajax({
      url: "signup.php",
      type: "POST",
      data: {
        username: username,
        password: password,
        phoneNum: phoneNum,
        schoolNum: schoolNum,
      }, // 비밀번호 및 아이디 전송 ( 건들지 마라 뒤진다 ) -> 전화번호랑 학번 저장 안함? 구별은 해야지
      success: function (response) {
        if (response === "Success") {
          // 가정: signup.php 에서 성공 시 'Success' 문자열 반환
          $(".message").addClass("success").text("회원가입 요청 성공");
          window.location.href = "/"; // 메인 페이지로 리다이렉트
        } else {
          $(".message")
            .addClass("error")
            .text(response + "\n서버와연결에실패했습니다."); // 실패 시 서버의 에러 메시지 출력
        }
      },
    });
  });
});
