import { grpc } from "@improbable-eng/grpc-web";
import { MyService } from "./generated/myservice_pb_service"; // Файлы, сгенерированные protoc
import { MyRequest } from "./generated/myservice_pb";

export function callMyMethod(name) {
  return new Promise((resolve, reject) => {
    const request = new MyRequest();
    request.setName(name);

    grpc.unary(MyService.MyMethod, {
      request: request,
      host: "http://localhost:8080", // Адрес gRPC-Web сервера
      onEnd: (res) => {
        if (res.status === grpc.Code.OK) {
          resolve(res.message.toObject());
        } else {
          reject(res);
        }
      },
    });
  });
}
