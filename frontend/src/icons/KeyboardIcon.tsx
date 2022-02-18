import React from "react";
import type { SVGProps } from "react";

export function KeyboardIcon(props: SVGProps<SVGSVGElement>) {
  return (
    <svg width="1em" height="1em" viewBox="0 0 20 20" {...props}>
      <g fill="none">
        <path
          d="M3.5 4A1.5 1.5 0 0 0 2 5.5v8A1.5 1.5 0 0 0 3.5 15h13a1.5 1.5 0 0 0 1.5-1.5v-8A1.5 1.5 0 0 0 16.5 4h-13zm2.755 3.252a.752.752 0 1 1-1.505 0a.752.752 0 0 1 1.505 0zm6 0a.752.752 0 1 1-1.505 0a.752.752 0 0 1 1.505 0zM5 12.5a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5zm9.502-4.495a.752.752 0 1 1 0-1.505a.752.752 0 0 1 0 1.505zm-7.504 2.5a.752.752 0 1 1 0-1.505a.752.752 0 0 1 0 1.505zm3.757-.753a.752.752 0 1 1-1.505 0a.752.752 0 0 1 1.505 0zm2.252.753a.752.752 0 1 1 0-1.505a.752.752 0 0 1 0 1.505zM9.255 7.252a.752.752 0 1 1-1.505 0a.752.752 0 0 1 1.505 0z"
          fill="currentColor"
        ></path>
      </g>
    </svg>
  );
}
