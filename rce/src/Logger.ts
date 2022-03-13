import { randomUUID } from "crypto";
import console from "console";
import grpc from "@grpc/grpc-js";
import { Environment, Level } from "@/stub/logger_pb";
import { LoggerClient } from "@/stub/logger_pb.grpc-client";

export class Logger {
  private _loggerClient: LoggerClient;
  private readonly _language = "Javascript";
  private environment: string;

  constructor(
    _loggerServerAddress = "",
    // eslint-disable-next-line no-unused-vars
    private readonly _loggerToken = "",
    // eslint-disable-next-line no-unused-vars
    private readonly _environment = ""
  ) {
    this._loggerClient = new LoggerClient(
      _loggerServerAddress,
      grpc.credentials.createInsecure()
    );

    this.environment = process.env?.ENVIRONMENT ?? "development";
  }

  private _getLogEnvironment() {
    switch (this._environment) {
      case "development":
        return Environment.DEVELOPMENT;
      case "production":
        return Environment.PRODUCTION;
      case "staging":
        return Environment.STAGING;
      case "testing":
        return Environment.TESTING;
      default:
        return Environment.UNSET;
    }
  }

  public log(
    message: string,
    level: Level,
    requestID: string,
    body: Record<string, string>
  ): void {
    if (this.environment !== "production") {
      console.log(`[${this._language}] ${message}`);
      return;
    }

    const env = this._getLogEnvironment();

    if (env !== Environment.PRODUCTION) {
      // eslint-disable-next-line no-console
      console.log(message);
    }

    if (requestID === "") {
      requestID = randomUUID();
    }

    this._loggerClient.createLog(
      {
        accessToken: this._loggerToken,
        data: {
          requestId: requestID,
          application: "rce",
          language: this._language,
          body: body,
          message: message,
          level: level,
          environment: env
        }
      },
      (err) => {
        if (err !== null) {
          // eslint-disable-next-line no-console
          console.error(
            `An error has occured while trying to create a log to the logger service: ${err}\n\nTrying to send: ${message}`
          );
        }
      }
    );
  }
}
