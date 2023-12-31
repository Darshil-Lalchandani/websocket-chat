$(function () {
  let num = Math.floor(Math.random() * 100);
  const websocket = new WebSocket(
    "ws://" + window.location.host + "/websocket/connect?id=" + num
  );
  let receivedMessages = $("#chat-text");
  receivedMessages.append(num);
  websocket.addEventListener("message", function (e) {
    let data = JSON.parse(e.data);
    let chatContent = `<p><strong>${data.text}</strong></p>`;
    receivedMessages.append(chatContent);
  });
  $("#input-form").on("submit", function (event) {
    event.preventDefault();
    let text = $("#input-text")[0].value;
    websocket.send(
      JSON.stringify({
        text: text,
      })
    );
  });
});
