import React from "react";
import type { SVGProps } from "react";

export function TemperatureIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 32 32" {...props}>
      <path
        d="M9 4v11.469C7.25 16.738 6 18.676 6 21c0 3.855 3.145 7 7 7s7-3.145 7-7c0-2.324-1.25-4.262-3-5.531V14h2v-2h-2v-2h2V8h-2V6h2V4H9zm15.5 0C22.57 4 21 5.57 21 7.5s1.57 3.5 3.5 3.5S28 9.43 28 7.5S26.43 4 24.5 4zM11 6h4v10.406l.5.282c1.496.867 2.5 2.46 2.5 4.312c0 2.773-2.227 5-5 5s-5-2.227-5-5c0-1.852 1.004-3.445 2.5-4.313l.5-.28V6zm13.5 0c.827 0 1.5.673 1.5 1.5S25.327 9 24.5 9S23 8.327 23 7.5S23.673 6 24.5 6zM12 16v2.188c-1.16.414-2 1.513-2 2.814a3.001 3.001 0 0 0 6 0c0-1.301-.84-2.4-2-2.814V16h-2z"
        fill="currentColor"
      ></path>
    </svg>
  );
}
