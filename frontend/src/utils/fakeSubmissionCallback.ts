// only for testing purposes
export type Languages = "java" | "javascript" | "php" | "python" | "cpp" | "c";

interface Payload {
  questionNumber: number;
  language: Languages;
  code: string;
}

interface SuccessResponse {
  data: Payload & {
    submissionType: "submit" | "refactor";
  };
  statusCode: 200;
}

interface ErrorResponse {
  data: {
    message: string;
  };
  statusCode: 500;
}

interface Options {
  onSuccess: (res: SuccessResponse) => void;
  onError: (err: ErrorResponse) => void;
}

const cachedData: Payload[] = [];

function getSubmissionType(questionNo: number) {
  const idx = cachedData.findIndex((cache) => cache.questionNumber === questionNo);

  if (idx > -1) {
    return "refactor";
  }

  return "submit";
}

function generateResponse(type: "success" | "error", payload: Payload) {
  if (type === "success") {
    return {
      data: {
        ...payload,
        submissionType: getSubmissionType(payload.questionNumber)
      },
      statusCode: 200
    };
  }

  return {
    data: {
      message: "Failed to Submit"
    },
    statusCode: 500
  };
}

export async function mutate(
  payload: Payload,
  { onSuccess, onError }: Partial<Options>
) {
  const fakePromise = new Promise<SuccessResponse>((resolve, reject) => {
    if (payload.questionNumber !== undefined && !Number.isNaN(payload.questionNumber)) {
      switch (true) {
        case payload.questionNumber === 0 || payload.questionNumber === 2:
          resolve(generateResponse("success", payload) as SuccessResponse);
          break;
        default:
          reject("only accepts question no 0  and 2");
      }
    } else {
      reject("please provide question no");
    }
  });

  try {
    const fakeRes = await fakePromise;
    if (fakeRes) {
      cachedData.push(fakeRes.data);
      onSuccess?.(fakeRes);
    }
  } catch (err) {
    console.log(err);
    onError?.(generateResponse("error", payload) as ErrorResponse);
  }
}
