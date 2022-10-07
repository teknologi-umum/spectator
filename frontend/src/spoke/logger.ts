import { LOGGER_URL } from "@/constants";
import { ILogger, LogLevel } from "@microsoft/signalr";

export class Logger implements ILogger {
  private readonly _loggerUrl: string;

  constructor(loggerUrl: string) {
    if (loggerUrl === "" || loggerUrl === null || loggerUrl === undefined) {
      throw new TypeError("Logger URL is not defined.");
    }

    this._loggerUrl = loggerUrl;
  }

  public async log(logLevel: LogLevel, message: string): Promise<void> {
    return; // TODO: remove
    let level: "error" | "warning" | "info" | "debug";
    switch (logLevel) {
      case LogLevel.Critical:
      case LogLevel.Error:
        level = "error";
        break;
      case LogLevel.Warning:
        level = "warning";
        break;
      case LogLevel.Information:
        level = "info";
        break;
      default:
        level = "debug";
    }

    if (import.meta.env.DEV) {
      console.log(`[${logLevel}] ${message}`);

      // no need to send these ones to the logger endpoint
      if (logLevel === LogLevel.Trace || logLevel === LogLevel.Debug) {
        return;
      }
    }

    try {
      await fetch(this._loggerUrl, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({
          level,
          message,
          timestamp: new Date(Date.now())
        })
      });
    } catch (err) {
      console.error(err);
    }
  }
}

export const loggerInstance = new Logger(LOGGER_URL);
