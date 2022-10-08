import React, { type SVGProps } from "react";

export function GradeIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 12 12" {...props}>
      <path
        fill="currentColor"
        d="M5.16 2.189a1.962 1.962 0 0 1 1.68 0l4.874 2.309a.5.5 0 0 1 .008.9l-4.85 2.406a1.962 1.962 0 0 1-1.744 0L1 5.756V8a.5.5 0 0 1-1 0V4.975a.502.502 0 0 1 .286-.477l4.874-2.31ZM2 7.369V9a.5.5 0 0 0 .147.354l.002.003l.023.021l.06.056a6.738 6.738 0 0 0 1.012.745C3.912 10.58 4.877 11 6 11c1.123 0 2.088-.42 2.757-.821a6.738 6.738 0 0 0 1.012-.745l.06-.056l.016-.016l.006-.006l.001-.001l.002-.001A.5.5 0 0 0 10 9V7.368L7.316 8.7a2.962 2.962 0 0 1-2.632 0L2 7.368Z"
      ></path>
    </svg>
  );
}
