$(document).ready(function () {
  var urlParams = new URLSearchParams(window.location.search);
  var username = urlParams.get("username");

  if (username) {
    $("#username").val(username);
  }
  $("#login-form").submit(function (e) {
    e.preventDefault();

    var username = $("#username").val();
    var password = $("#password").val();

    $(".message").removeClass("error success");

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
          $(".message").addClass("success").text("로그인 성공");
          window.location.href = "../board/main.html";
        } else {
          $(".message")
            .addClass("error")
            .text(response + "서버와 연결에 실패했습니다.");
        }
      },
    });
  });
});
