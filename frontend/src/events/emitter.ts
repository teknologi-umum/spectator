/**
 * `emit` will emit a SignalR event to the server.
 * @param connection - Connection to SignalR hub
 * @param data - The data you want to send
 */
// eslint-disable-next-line require-await
export async function emit<T>(connection: unknown, data: T) {
  // TODO(elianiva):
  // - nanti kira kira bakal gini implementasinya:
  //   await connection.invoke("SendMessage", data);
  //   tapi karena servernya belom ada, yaudah pake console.log dulu
  // - ga perlu emit id user lewat sini, udah ada jwt di header requestnya
  // - perlu pake await? atau langsung `return connection.invoke` aja?

  // eslint-disable-next-line no-console
  console.log(data);
}
