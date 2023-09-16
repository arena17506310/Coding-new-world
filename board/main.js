var page = 1; // 페이지 번호 초기화
var isLoading = false; // 데이터 로딩 중인지 여부

// 스크롤 이벤트 감지
window.addEventListener("scroll", function () {
  if (window.innerHeight + window.scrollY >= document.body.offsetHeight) {
    // 페이지 하단에 도달하면
    loadMoreData(); // 추가 데이터 로드
  }
});

// 초기 데이터 로드
loadMoreData();

// 추가 데이터 로드 함수 정의
function loadMoreData() {
  if (!isLoading) {
    // 데이터 로딩 중이 아닌 경우에만 실행

    isLoading = true; // 데이터 로딩 상태로 변경

    var xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function () {
      if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
        // 요청 완료 및 성공적인 응답인 경우

        var response = JSON.parse(xhr.responseText); // 응답 데이터 파싱

        var contentDiv = document.getElementById("content");

        for (var i = 0; i < response.length; i++) {
          var itemDiv = document.createElement("div");
          itemDiv.className = "item";
          itemDiv.textContent = response[i];

          contentDiv.appendChild(itemDiv);
        }

        page++; // 페이지 번호 증가

        isLoading = false; // 데이터 로딩 완료 상태로 변경
      }
    };

    xhr.open("GET", "data.php?page=" + page, true); // 서버에서 새로운 데이터 가져오기 요청

    xhr.send();
  }
}
