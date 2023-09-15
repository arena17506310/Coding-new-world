$(document).ready(function () {
  $("#signup-form").submit(function (e) {
    e.preventDefault();

    var username = $("#username").val(); // username 필드 추가
    var password = $("#password").val(); // password 필드 추가

    $(".message").removeClass("error success");

    if (username === "" || password === "") {
      $(".message").addClass("error").text("빈 칸에 정보를 모두 기입하여 주세요.");
      return;
    }

    // AJAX 요청
    $.ajax({
      url: "signup.php",
      type: "POST",
      data: { username: username, password: password }, // 비밀번호 및 아이디 전송 ( 건들지 마라 뒤진다 )
      success: function (response) {
        $(".message").addClass("success").text("회원가입 요청 성공");
        // 성공적으로 처리된 후 실행할 코드 작성
      },
      error: function () {
        $(".message")
          .addClass("error")
          .text("서버와의 통신 중 오류가 발생했습니다.");
        // 오류 발생 시 실행할 코드 작성
      }
    });
  });
});
