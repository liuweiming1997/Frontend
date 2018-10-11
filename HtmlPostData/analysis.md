如果仔细观察一个Form的提交，你就会发现，一旦用户点击“Submit”按钮，表单开始提交，
浏览器就会刷新页面，然后在新页面里告诉你操作是成功了还是失败了。
如果不幸由于网络太慢或者其他原因，就会得到一个404页面。

这就是Web的运作原理：一次HTTP请求对应一个页面。

如果要让用户留在当前页面中，同时发出新的HTTP请求，就必须用JavaScript发送这个新请求，
接收到数据后，再用JavaScript更新页面，这样一来，用户就感觉自己仍然停留在当前页面，
但是数据却可以不断地更新。