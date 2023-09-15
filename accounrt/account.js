$(document).ready(function () {
  $("#signup-form").submit(function (e) {
    e.preventDefault();

    var username = $("#username").val();

    $(".message").removeClass("error success");

    if (username === "" || password === "") {
      $(".message").addClass("error").text("빈 칸에 요청하는 정보를 모두 기입 하여 주세요");
      return;
    }

    // 여기서부터 서버와의 통신 등 추가 작업을 진행할 수 있습니다.

    $(".message").addClass("success").text("회원가입 성공!");
  });
});
