$(document).ready(function () {
  $("#post-form").on("submit", function (event) {
    event.preventDefault();

    var title = $("#title").val();
    var content = $("#content").val();

    $.ajax({
      url: "create_post.php",
      method: "POST",
      data: { title: title, content: content },
      success: function (data) {
        alert(data);
        fetchPosts();
        $("#post-form").trigger("reset");
      },
    });
  });

  function fetchPosts() {
    $.ajax({
      url: "fetch_posts.php",
      method: "GET",
      success: function (data) {
        $("#posts").html(data);
      },
    });
  }

  fetchPosts();
});
