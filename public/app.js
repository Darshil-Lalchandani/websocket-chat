$(function () {
    let num = Math.floor(Math.random() * 100);
    const websocket = new WebSocket("ws://" + window.location.host + "/websocket?id=" + num);
    let room = $("#chat-text");
    room.append(num)
    websocket.addEventListener("message", function (e) {
      let data = JSON.parse(e.data);
      let chatContent = `<p><strong>${data.text}</strong></p>`;
      room.append(chatContent);
      room.scrollTop = room.scrollHeight; // Auto scroll to the bottom
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