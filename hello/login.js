$(document).ready(function () {
  $("#re_main_button").click(function (e) {
    e.preventDefault();
    window.location.href = "../account/ccount.html";
  });

  $(".login-form").submit(function (e) {
    e.preventDefault();

    var username = $("#username").val();
    var password = $("#password").val();

    // AJAX 요청
    $.ajax({
      url: "login.php",
      type: "POST",
      data: {
        username: username,
        password: password,
      },
      success: function (response) {
        if (response.trim() === "Success") {
          window.location.href = "../board/main.html"; // 로그인 성공 시 메인 페이지로 리다이렉트
        } else {
          alert("로그인 실패 : " + response); // 실패 시 서버의 에러 메시지 출력
        }
      },
    });
  });
});
