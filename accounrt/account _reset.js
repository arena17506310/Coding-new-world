$(document).ready(function () {
  $("#signup-form").submit(function (e) {
    e.preventDefault();

    var username = $("#username").val();

    $(".message").removeClass("error success");

    if (username === "") {
      $(".message").addClass("error").text("빈 칸에 정보를 모두 기입하여 주세요.");
      return;
    }

    // AJAX 요청 ( 백엔드가 건들면 ㅈ된다 개새기들아 )
    $.ajax({
      url: "signup.php",
      type: "POST",
      data: { username: username },
      success: function (response) {
        $(".message").addClass("success").text("비밀번호 초기화 요청 성공");
        // 성공적으로 처리된 후 실행할 코드 작성 ( 백엔드가 건들면 ㅈ된다 개새기들아 )
      },
      error: function () {
        $(".message")
          .addClass("error")
          .text("서버와의 통신 중 오류가 발생했습니다.");
        // 오류 발생 시 실행할 코드 작성 ( 통신 에러 말고 건들면 죽인다 )
      }
    });
  });
});
